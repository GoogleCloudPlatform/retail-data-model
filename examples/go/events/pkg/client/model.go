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
