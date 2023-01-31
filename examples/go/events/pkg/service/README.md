# DAO Design Pattern

Building DAOs have the following requirements:

1. Access to the underlying table or tables.
2. Access to GCP Project and Cloud Utilities.
3. Any additional helper classes.

## Common Imports

```go
import(
"cloud.google.com/go/bigquery"
"github.com/rrmcguinness/retail-data-services/gcp/bq"
"github.com/rrmcguinness/retail-data-services/gcp/meta"
"github.com/rrmcguinness/retail-data-services/gcp/utils"
"google.golang.org/api/iterator"
)
```
## Project Functions

```go
// Get the BigQuery Client
project.GetBigQueryClient()
// Execute a custom query
statement := "select ..."
project.GetBigQueryClient().Query(statement)
```

## Common Functions from Cloud Utilities

```go
// Get the table name to be used in query
cloud.QueryableNameFromQualifiedName(table.FullyQualifiedName())

// Format a time for Big Query
startTime := time.Now()
startTimeForQuery := startTime.UTC().Format(cloud.BigQueryDateFormat)

// Convert and inbound protocol buffer to a savable entity, this
// does the heavy lifting for deserializing proto.Message
// into a BigQuery struct.
record := pb.SomeProtobuf
newValueSaver := cloud.ConvertMessageToBigQueryType(record)

// Loading a record directly into a Protocol Buffer message
// uses a value loader
out := &pb.RecordType{}
loader := cloud.ConvertFromBigQueryToMessage(out)
// Row Iterator
err := itr.Next(loader)

// Use 'out' to either add to a list or stream.
 ```