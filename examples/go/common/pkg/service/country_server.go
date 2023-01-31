package service

// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	context "context"
	sync "sync"

	"github.com/google/uuid"
	"github.com/rrmcguinness/retail-data-model/common/grpc"
	"github.com/rrmcguinness/retail-data-model/common/pb"
	"github.com/rrmcguinness/retail-data-model/common/pkg/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CountryServer Is a server that will listen on a random available port
// specifically used for testing GRPC functions.
type CountryServer struct {
	grpc.UnimplementedCountriesServer
	savedCountries []*pb.Country
	mu             sync.Mutex
}

// List pulls a list of countries from the saved country list
func (cs *CountryServer) List(value *emptypb.Empty, stream grpc.Countries_ListServer) (err error) {
	for _, country := range cs.savedCountries {
		err = stream.Send(country)
		if err != nil {
			break
		}
	}
	return err
}

// Create adds a country to the saved country list
func (cs *CountryServer) Create(ctx context.Context, country *pb.Country) (*pb.Country, error) {
	if len(country.Id) == 0 {
		id, _ := uuid.NewRandom()
		country.Id = id.String()
	}
	cs.savedCountries = append(cs.savedCountries, country)
	return country, nil
}

// Update updates an existing country if it exists in the map, otherwise it adds it.
func (cs *CountryServer) Update(ctx context.Context, country *pb.Country) (*pb.Country, error) {
	cs.Delete(ctx, country)
	cs.savedCountries = append(cs.savedCountries, country)
	return country, nil
}

// Delete removes a country from the saved countries list.
func (cs *CountryServer) Delete(ctx context.Context, country *pb.Country) (*pb.StatusResponse, error) {
	newArray := make([]*pb.Country, 0)

	affectedIds := make([]string, 0)
	for _, ctry := range cs.savedCountries {
		if ctry.Id != country.Id {
			newArray = append(newArray, ctry)
		}
		affectedIds = append(affectedIds, ctry.Id)
	}
	cs.savedCountries = newArray

	return model.NewDeletedMessage(country.Id, ""), nil
}
