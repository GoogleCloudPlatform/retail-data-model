package client

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/rrmcguinness/retail-data-model/events/grpc"
	"github.com/rrmcguinness/retail-data-model/events/pb"
	"github.com/rrmcguinness/retail-data-model/events/pkg/model"
	grpcCore "google.golang.org/grpc"
)

func appendTrait(ctx context.Context, event *pb.Event, key string) {
	// Add the trace id if present
	v := ctx.Value(key)
	if v != nil {
		event.Traits = append(
			event.GetTraits(), model.NewEventTrait(KeyTraceId, fmt.Sprintf("%s", v)))
	}
}

func PrepareEvent(ctx context.Context, event *pb.Event) {
	if event != nil {
		appendTrait(ctx, event, KeyTraceId)
		appendTrait(ctx, event, KeyParentId)
		appendTrait(ctx, event, KeyTraceFlags)

		if len(event.EventId) == 0 {
			id, _ := uuid.NewRandom()
			event.EventId = id.String()
		}
		// Set the parent id to the event id if nil
		if len(event.EventParentId) == 0 {
			event.EventParentId = event.EventId
		}
	}
}

func Emit(ctx context.Context,
	client grpc.EventsClient,
	callback Callable,
	events ...*pb.Event) error {

	var opts []grpcCore.CallOption
	stream, err := client.Emit(ctx, opts...)
	if err != nil {
		return err
	}

	for _, event := range events {
		tE := event
		PrepareEvent(ctx, tE)
		if err := stream.Send(tE); err != nil {
			log.Default().Printf("failed to send event, %s, with error: %v", tE.EventId, err)
		}
	}

	// Once send is complete, wait for all responses, these may trigger
	// latch open events
	closeErr := stream.CloseSend()
	if closeErr != nil {
		log.Default().Printf("Error closing the send stream: %v", closeErr)
	}

	<-callback.GetWaitChannel()

	return nil
}
