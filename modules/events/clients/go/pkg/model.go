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
	"io"

	"github.com/google/uuid"
	pb2 "github.com/rrmcguinness/retail-data-model/common/pb"
	"github.com/rrmcguinness/retail-data-model/common/pkg/model"
	"github.com/rrmcguinness/retail-data-model/events/grpc"
)

type Callable interface {
	Callable()
	GetWaitChannel() chan struct{}
	GetCommChannel() chan *pb2.StatusResponse
}

type EventClientCallback struct {
	waitc  chan struct{}
	stream grpc.Events_EmitClient
	comm   chan *pb2.StatusResponse
}

func NewEventClientCallback(stream grpc.Events_EmitClient) *EventClientCallback {
	return &EventClientCallback{
		waitc:  make(chan struct{}),
		stream: stream,
		comm:   make(chan *pb2.StatusResponse),
	}
}

func (e EventClientCallback) GetCommChannel() chan *pb2.StatusResponse {
	return e.comm
}

func (e EventClientCallback) SetWaitChannel(in chan struct{}) {
	e.waitc = in
}

func (e EventClientCallback) GetWaitChannel() chan struct{} {
	return e.waitc
}

func (e EventClientCallback) RegisterCommunicationsChannel(comm chan *pb2.StatusResponse) {
	e.comm = comm
}

func (e EventClientCallback) Callback() {
	for {
		in, err := e.stream.Recv()
		if err == io.EOF {
			close(e.waitc)
			return
		}
		if err != nil {
			id, _ := uuid.NewRandom()
			e.comm <- model.NewErrorMessage(id.String(), err, nil)
			return
		}
		e.comm <- in
	}
}

func (e EventClientCallback) close() {
	close(e.comm)
}

type EventClient struct {
	callbackHandler EventClientCallback
	client          grpc.EventsClient
}
