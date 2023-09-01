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

package module

import (
	"flag"
	"path/filepath"

	"github.com/GoogleCloudPlatform/retail-data-model/retail/pkg/common"
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
