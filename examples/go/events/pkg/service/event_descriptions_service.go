package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	commonProto "github.com/rrmcguinness/retail-data-model/common/pb"
	common "github.com/rrmcguinness/retail-data-model/common/pkg/model"
	"github.com/rrmcguinness/retail-data-model/events/grpc"
	eventProto "github.com/rrmcguinness/retail-data-model/events/pb"
	"github.com/rrmcguinness/retail-data-model/events/pkg/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EventDescriptionsServerImpl struct {
	grpc.UnimplementedEventsServer
	mu  sync.Mutex
	log *log.Logger
	dao *EventDescriptionDao
}

func NewEventDescriptionServer(dao *EventDescriptionDao, log *log.Logger) *EventDescriptionsServerImpl {
	return &EventDescriptionsServerImpl{dao: dao, log: log}
}

func (server *EventDescriptionsServerImpl) List(in *emptypb.Empty, stream grpc.EventDescriptions_ListServer) (err error) {
	if server.dao != nil {
		for _, description := range server.dao.List() {
			err = stream.Send(description)
		}
	}
	return err
}

func (server *EventDescriptionsServerImpl) Get(ctx context.Context, in *commonProto.IDRequest) (out *eventProto.EventDescription, err error) {
	return server.dao.Read(ctx, in.Id)
}

func (server *EventDescriptionsServerImpl) Create(stream grpc.EventDescriptions_CreateServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		existing, readErr := server.dao.Read(ctx, in.Name)
		if readErr != nil {
			_ = stream.Send(common.NewErrorMessage(in.Name, model.RecordNotChanged, in))
		} else {
			if existing != nil {
				_ = stream.Send(common.NewErrorMessage(in.Name, model.RecordIdExists, in))
			} else {
				out, createErr := server.dao.Create(ctx, in)
				if createErr == nil {
					_ = stream.Send(common.NewCreatedMessage(out.Name, ""))
				} else {
					server.log.Printf("failed to persist: %v", in)
				}
			}
		}
	}
	return nil
}

func (server *EventDescriptionsServerImpl) Update(stream grpc.EventDescriptions_UpdateServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}

		exiting, readErr := server.dao.Read(ctx, in.Name)

		if readErr == nil {
			if !model.DescriptionsAreEqual(exiting, in) {
				model.CreateChangeRecord(exiting, in)
				d, e := server.dao.Update(ctx, exiting)
				if e != nil {
					err = stream.Send(common.NewUpdatedMessage(d.Name, ""))
				}
			} else {
				err = stream.Send(common.NewErrorMessage(exiting.Name, model.RecordNotChanged, in))
			}
		}
		if err != nil {
			server.log.Printf("failed to update record: %s from %v", in.Name, in)
		}
	}
	return nil
}

func (server *EventDescriptionsServerImpl) Delete(stream grpc.EventDescriptions_DeleteServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for {
		in, receiveErr := stream.Recv()
		if receiveErr == io.EOF {
			break
		}

		stats, deleteErr := server.dao.Delete(ctx, in)
		if deleteErr != nil {
			return deleteErr
		}

		var txId string
		if stats != nil {
			txId = stats.TransactionInfo.TransactionID
		}
		if err := stream.SendMsg(common.NewDeletedMessage(in.Name, fmt.Sprintf("Transaction Id: %s", txId))); err != nil {
			log.Default().Printf("Failed to send message: %v", err)
		}
	}
	return nil
}
