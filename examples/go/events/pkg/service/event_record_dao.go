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
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/retail-data-model/events/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/retail/pkg/common"
	"google.golang.org/api/iterator"
)

// EventRecordDao is a data access object for saving events to BigQuery
type EventRecordDao struct {
	Dao
}

// NewEventRecordDao a constructor for the EventDao
func NewEventRecordDao(
	module common.Module,
	table *bigquery.Table) *EventRecordDao {

	return &EventRecordDao{
		Dao: Dao{module, table}}
}

// Create is the method responsible for converting events to event records and
// writing them to big query.
func (dao *EventRecordDao) Create(ctx context.Context, eventRecord *pb.EventRecord) error {
	return dao.Table.Inserter().Put(ctx, common.ConvertMessageToBigQueryType(eventRecord))
}

const StmtFindByDateRange = "SELECT * from `%s` WHERE emit_ts BETWEEN '%s' AND '%s'"

// FindByEmittedDateRange Finds a set of event records that were
// written during a range of time.
func (dao *EventRecordDao) FindByEmittedDateRange(ctx context.Context, startTime time.Time, endTime time.Time) (out []*pb.EventRecord, err error) {

	// Create a string that can be queried in BQ using UTC time
	queryString := fmt.Sprintf(
		StmtFindByDateRange,
		common.QueryableNameFromQualifiedName(dao.Table.FullyQualifiedName()),
		startTime.UTC().Format(common.BigQueryDateFormat),
		endTime.UTC().Format(common.BigQueryDateFormat))

	query := dao.Module.GetBigQuery().Query(queryString)

	out = make([]*pb.EventRecord, 0)

	var itr *bigquery.RowIterator
	itr, err = query.Read(ctx)
	if err == nil {
		for {

			var cEvent = &pb.EventRecord{}
			var valueLoader = common.ConvertFromBigQueryToMessage(cEvent)

			err = itr.Next(valueLoader)

			if err == iterator.Done {
				// Clear the error so it does not propagate
				err = nil
				break
			}

			if err != nil {
				dao.Module.GetDefaultLog().Printf("Error reading EventRecord from range; %v", err)
				break
			}

			if err == nil {
				out = append(out, cEvent)
			} else {
				dao.Module.GetDefaultLog().Printf("Failed to unmarshall EventRecord %v", err)
			}
		}
	} else {
		dao.Module.GetDefaultLog().Printf("Failed to read from query [[%s]] with error: %v", queryString, err)
	}
	return out, err
}
