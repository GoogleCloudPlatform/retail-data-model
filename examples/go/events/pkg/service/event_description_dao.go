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
	"log"

	"github.com/GoogleCloudPlatform/retail-data-model/events/pkg/model"
	"github.com/GoogleCloudPlatform/retail-data-model/retail/pkg/common"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/retail-data-model/events/pb"
	"google.golang.org/api/iterator"
)

type EventDescriptionDao struct {
	Dao
}

// NewEventDescriptionDao constructor method
func NewEventDescriptionDao(module common.Module, table *bigquery.Table) (out *EventDescriptionDao, err error) {
	return &EventDescriptionDao{Dao{module, table}}, err
}

func (dao EventDescriptionDao) Create(ctx context.Context, in *pb.EventDescription) (out *pb.EventDescription, err error) {
	err = dao.Dao.Table.Inserter().Put(ctx, in)
	return in, err
}

func (dao EventDescriptionDao) Read(ctx context.Context, name string) (out *pb.EventDescription, err error) {
	itr, err := dao.Dao.Module.GetBigQuery().Query(fmt.Sprintf("select * from `%s` where name = '%s'", dao.Table.FullyQualifiedName(), name)).Read(ctx)
	for {
		var eventDescription pb.EventDescription
		err := itr.Next(&eventDescription)
		if err == iterator.Done {
			break
		} else {
			out = &eventDescription
		}
	}
	if out == nil {
		err = model.RecordNotFound
	}
	return out, err
}

func (dao EventDescriptionDao) Update(ctx context.Context, in *pb.EventDescription) (out *pb.EventDescription, err error) {
	err = dao.Table.Inserter().Put(ctx, in)
	return in, err
}

func (dao EventDescriptionDao) Delete(ctx context.Context, in *pb.EventDescription) (stats *bigquery.JobStatistics, err error) {
	query := dao.Module.GetBigQuery().Query(fmt.Sprintf("DELETE FROM `%s` WHERE name='%s", dao.Table.FullyQualifiedName(), in.Name))
	job, err := query.Run(ctx)
	if err == nil {
		status, err := job.Status(ctx)
		if err == nil {
			stats = status.Statistics
		}
	}
	return stats, err
}

// List returns a list of all EventDescription objects
func (dao EventDescriptionDao) List() (out []*pb.EventDescription) {
	out = make([]*pb.EventDescription, 0)
	ctx := context.Background()
	itr := dao.Table.Read(ctx)
	for {
		var eventDescription pb.EventDescription
		err := itr.Next(&eventDescription)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("failed to load event descritpion from table into the struct defintion")
		}
		out = append(out, &eventDescription)
	}
	return out
}
