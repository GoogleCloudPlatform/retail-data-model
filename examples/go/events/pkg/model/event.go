package model

import (
	"strings"

	"github.com/google/uuid"
	"github.com/rrmcguinness/retail-data-model/events/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func NewEvent(id string, parentId string, name string) *pb.Event {
	if len(id) == 0 {
		u, _ := uuid.NewRandom()
		id = u.String()
	}

	if len(parentId) == 0 {
		parentId = id
	}
	return &pb.Event{
		EventId:       id,
		EventParentId: parentId,
		Name:          strings.ToLower(name),
		Traits:        make([]*pb.Event_Trait, 0),
	}
}

func NewEventTrait(name string, values ...string) *pb.Event_Trait {
	return &pb.Event_Trait{
		Name:   name,
		Values: values,
	}
}

func GetTraitByName(event *pb.Event, traitName string) (out *pb.Event_Trait) {
	for _, trait := range event.Traits {
		if trait.Name == traitName {
			out = trait
			break
		}
	}
	return out
}

func EventToMessage(event *pb.Event) (proto.Message, error) {
	b, e := protojson.Marshal(event)
	if e != nil {
		var out proto.Message
		e = protojson.Unmarshal(b, out)
		if e == nil {
			return out, e
		}
	}
	return nil, e
}
