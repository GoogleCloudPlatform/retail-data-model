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
