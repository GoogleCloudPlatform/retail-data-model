---
title: "Using Google Retail Data Model with Go"
---

If you're using the Go Language, or have been wanting to explore cloud native
development using Go, there are two ways to get started:

1. Deploy a predefined implementation using Terraform directly to GCP.
2. Develop your own service layer, and take advantage of the definitions
   provided by the API.

## Option One

See the [Go Service Implementation Project]()

> NOTE: These projects are built using a Google Cloud account. Depending
> on the Terraform Script run various billable infrastructure will be provisioned.
> Please refer to each script to understand the implications.


## Option Two

Option one provides the absolute flexibility in development, but requires a
significant amount of investment to use the model. This option MAY be appealing
to large retailers that want to do complex extension on the existing model.

### Steps
1. Create a new Go project
2. Add a dependency on the Go libraries

```shell
% mkdir my_project
% cd my_project
% go mod init mycompany.com/my_project
% go get github.com/GoogleCloudPlatform/google-retail-cloud-api@go
```

### Creating Utilities

Utilities are useful in struct construction and normalizing behaviors
across packages.

#### Example Creation Pattern
```go
package model

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	"google.com/retail/events/pb"
	ts "google.golang.org/protobuf/types/known/timestamppb"
)

func NewEventRecordTrait(trait *pb.Event_Trait) *pb.EventRecord_Trait {
	// ... 
	return nil
}

// Converts an Event to an EventRecord
func NewEventRecord(event *pb.Event) *pb.EventRecord {
   recordId, _ := uuid.NewRandom()

   record := pb.EventRecord{
      TxId:          recordId.String(),
      EmitTs:        event.Created,
      ObserveTs:     ts.Now(),
      Name:          strings.ToLower(event.Name),
      EventId:       event.EventId,
      EventParentId: event.EventParentId,
      Traits:        make([]*pb.EventRecord_Trait, 0),
   }

   for _, trait := range event.Traits {
      record.Traits = append(record.Traits, NewEventRecordTrait(trait))
   }

   return &record
}

```


### Creating a service

Create the service implementation for a GRPC Endpoint (event_service.go)
```go
package service

import (
	"sync"

	"cloud.google.com/go/bigquery"
	"google.com/retail/events/grpc"
	"google.com/retail/events/pb"
	"google.com/retail/events/model"
)

type EventsServerImpl struct {
	grpc.UnimplementedEventsServer
	mu             sync.Mutex
	eventTable     *bigquery.Table
	eventProcessor model.EventProcessor
}

func NewEventsServer(eventTable *bigquery.Table) (grpc.EventsServer, error) {
	// Attempt to load event descriptors from the database
	return &EventsServerImpl{eventTable: eventTable}, nil
}

func (server *EventsServerImpl) Emit(stream grpc.Events_EmitServer) error {
	// Convert event from stream to an EventRecord and write to cloud storage
	return nil
}
```

