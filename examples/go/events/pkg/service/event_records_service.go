package service

import (
	"context"
	"sync"

	common "github.com/rrmcguinness/retail-data-model/common/pb"
	"github.com/rrmcguinness/retail-data-model/events/grpc"
	"github.com/rrmcguinness/retail-data-model/events/pb"
)

type EventRecordsServerImpl struct {
	grpc.UnimplementedEventRecordsServer
	mu  sync.Mutex
	dao *EventRecordDao
}

func NewEventRecordsServer(dao *EventRecordDao) *EventRecordsServerImpl {
	return &EventRecordsServerImpl{dao: dao}
}

func (server *EventRecordsServerImpl) List(in *common.TimeBoundRequest, stream grpc.EventRecords_ListServer) (err error) {

	return err
}

func (server *EventRecordsServerImpl) FindTransactionById(ctx context.Context, request *common.IDRequest) (out *pb.EventRecord, err error) {

	return out, err
}
