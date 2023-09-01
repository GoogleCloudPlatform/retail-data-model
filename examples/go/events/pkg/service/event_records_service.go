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

package service

import (
	"context"
	"sync"

	common "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/events/grpc"
	"github.com/GoogleCloudPlatform/retail-data-model/events/pb"
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
