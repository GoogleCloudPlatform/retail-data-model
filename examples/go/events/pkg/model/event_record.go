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

package model

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/GoogleCloudPlatform/retail-data-model/events/pb"
	ts "google.golang.org/protobuf/types/known/timestamppb"
)

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func GetStructFromJson(str string) map[string]interface{} {
	var input map[string]interface{}
	_ = json.Unmarshal([]byte(str), input)
	return input
}

func NewEventRecordTrait(trait *pb.Event_Trait) (*pb.EventRecord_Trait, error) {
	numericValues := make([]float64, len(trait.Values))
	var intCount = 0
	for _, value := range trait.Values {
		if fValue, err := strconv.ParseFloat(value, 64); err == nil {
			numericValues[intCount] = fValue
			intCount = intCount + 1
		}
	}

	//if intCount == len(trait.Values) {
	//	return &pb.EventRecord_Trait{
	//		Name: strings.ToLower(trait.Name),
	//		Values: &pb.EventRecord_Trait_Numeric{
	//			Numeric: &pb.EventRecord_Trait_Number{
	//				Value: numericValues[:],
	//			},
	//		},
	//	}, nil
	//} else {
	//	if len(trait.Values) == 1 && IsJSON(trait.Values[0]) {
	//		out, err := structpb.NewStruct(GetStructFromJson(trait.Values[0]))
	//		if err == nil {
	//			return &pb.EventRecord_Trait{
	//				Name: "",
	//				Values: &pb.EventRecord_Trait_Object{
	//					Object: out,
	//				},
	//			}, nil
	//		}
	//	} else {
	//		return &pb.EventRecord_Trait{
	//			Name: strings.ToLower(trait.Name),
	//			Values: &pb.EventRecord_Trait_String_{
	//				String_: &pb.EventRecord_Trait_String{
	//					Value: trait.Values[:],
	//				},
	//			},
	//		}, nil
	//	}
	//}
	return nil, ConversionError
}

func NewEventRecord(event *pb.Event) (*pb.EventRecord, error) {
	if event != nil {
		recordId, _ := uuid.NewRandom()

		record := &pb.EventRecord{
			TxId:          recordId.String(),
			EmitTs:        event.Created,
			ObserveTs:     ts.Now(),
			Name:          strings.ToLower(event.Name),
			EventId:       event.EventId,
			EventParentId: event.EventParentId,
			Traits:        make([]*pb.EventRecord_Trait, 0),
		}

		for _, trait := range event.Traits {
			traitRecord, err := NewEventRecordTrait(trait)
			if err == nil {
				record.Traits = append(record.Traits, traitRecord)
			} else {
				return nil, err
			}
		}

		return record, nil
	}
	return nil, errors.New("invalid event")
}
