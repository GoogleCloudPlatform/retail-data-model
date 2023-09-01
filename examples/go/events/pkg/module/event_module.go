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

package module

import (
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/retail-data-model/events/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/events/pkg/service"
	"github.com/GoogleCloudPlatform/retail-data-model/retail/pkg/common"
)

const (
	FlagSampleRate      = "sampleRate"
	TblEventDescription = "event_description"
	TblEventRecord      = "event_record"
)

type EventModule struct {
	BaseModule              common.Module
	EventDescriptionService *service.EventDescriptionsServerImpl
	EventRecordServer       *service.EventRecordsServerImpl
	EventServer             *service.EventsServerImpl
}

func (m *EventModule) GetConfig() *common.Config {
	return m.BaseModule.GetConfig()
}

func (m *EventModule) GetBigQuery() *common.BigQuery {
	return m.BaseModule.GetBigQuery()
}

func (m *EventModule) GetLogger() *common.Logger {
	return m.BaseModule.GetLogger()
}

func (m *EventModule) GetDefaultLog() *log.Logger {
	return m.BaseModule.GetDefaultLog()
}

func (m *EventModule) Close() {
	m.BaseModule.Close()
}

func (m *EventModule) bootstrap(sampleRate int32) {
	// Initialize Tables
	t1, err := m.GetBigQuery().InitializeTable(
		TblEventDescription,
		"",
		&pb.EventDescription{},
		m.GetConfig().Labels)

	HandleCriticalError(err)

	eventDescriptionDao, err := service.NewEventDescriptionDao(m, t1)

	HandleCriticalError(err)

	m.EventDescriptionService = service.NewEventDescriptionServer(eventDescriptionDao, m.GetDefaultLog())

	t2, err := m.GetBigQuery().InitializeTable(
		TblEventRecord,
		"",
		&pb.EventRecord{},
		m.GetConfig().Labels)

	HandleCriticalError(err)

	eventRecordDao := service.NewEventRecordDao(m, t2)
	m.EventRecordServer = service.NewEventRecordsServer(eventRecordDao)

	eventProcessor := service.NewEventProcessor(eventDescriptionDao, sampleRate, m.GetDefaultLog())

	// Initialize the Events Service
	m.EventServer = service.NewEventsServer(eventProcessor, eventRecordDao, m.GetDefaultLog())
}

func HandleCriticalError(err error) {
	if err != nil {
		panic(fmt.Sprintf("critical error loading the events common: %v", err))
	}
}
