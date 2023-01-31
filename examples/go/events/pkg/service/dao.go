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

package service

import (
	"cloud.google.com/go/bigquery"
	"github.com/rrmcguinness/retail-data-model/retail/pkg/common"
)

type Dao struct {
	Module common.Module
	Table  *bigquery.Table
}

func NewDao(module common.Module, table *bigquery.Table) *Dao {
	return &Dao{Module: module, Table: table}
}
