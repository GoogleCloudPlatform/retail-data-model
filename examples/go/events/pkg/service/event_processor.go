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
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/madflojo/tasks"
	"github.com/GoogleCloudPlatform/retail-data-model/events/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/events/pkg/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	MinSampleRate int32 = 10000
)

type EventProcessor struct {
	mu                  sync.Mutex
	eventDescriptionDao *EventDescriptionDao
	state               map[string]*pb.EventDescription
	sampleRate          int32
	sampleCount         int64
	scheduler           *tasks.Scheduler
	runningTasks        []string
	log                 *log.Logger
}

func NewEventProcessor(
	eventDescriptionDao *EventDescriptionDao,
	sampleRate int32,
	log *log.Logger) *EventProcessor {
	// Ensure some sanity
	if sampleRate < MinSampleRate {
		log.Printf("The minimum sample rate is %d, the attempt was made to set it to a value of: %d was ignored.\n", MinSampleRate, sampleRate)
		sampleRate = MinSampleRate
	}
	processor := &EventProcessor{
		eventDescriptionDao: eventDescriptionDao,
		state:               make(map[string]*pb.EventDescription, 0),
		sampleRate:          sampleRate,
		sampleCount:         0,
		scheduler:           tasks.New(),
		runningTasks:        make([]string, 0),
		log:                 log}

	processor.UpdateState()

	return processor
}

func (processor *EventProcessor) InitializeScheduler() {
	id, err := processor.scheduler.Add(
		&tasks.Task{
			Interval: 5 * time.Minute,
			TaskFunc: func() error {
				processor.UpdateState()
				return nil
			},
			ErrFunc: func(err error) {
				log.Printf("An error occurred while running event processor update: %v", err)
			},
		})
	if err != nil {
		log.Printf("error setting up task scheduler: %v", err)
	} else {
		processor.runningTasks = append(processor.runningTasks, id)
		log.Printf("started schedule for updating state: %s", id)
	}
}

func (processor *EventProcessor) ContainsDescription(name string) bool {
	ans := processor.state[name]
	if ans != nil {
		return true
	}
	return false
}

func (processor *EventProcessor) GetDescription(name string) *pb.EventDescription {
	return processor.state[name]
}

func (processor *EventProcessor) IncrementSampleCount() {
	atomic.AddInt64(&processor.sampleCount, 1)
}

func (processor *EventProcessor) DoSample() bool {
	return processor.sampleRate%processor.sampleRate == 0
}

func (processor *EventProcessor) DefinitionCount() int {
	return len(processor.state)
}

func (processor *EventProcessor) UpdateState() {
	processor.mu.Lock()
	processor.log.Println("Updating processor definitions state")
	count := 0
	for _, e := range processor.eventDescriptionDao.List() {
		processor.state[e.Name] = e
		count = count + 1
	}
	processor.mu.Unlock()
	processor.log.Printf("Processor state updated %d records processed", count)
}

func (processor *EventProcessor) Process(event *pb.Event) (record *pb.EventRecord, eventDescription *pb.EventDescription, err error) {
	processor.IncrementSampleCount()
	processor.mu.Lock()
	defer processor.mu.Unlock()
	record, err = model.NewEventRecord(event)
	if err == nil {
		if !processor.ContainsDescription(record.Name) {
			eventDescription = model.NewEventDescription(record)
			processor.state[eventDescription.Name] = eventDescription
			processor.eventDescriptionDao.Create(context.Background(), eventDescription)
		} else {
			eventDescription = processor.GetDescription(record.Name)
			if eventDescription != nil {
				if eventDescription.TotalObserved%int64(processor.sampleRate) == 0 {
					currentObservation := model.NewEventDescription(record)
					areEqual := model.DescriptionsAreEqual(eventDescription, currentObservation)
					if !areEqual {
						//log.Print("Event Descriptions are Not Equal !=")
						model.CreateChangeRecord(eventDescription, currentObservation)
						_, _ = processor.eventDescriptionDao.Update(context.Background(), currentObservation)
						// Update other counts
					} else {
						//log.Print("Event Descriptions are Equal ==")
						//Update sample count
						eventDescription.SampleCounts = append(eventDescription.SampleCounts, &pb.EventDescription_SampleCount{
							Ts:          timestamppb.Now(),
							ProducerId:  "event_service",
							SampleCount: currentObservation.TotalObserved,
							ChangeCount: 0,
						})
						processor.eventDescriptionDao.Create(context.Background(), eventDescription)
					}
				}
			}
		}
		if eventDescription != nil {
			eventDescription.TotalObserved += 1
		}
	}
	return record, eventDescription, err
}

func (processor *EventProcessor) Close() {
	// Remove and close all background tasks
	for _, id := range processor.runningTasks {
		processor.scheduler.Del(id)
	}
}
