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
	"github.com/rrmcguinness/retail-data-model/common/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewEmpty() *emptypb.Empty {
	return &emptypb.Empty{}
}

// NewCountry A factory method for creating Country objects
func NewCountry(name string,
		alpha2 string,
		alpha3 string,
		code string,
		iso2 string,
		region string,
		subRegion string,
		intermediateRegion string,
		regionCode string,
		subRegionCode string,
		intermediateRegionCode string) *pb.Country {
	return &pb.Country{Name: name,
		Alpha2:                 alpha2,
		Alpha3:                 alpha3,
		Code:                   code,
		Iso2:                   iso2,
		Region:                 region,
		SubRegion:              subRegion,
		IntermediateRegion:     intermediateRegion,
		RegionCode:             regionCode,
		SubRegionCode:          subRegionCode,
		IntermediateRegionCode: intermediateRegionCode}
}
