/*
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
