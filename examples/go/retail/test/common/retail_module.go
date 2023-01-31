package common

import (
	"log"

	"github.com/rrmcguinness/retail-data-model/retail/pkg/common"
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
