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

import (
	"context"
	"io"
	"log"
	"time"

	"cloud.google.com/rdm/common/grpc"
	"cloud.google.com/rdm/common/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListCountries provides a list of countries from the server
func ListCountries(client grpc.CountriesClient) ([]*pb.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	empty := &emptypb.Empty{}

	stream, err := client.List(ctx, empty)

	out := make([]*pb.Country, 0)

	for {
		country, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Country list failed to recieve %v", err)
		}
		out = append(out, country)
	}

	return out, err
}

// Create creates a country on the server
func Create(client grpc.CountriesClient, country *pb.Country) (*pb.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	out, err := client.Create(ctx, country)
	if err != nil {
		log.Fatalf("Failed to create country")
	}
	return out, err
}

// Update updates a country on the server
func Update(client grpc.CountriesClient, country *pb.Country) (*pb.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	out, err := client.Update(ctx, country)
	if err != nil {
		log.Fatalf("Failed to create country")
	}
	return out, err
}

// Delete removed a country definition from the server
func Delete(client grpc.CountriesClient, country *pb.Country) (*pb.StatusResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	out, err := client.Delete(ctx, country)
	if err != nil {
		log.Fatalf("Failed to create country")
	}
	return out, err
}
