package common

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
)

/*
bigquery.go contains the state and helper methods required for working with
a BigQuery data source and marshalling and unmarshalling structures betwee
protobuf and big query
*/

const (
	BigQueryDateFormat = "2006-01-02 15:04:05"
)

// BigQuery is a state holder for the bigquery.Client, and bigquery.Dataset.
// This allows for safe concurrent access to the datastore and its tables.
type BigQuery struct {
	client  *bigquery.Client
	dataset *bigquery.Dataset
	log     *log.Logger
	tables  map[string]*bigquery.Table
}

// NewBigQuery is the default constructor for the BigQuery object
func NewBigQuery(
	ctx context.Context,
	projectId string,
	log *log.Logger,
	datasetName string) (*BigQuery, error) {

	var bq *BigQuery
	client, err := bigquery.NewClient(ctx, projectId)

	if err == nil {
		bq = &BigQuery{
			client: client,
			log:    log,
			tables: make(map[string]*bigquery.Table)}

		bq.initializeDataset(ctx, datasetName)

		if bq.dataset == nil {
			err = errors.New("failed to create or read dataset")
		}

	}
	return bq, err
}

// GetOrCreateDataset Gets a dataset, if the dataset does not exist, it is created.
// The method should be invoked with a time safe context:
// ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
// defer cancel()
// GetOrCreateDataset(ctx, "test")
func (bq *BigQuery) initializeDataset(ctx context.Context, name string) {
	if err := bq.client.Dataset(name).Create(ctx, nil); err != nil {
		bq.log.Printf("\n\n========================== Failed to create name: %s; \n%v\n\n", name, err)
	}
	bq.dataset = bq.client.Dataset(name)
}

// Closes all client connections, this is kept package private to allow close to happen at the common level
func (bq *BigQuery) close() (err error) {
	return bq.client.Close()
}

// CreateTableFromSchema creates a table from a bq Schema
func (bq *BigQuery) CreateTableFromSchema(
	tableName string,
	schema bigquery.Schema,
	table *bigquery.Table,
	labels map[string]string) {

	// Allow a maximum of 5 seconds to create a table, fail otherwise
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := table.Create(ctx, &bigquery.TableMetadata{
		Name:   tableName,
		Schema: schema,
		Labels: labels,
	}); err != nil {
		log.Printf("Failed to create table %s\n\nwith schema:\n%v\n", tableName, err)
	}
}

// GetTable will get a table from BQ if the table exists, otherwise returns
// an error
func (bq *BigQuery) GetTable(tableName string) (table *bigquery.Table, err error) {
	table = bq.tables[tableName]
	if table == nil {
		err = errors.New(fmt.Sprintf("Table %s not initialized", tableName))
	}
	return table, err
}

// TableExists determines if the table is already in the dataset
func (bq *BigQuery) TableExists(tableName string) (out bool) {
	table, _ := bq.GetTable(tableName)

	if table == nil && bq.dataset != nil {
		// Timeout in 5 seconds (considering large dataset)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		tblItr := bq.dataset.Tables(ctx)
		out = false
		if tblItr != nil {
			for {
				tbl, err := tblItr.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return false
				}
				if tbl.TableID == tableName {
					// Register the table
					bq.tables[tableName] = tbl
					out = true
				}
			}
		}
	}
	return out
}

// InitializeTable - reads a table from the provided schema or creates the
// schema from the struct. Once the schema is created, it is used to create
// the table if the table does not exist. Once created or read, the table is cached
// for future use.
func (bq *BigQuery) InitializeTable(
	tableName string,
	schemaFileName string,
	t proto.Message,
	additionalTags map[string]string) (out *bigquery.Table, err error) {

	tags := make(map[string]string)

	for k, v := range additionalTags {
		tags[k] = v
	}

	if bq.dataset != nil {
		out = bq.dataset.Table(tableName)

		if !bq.TableExists(tableName) {
			fil, err := os.OpenFile(schemaFileName, os.O_RDONLY, 0644)
			if os.IsNotExist(err) {
				// Attempt to create the table from the struct
				log.Printf("Creating Schema from Type")

				schema := protobq.InferSchema(t)
				//schema, err := bigquery.InferSchema(t)
				if err != nil && schema != nil {
					log.Printf("failed to create schema from struct: %s\n%v", tableName, err)
				}
				bq.CreateTableFromSchema(tableName, schema, out, tags)
			} else {
				log.Printf("Creating Schema from file")
				schema, err := SchemaFromFile(fil)
				if err != nil {
					log.Printf("failed to create schema from file: %s\n%v", schemaFileName, err)
				}
				bq.CreateTableFromSchema(tableName, schema.Relax(), out, tags)
			}
		}
	}
	return out, err
}

// Utility Functions

// QueryableNameFromQualifiedName - removes the project name with : from the qualified name
func QueryableNameFromQualifiedName(qualifiedName string) string {
	idx := strings.Index(qualifiedName, ":")
	if idx > 0 {
		qualifiedName = qualifiedName[idx+1:]
	}
	return qualifiedName
}

// ConvertMessageToBigQueryType converts from a message into a message save type
func ConvertMessageToBigQueryType(t proto.Message) bigquery.ValueSaver {
	return &protobq.MessageSaver{Message: t}
}

// ConvertFromBigQueryToMessage converts from the BQ schema into a structured
// protobuf message type.
func ConvertFromBigQueryToMessage(t proto.Message) *MessageLoaderWrapper {
	wrapped := protobq.MessageLoader{Message: t}
	return &MessageLoaderWrapper{loader: wrapped}
}

// SchemaFromFile reads a JSON file from the operating system
// this is used during the creation phase of data assets.
// TODO - Verify this will work with proto bq, it may need to be dropped
func SchemaFromFile(fil *os.File) (schema bigquery.Schema, err error) {
	if fil != nil {
		reader := bufio.NewReader(fil)
		if reader.Size() > 0 {
			source, err := io.ReadAll(reader)
			if err == nil {
				schema, err = bigquery.SchemaFromJSON(source)
			}
		} else {
			err = errors.New("file size is zero")
		}
	} else {
		err = errors.New("file must not be nil")
	}
	return schema, err
}

func (bq *BigQuery) Query(sql string) *bigquery.Query {
	return bq.client.Query(sql)
}
