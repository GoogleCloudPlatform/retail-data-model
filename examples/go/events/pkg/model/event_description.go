package model

import (
	"log"
	"strings"

	"github.com/rrmcguinness/retail-data-model/events/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewEventDescription(record *pb.EventRecord) *pb.EventDescription {
	description := &pb.EventDescription{
		Name:                  strings.ToLower(record.Name),
		TotalObserved:         0,
		ChangeThresholdCommit: 0.33,
		SampleCounts:          make([]*pb.EventDescription_SampleCount, 0),
		TraitDescriptions:     make([]*pb.EventDescription_TraitDescription, 0),
		TraitChangeRecords:    make([]*pb.EventDescription_TraitChangeRecord, 0),
	}

	for _, trait := range record.Traits {
		description.TraitDescriptions = append(description.TraitDescriptions, NewEventDescriptionTraitDescription(trait))
	}
	return description
}

func NewEventDescriptionTraitDescription(trait *pb.EventRecord_Trait) *pb.EventDescription_TraitDescription {
	out := &pb.EventDescription_TraitDescription{
		Name:         strings.ToLower(trait.Name),
		UserOverride: false,
	}

	switch v := trait.Values.(type) {
	case *pb.EventRecord_Trait_Numeric:
		out.IsNumeric = true
		out.IsList = len(v.Numeric.Value) > 1
	case *pb.EventRecord_Trait_String_:
		out.IsNumeric = false
		out.IsList = len(v.String_.Value) > 1
	default:
		log.Print("Unhandled type")
	}
	return out
}

func GetTraitDescriptionByName(description *pb.EventDescription, traitName string) (out *pb.EventDescription_TraitDescription) {
	for _, traitDescription := range description.TraitDescriptions {
		if traitDescription.Name == traitName {
			out = traitDescription
			break
		}
	}
	return out
}

func DescriptionsAreEqual(d1 *pb.EventDescription, d2 *pb.EventDescription) bool {
	traitDescriptionsAreEqual := len(d1.TraitDescriptions) == len(d2.TraitDescriptions)
	if traitDescriptionsAreEqual {
		for _, t1 := range d1.TraitDescriptions {
			t2 := GetTraitDescriptionByName(d2, t1.Name)
			if t2 != nil && t1.Name == t2.Name {
				traitDescriptionsAreEqual =
					t1.IsNumeric == t2.IsNumeric &&
						t1.IsList == t2.IsList
			}
		}
	}
	return traitDescriptionsAreEqual
}

func CreateChangeRecord(old *pb.EventDescription, new *pb.EventDescription) {
	old.Name = new.Name
	old.Description = new.Description
	old.ChangeThresholdCommit = new.ChangeThresholdCommit
	old.TraitChangeRecords = append(old.TraitChangeRecords, &pb.EventDescription_TraitChangeRecord{
		Ts:                timestamppb.Now(),
		TraitDescriptions: old.GetTraitDescriptions()[:],
	})
	old.TraitDescriptions = new.TraitDescriptions[:]
}
