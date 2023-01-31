package module

import (
	"flag"
	"path/filepath"

	"github.com/rrmcguinness/retail-data-model/retail/pkg/common"
)

var (
	sampleRate = flag.Int(FlagSampleRate, 100, "Defines the rate for sampling events, for example if the value is 100, 1 in every 100 events will be sampled.")
	configFile = flag.String("configFile", "", "The relative or exact path of a configuration file.")
)

func Load() (*EventModule, error) {
	pth, err := filepath.Abs(*configFile)
	config, err := common.LoadConfig(pth)

	if err == nil {
		defaultModule, err := common.NewBaseModule(config)
		if err == nil {
			eventModule := &EventModule{BaseModule: defaultModule}
			eventModule.bootstrap(int32(*sampleRate))
			return eventModule, nil
		}
	}
	return nil, err
}
