package service

import (
	"context"
	"io"
	"log"
	"sync"

	common "github.com/rrmcguinness/retail-data-model/common/pkg/model"
	"github.com/rrmcguinness/retail-data-model/events/grpc"
	"github.com/rrmcguinness/retail-data-model/events/pkg/model"
)

type EventsServerImpl struct {
	grpc.UnimplementedEventsServer
	mu             sync.Mutex
	eventProcessor *EventProcessor
	eventRecordDao *EventRecordDao
	log            *log.Logger
}

func NewEventsServer(
	eventProcessor *EventProcessor,
	eventRecordDao *EventRecordDao,
	log *log.Logger) *EventsServerImpl {

	return &EventsServerImpl{
		eventProcessor: eventProcessor,
		eventRecordDao: eventRecordDao,
		log:            log}
}

// Emit is used by client APIs to record events on the server.
// The event processor is responsible for streamlining event evaluation
// and description evaluation
func (server *EventsServerImpl) Emit(stream grpc.Events_EmitServer) error {
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			return io.EOF
		}
		if err != nil {
			return err
		}
		record, eventDescription, processErr := server.eventProcessor.Process(event)
		eventDescription.TotalObserved += 1
		if processErr != nil {
			log.Default().Printf("An error occurred processing an event: %v", processErr)
		} else {
			if createErr := server.eventRecordDao.Create(context.Background(), record); createErr != nil {
				if event != nil {
					message, convertErr := model.EventToMessage(event)
					if convertErr == nil {
						if sendErr := stream.Send(common.NewErrorMessage(event.EventId, createErr, message)); sendErr != nil {
							log.Printf("An error occurred while sending ane error response: %v", err)
						}
					}
				}
			}
			if sendErr := stream.Send(common.NewCreatedMessage(event.EventId, "")); sendErr != nil {
				return sendErr
			}
		}
		return nil
	}
}
