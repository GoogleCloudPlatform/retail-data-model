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
	"io"
	"testing"

	pb2 "github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/events/grpc"
	"github.com/GoogleCloudPlatform/retail-data-model/events/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServerMock struct {
	grpc.EventsServer
	ctx            context.Context
	recvToServer   chan *pb.Event
	sentFromServer chan *pb2.StatusResponse
}

func (sm *ServerMock) Context() context.Context {
	return sm.ctx
}

func (sm *ServerMock) Send(resp *pb2.StatusResponse) error {
	sm.sentFromServer <- resp
	return nil
}

func (sm *ServerMock) Recv() (*pb.Event, error) {
	req, more := <-sm.recvToServer
	if !more {
		return nil, io.EOF
	}
	return req, nil
}

type MockClient struct {
	grpc.Events_EmitClient
}

func (sm *ServerMock) SendFromClient(event *pb.Event) error {
	sm.recvToServer <- event
	return nil
}

func (sm *ServerMock) RecvToClient() (*pb2.StatusResponse, error) {
	resp, more := <-sm.sentFromServer
	if !more {
		return nil, io.EOF
	}
	return resp, nil
}

func NewServerMock() *ServerMock {
	return &ServerMock{
		ctx:            context.Background(),
		recvToServer:   make(chan *pb.Event),
		sentFromServer: make(chan *pb2.StatusResponse),
	}
}

func TestClientCallback(t *testing.T) {
	server := NewServerMock()

	evt := &pb.Event{
		Created:       timestamppb.Now(),
		Name:          "test",
		EventId:       "test_event_01",
		EventParentId: "test_event_01",
		Traits:        make([]*pb.Event_Trait, 0),
	}

	server.SendFromClient(evt)

	server.Recv()

}
