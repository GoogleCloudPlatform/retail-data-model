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
	"fmt"

	"github.com/rrmcguinness/retail-data-model/common/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func StructToJson(in *proto.Message) (string, error) {
	b, e := protojson.Marshal(*in)
	return string(b), e
}

func GetStructFromJson(str string) map[string]interface{} {
	var input map[string]interface{}
	_ = json.Unmarshal([]byte(str), input)
	return input
}

func NewStatusResponse(id string, message string, responseType pb.StatusResponse_Type, in *proto.Message) *pb.StatusResponse {
	var payload *structpb.Struct
	if in != nil {
		p, e := StructToJson(in)
		if e == nil {
			payload, e = structpb.NewStruct(GetStructFromJson(p))
		}
	}

	return &pb.StatusResponse{
		Ts:      timestamppb.Now(),
		Type:    responseType,
		Id:      id,
		Message: message,
		Payload: payload,
	}
}

func NewSuccessMessage(id string, message string) *pb.StatusResponse {
	return NewStatusResponse(id, message, pb.StatusResponse_SUCCESS, nil)
}

func NewErrorMessage(id string, message error, in proto.Message) *pb.StatusResponse {
	errStr := fmt.Sprintf("%v", message)
	return NewStatusResponse(id, errStr, pb.StatusResponse_ERROR, &in)
}

func NewCreatedMessage(id string, message string) *pb.StatusResponse {
	return NewStatusResponse(id, message, pb.StatusResponse_CREATED, nil)
}

func NewUpdatedMessage(id string, message string) *pb.StatusResponse {
	return NewStatusResponse(id, message, pb.StatusResponse_UPDATED, nil)
}

func NewDeletedMessage(id string, message string) *pb.StatusResponse {
	return NewStatusResponse(id, message, pb.StatusResponse_DELETED, nil)
}
