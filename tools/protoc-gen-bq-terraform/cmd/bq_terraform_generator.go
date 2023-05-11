package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"log"
	"text/template"
	"bufio"
	"path/filepath"
)

const (
	Space = " "
)

var (
	inputDirectory string
	outputFile     string
	projectId      string
	location string
	accessOwner string
	accessOwnerRole string
	accessReaderRole string
	accessReaderDomain string
	expirationTime string
	datasetId string
	tableId string
	timePartitionType string
	timePartitionExpirationMs string
	timePartitionField string

	debug bool
)
func init() {
	flag.BoolVar(&debug, "debug", false, "Determines if debug output is visible.")
	flag.StringVar(&inputDirectory, "input", "", "The input schema file to process.")
	flag.StringVar(&outputFile, "output", "", "The destination file to write to")
	flag.StringVar(&projectId, "project_id", "var.project", "The project id to use")
	flag.StringVar(&location, "location", "var.location", "The location of the tableset")
	flag.StringVar(&accessOwner, "access_owner", "var.access.owner", "The access owners email")
	flag.StringVar(&accessOwnerRole, "access_owner_role", "OWNER", "The name of the role of the owner")
	flag.StringVar(&accessReaderRole, "access_reader_role", "READER", "The name of the reader role")
	flag.StringVar(&accessReaderDomain, "access_reader_domain", "", "The domain name of the reader group")
	flag.StringVar(&expirationTime, "expiration_time", "", "The time in milliseconds the data should be accessible")
	flag.StringVar(&datasetId, "dataset_id", "", "The dataset ID")
	flag.StringVar(&tableId, "table_id", "", "The table ID")
	flag.StringVar(&timePartitionType, "time_partition_type", "DAY", "The time partition type.")
	flag.StringVar(&timePartitionExpirationMs, "time_partition_expiration_ms", "", "The number of milliseconds the data should be accessible via the index")
	flag.StringVar(&timePartitionField, "time_partition_field", "", "The table field to use for the time partition")
}

type BQConfiguration struct {
	InputDirectory string
	OutputFile     string
	ProjectId string
	Location string
	AccessOwner string
	AccessOwnerRole string
	AccessReaderRole string
	AccessReaderDomain string
	ExpirationTime string
	DatasetId string
	TableId string
	TimePartitionType string
	TimePartitionExpirationMills string
	TimePartitionField string
	Schema string
}

func (conf *BQConfiguration) PrintDebug() {
	if debug {
		log.Printf("Initialized protoc-gen-bq-terraform with:\n" +
			"Input Directory: %s\n" +
			"Output File: %s\n" +
			"Project Id: %s\n" +
			"Location: %s\n" +
			"Access Owner: %s\n" +
			"Access Owner Role: %s\n" +
			"Access Reader Role: %s\n" +
			"Access Reader Domain: %s\n" +
			"Expiration Time: %s\n" +
			"Dataset ID: %s\n" +
			"Table ID: %s\n" +
			"Time Partition Type: %s\n" +
			"Time Partition Expiration MS: %s\n" +
			"Time Partition Field: %s\n",
			inputDirectory,
			outputFile,
			projectId,
			location,
			accessOwner,
			accessOwnerRole,
			accessReaderRole,
			accessReaderDomain,
			expirationTime,
			datasetId,
			tableId,
			timePartitionType,
			timePartitionExpirationMs,
			timePartitionField)
	}
}

func (conf *BQConfiguration) IsValid() bool {
	return !(len(strings.Trim(conf.InputDirectory, Space)) ==0 ||
		len(strings.Trim(conf.OutputFile, Space)) == 0 ||
		len(strings.Trim(conf.AccessOwner, Space)) == 0)
}

func (conf *BQConfiguration) WriteFile() {

	tmpl, err := template.New(conf.TableId).Parse(BaseTemplate)

	if err != nil {
		log.Fatalf("failed to create from template: %v\n", err)
	}

	f, err := os.Create(conf.OutputFile)

	log.Printf("Writing to file: %v", f.Name())

	if err != nil {
		log.Fatalf("failed to create output file: %v\n", err)
	}

	// For each schema file in the directory
	err = filepath.Walk(conf.InputDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Failed to walk directory: %v", err)
		}
		if strings.HasSuffix(info.Name(), ".schema") {

			w := bufio.NewWriter(f)

			tableId = path[strings.LastIndex(path, "/")+1:strings.LastIndex(path, ".")]
			conf.TableId = tableId;
			conf.Schema = readFile(path)
			err = tmpl.Execute(w, conf)
			if err != nil {
				log.Fatalf("Failed to apply template to writer: %v\n", err)
			}
			w.Flush()
		}
		return nil
	})
	f.Close()
}

func varOrQuoted(value string) string {
	if strings.HasPrefix(value, "var.") {
		return value
	} else {
		return fmt.Sprintf("\"%s\"", value)
	}
}

func readFile(fileName string) string {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("Failed to read file: %s", fileName))
	}
	return string(dat)
}

func NewBQConfiguration() *BQConfiguration {
	return &BQConfiguration{
		InputDirectory: inputDirectory,
		OutputFile: outputFile,
		ProjectId: varOrQuoted(projectId),
		Location: varOrQuoted(location),
		AccessOwner: varOrQuoted(accessOwner),
		AccessOwnerRole: varOrQuoted(accessOwnerRole),
		AccessReaderRole: varOrQuoted(accessReaderRole),
		AccessReaderDomain: varOrQuoted(accessReaderDomain),
		ExpirationTime: varOrQuoted(expirationTime),
		DatasetId: datasetId,
		TableId: tableId,
		TimePartitionType: varOrQuoted(timePartitionType),
		TimePartitionExpirationMills: varOrQuoted(timePartitionExpirationMs),
		TimePartitionField: varOrQuoted(timePartitionField),
	}
}

const BaseTemplate = `
resource "google_bigquery_table" "{{.TableId}}" {
  dataset_id = google_bigquery_dataset.{{.DatasetId}}.dataset_id
  table_id   = "{{.TableId}}"
  time_partitioning {
    type = {{.TimePartitionType}}
  }
  labels = {
    provider = "google_cloud_retail"
  }
  schema = <<EOF
	{{.Schema}}
EOF

}`
