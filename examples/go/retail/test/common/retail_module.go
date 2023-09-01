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
	"log"

	"github.com/GoogleCloudPlatform/retail-data-model/retail/pkg/common"
)

type RetailModule struct {
	mod common.Module
}

func (retail *RetailModule) GetBigQuery() *common.BigQuery {
	return retail.mod.GetBigQuery()
}

func (retail *RetailModule) GetLabels() map[string]string {
	return retail.mod.GetConfig().Labels
}

func (retail *RetailModule) GetLogger() *common.Logger {
	return retail.mod.GetLogger()
}

func (retail *RetailModule) GetLog() *log.Logger {
	return retail.mod.GetDefaultLog()
}

func (retail *RetailModule) Close() {
	retail.mod.Close()
}

func NewRetailModule(configName string) (mod *RetailModule, err error) {
	// Load the common from a config file
	baseModule, err := common.LoadBaseModule(configName)

	// If not in error, create the retail common
	if err == nil && baseModule != nil {
		mod = &RetailModule{mod: baseModule}
	}
	return mod, err
}
