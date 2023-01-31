package service

import (
	"context"
	"errors"
	"sort"
	"sync"
	"time"

	common "github.com/rrmcguinness/retail-data-model/common/pb"
	"github.com/rrmcguinness/retail-data-model/enterprise/grpc"
	"github.com/rrmcguinness/retail-data-model/enterprise/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type VersionKey struct {
	key     string
	version int32
}

type TaxCodeServer struct {
	grpc.UnimplementedTaxCodesServer
	savedTaxCodes map[*VersionKey]*pb.TaxCode
	mu            sync.Mutex
}

func (srv *TaxCodeServer) GetTaxCodes(value *emptypb.Empty, stream grpc.TaxCodes_GetTaxCodesServer) (err error) {
	for _, taxCode := range srv.savedTaxCodes {
		err = stream.Send(taxCode)
		if err != nil {
			break
		}
	}
	return err
}

func (srv *TaxCodeServer) getTaxCodeById(id string) (out []*pb.TaxCode, err error) {
	list := make([]*pb.TaxCode, 0)
	for k, v := range srv.savedTaxCodes {
		if k.key == id {
			list = append(list, v)
		}
	}

	if len(list) == 0 {
		return list, errors.New("key not present")
	}

	srv.sortTaxCodesByDate(list)

	return list, nil
}

func (srv *TaxCodeServer) sortTaxCodesByDate(list []*pb.TaxCode) {
	sort.Slice(list[:], func(i, j int) bool {
		return list[i].Id.Effective.Seconds < list[j].Id.Effective.Seconds
	})
}

// GetTaxCode returns the latest active TaxCode
func (srv *TaxCodeServer) GetTaxCode(ctx context.Context, value *common.IDRequest) (out *pb.TaxCode, err error) {
	list, err := srv.getTaxCodeById(value.Id)
	if err == nil {
		for _, taxCode := range list {
			if taxCode.Id.Effective.AsTime().Before(time.Now()) && taxCode.Deleted == nil {
				return taxCode, nil
			}
		}
	}
	return nil, errors.New("invalid ID")
}

func (srv *TaxCodeServer) GetTaxCodeVersionHistory(value *common.IDRequest, stream grpc.TaxCodes_GetTaxCodeVersionHistoryServer) (err error) {
	list, err := srv.getTaxCodeById(value.Id)
	if err != nil {
		return err
	}
	for _, taxCode := range list {
		stream.Send(taxCode)
	}
	return nil
}

func (srv *TaxCodeServer) GetTaxCodeVersion(ctx context.Context, request *common.VersionIDRequest) (out *pb.TaxCode, err error) {
	vKey := &VersionKey{key: request.Id, version: request.Version}
	out = srv.savedTaxCodes[vKey]
	if out == nil {
		return out, errors.New("invalid ID")
	}
	return out, nil
}

/* Tax Code Methods */

// Create - creates a new tax code in the data backing store
func (srv *TaxCodeServer) Create(ctx context.Context, request *pb.TaxCode) (out *pb.TaxCode, err error) {

	return out, err
}

// Update - updates an existing tax code based on its id, incrementing the version if change is detected
func (srv *TaxCodeServer) Update(ctx context.Context, request *pb.TaxCode) (out *pb.TaxCode, err error) {
	return out, err
}

// Delete - removes a tax code, and creates an audit record
func (srv *TaxCodeServer) Delete(ctx context.Context, request *pb.TaxCode) (out *common.StatusResponse, err error) {
	return out, err
}

/* Rate Methods */

func (srv *TaxCodeServer) CreateRate(ctx context.Context, request *pb.TaxCode_Rate) (out *pb.TaxCode_Rate, err error) {
	return out, err
}

func (srv *TaxCodeServer) UpdateRate(ctx context.Context, request *pb.TaxCode_Rate) (out *pb.TaxCode_Rate, err error) {
	return out, err
}

func (srv *TaxCodeServer) DeleteRate(ctx context.Context, request *pb.TaxCode_Rate) (out *common.StatusResponse, err error) {
	return out, err
}
