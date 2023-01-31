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

package utils

import (
	"strings"
)

// ParseMapFromString is used for converting name/value pairs
// from command line arguments. For example adding tags to Cloud
// assets
func ParseMapFromString(in string) (out map[string]string) {
	out = make(map[string]string, 0)

	for _, kv := range strings.Split(in, ";") {
		kv1 := strings.Split(kv, "=")
		if len(kv1) == 2 {
			out[kv1[0]] = kv1[1]
		}
	}
	return out
}
