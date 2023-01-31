package client

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/rrmcguinness/retail-data-model/common/grpc"
	"github.com/rrmcguinness/retail-data-model/common/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuditServer struct {
	grpc.UnimplementedAuditRecordsServer
	mu sync.Mutex
}

func (srv AuditServer) Create(ctx context.Context, auditRecord pb.AuditRecord) (out pb.StatusResponse, err error) {
	uuid, err := uuid.NewRandom()
	auditRecord.Id = uuid.String()
	auditRecord.Created = timestamppb.Now()
	return out, err
}

func (srv AuditServer) Search(ctx context.Context, criteria pb.AuditRecordSearchRequest, client grpc.AuditRecords_SearchClient) (err error) {
	crit := criteria.Criteria
	bounds := criteria.Bounds

	hasCriteria := crit != nil
	isTimeBound := bounds != nil && bounds.StartDate.AsTime().Before(bounds.EndDate.AsTime())

	if hasCriteria && len(crit.Id) > 0 {
		// TODO - Attempt a read from ID first
	} else if hasCriteria || isTimeBound {
		// TODO - Build Criteria Object
	} else {
		err = errors.New("invalid request")
	}
	return err
}
