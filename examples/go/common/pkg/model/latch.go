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
	"time"
)

// Latch is a simple latch class that creates a time based latch
// which MAY participate in a WaitGroup.
// Note, when a Latch is Closed, that is considered the blocking state,
// whereas when a Latch is open, something may pass through.
type Latch struct {
	mu          sync.Mutex
	MaxDuration uint8
	Multiplier  uint8
	TimeSymbol  string
	Timer       *time.Timer
	countDown   *time.Timer
}

func NewLatch(maxDuration uint8, timeSymbol string) *Latch {
	if len(timeSymbol) == 0 {
		timeSymbol = "s"
	}

	return &Latch{
		MaxDuration: maxDuration,
		Multiplier:  0,
		TimeSymbol:  timeSymbol,
	}
}

func (l *Latch) IsClosed() bool {
	return l.Multiplier > 0
}

func (l *Latch) Open() {
	fmt.Println("Opening Latch")
	l.mu.Lock()
	l.Multiplier = 0
	l.Timer.Stop()
	l.mu.Unlock()
}

func (l *Latch) startCountDown() {
	d, _ := time.ParseDuration(fmt.Sprintf("1%s", l.TimeSymbol))
	if l.countDown == nil {
		l.countDown = time.NewTimer(d)
		go func() {
			<-l.countDown.C
			if l.Multiplier > 0 {
				l.mu.Lock()
				l.Multiplier -= 1
				l.mu.Unlock()
			} else {
				l.countDown.Stop()
			}
		}()
	} else {
		l.countDown.Reset(d)
	}
}

func (l *Latch) GetCurrentDuration() time.Duration {
	fmt.Printf("Starting timer %d%s\n", l.Multiplier, l.TimeSymbol)
	d, _ := time.ParseDuration(fmt.Sprintf("%d%s", l.Multiplier, l.TimeSymbol))
	return d
}

func (l *Latch) Close(wg *sync.WaitGroup) {
	l.mu.Lock()

	if l.Multiplier < l.MaxDuration {
		fmt.Println("Incrementing Multiplier")
		l.Multiplier += 1
	}

	if l.Timer == nil {
		l.Timer = time.NewTimer(l.GetCurrentDuration())
		if wg != nil {
			wg.Add(1)
		}
		go func() {
			<-l.Timer.C
			if wg != nil {
				wg.Done()
			}
			l.Open()
		}()
	} else {
		l.Timer.Reset(l.GetCurrentDuration())
	}
	l.mu.Unlock()
	l.startCountDown()
}
