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

package model

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestLatch_Close_Simple tests a simple latch closing with a 1-second timer for reset
func TestLatch_Close_Simple(t *testing.T) {

	l1 := NewLatch(2, "s")

	l1.Close(nil)
	assert.True(t, l1.IsClosed())

	time.Sleep(1 * time.Second)
	assert.False(t, l1.IsClosed())
}

// TestLatch_Close_Complex tests a more complex observer where a single
// start triggers the timer, then a second event increments the timer
// while the background watcher releases the latch. In addition, if a WaitGroup
// is passed in, it will become part of the WaitGroup blocker until released.
func TestLatch_Close_Complex(t *testing.T) {
	var wg sync.WaitGroup

	l1 := NewLatch(2, "s")
	l1.Close(&wg)
	fmt.Println("Starting Timer Watcher")
	assert.Equal(t, uint8(1), l1.Multiplier)

	assert.True(t, l1.IsClosed())
	l1.Close(&wg)
	assert.Equal(t, uint8(2), l1.Multiplier)
	time.Sleep(1 * time.Second)

	assert.True(t, l1.IsClosed())
	time.Sleep(2 * time.Second)
	assert.False(t, l1.IsClosed())
	assert.Equal(t, uint8(0), l1.Multiplier)

	wg.Wait()
}
