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
	"os"
	"testing"

	"github.com/rrmcguinness/retail-data-model/events/pkg/module"
)

var eventModule *module.EventModule

const (
	ConfigFile = "-configFile=examples/configs/events-test-config.json"
	SampleRate = "-sampleRate=1"
)

func setup() {
	var err error
	os.Args = append(os.Args, ConfigFile)
	os.Args = append(os.Args, SampleRate)

	// Parse all defined flags
	flag.Parse()

	// Load the event module
	eventModule, err = module.Load()
	if err != nil {
		panic(err)
	}
}

func teardown() {
	if eventModule != nil {
		eventModule.Close()
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
