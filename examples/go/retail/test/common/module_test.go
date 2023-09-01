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

package common

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/retail-data-model/common/pb"
	"github.com/GoogleCloudPlatform/retail-data-model/retail/pkg/common"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// WriteToTable illustrates the Write function, note the conversion from contact (protobuf)
// into a proto mapper for BQ.
func WriteToTable(r *RetailModule, table *bigquery.Table, t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	snContacts := make([]*pb.Contact_SocialMediaNetworkIdentity, 0)

	snContacts = append(snContacts,
		&pb.Contact_SocialMediaNetworkIdentity{Platform: "TWITTER", PlatformId: "@googlecloud"})

	snContacts = append(snContacts,
		&pb.Contact_SocialMediaNetworkIdentity{Platform: "LINKEDIN", PlatformId: "google_cloud"})

	snContacts = append(snContacts,
		&pb.Contact_SocialMediaNetworkIdentity{Platform: "GITHUB", PlatformId: "https://github.com/googleapis"})

	contact := &pb.Contact{
		Id:              "TEST_INVENTORY2",
		ContactPurpose:  "PRESALES",
		ContactMethod:   pb.Contact_BUSINESS,
		EffectiveDate:   timestamppb.Now(),
		EmailAddress:    "tester_grdm@google.com",
		Telephone:       "303-555-1212",
		Website:         "https://cloud.google.com",
		SocialNetworkId: snContacts,
		Status:          "ACTIVE",
	}

	pContact := common.ConvertMessageToBigQueryType(contact)

	err := table.Inserter().Put(ctx, pContact)

	assert.Nil(t, err)

}

func ReadTable(r *RetailModule, table *bigquery.Table, t *testing.T) *pb.Contact {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	tableName := common.QueryableNameFromQualifiedName(table.FullyQualifiedName())
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE id = 'TEST_INVENTORY2'", tableName)

	itr, err := r.GetBigQuery().Query(query).Read(ctx)

	var existing *pb.Contact

	if err == nil {
		for {
			cContact := &pb.Contact{}
			valueLoader := common.ConvertFromBigQueryToMessage(cContact)
			err := itr.Next(valueLoader)
			if err == iterator.Done {
				break
			}
			if cContact != nil {
				existing = cContact
			}
		}

		if existing == nil {
			WriteToTable(r, table, t)
			existing = ReadTable(r, table, t)
		}

		assert.NotNil(t, existing)
		assert.Equal(t, "TEST_INVENTORY2", existing.Id)
	} else {
		assert.Fail(t, "failed to read table", err)
	}
	return existing
}

const (
	ConfigFileName = "examples/configs/retail-test-config.json"
)

func TestRetailModule(t *testing.T) {
	pth, err := filepath.Abs(ConfigFileName)
	assert.Nil(t, err)

	retail, err := NewRetailModule(pth)

	assert.Nil(t, err)
	assert.NotNil(t, retail)

	defer retail.Close()

	// Load Data Structures

	contactTable, err := retail.GetBigQuery().InitializeTable(
		"contacts",
		"",
		&pb.Contact{},
		retail.GetLabels())

	retail.GetLog().Printf("Initialized Table: %s", contactTable.TableID)

	contact := ReadTable(retail, contactTable, t)

	assert.NotNil(t, contact)
}
