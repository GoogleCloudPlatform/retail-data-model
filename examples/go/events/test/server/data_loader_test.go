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

package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	model2 "github.com/rrmcguinness/retail-data-model/common/pkg/model"
	"github.com/rrmcguinness/retail-data-model/events/pb"
	clientPkg "github.com/rrmcguinness/retail-data-model/events/pkg/client"
	"github.com/rrmcguinness/retail-data-model/events/pkg/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const fileName = "/Users/rmcguinness/Downloads/event_input.json"

func handleMarshallError(err error) {
	fmt.Println("Error m")
}

const recordTimeFormat = "2006-01-02 15:04:05.999 UTC"

func getTimeFromBQ(value string) (ts *timestamppb.Timestamp) {
	oTime, err := time.Parse(recordTimeFormat, value)
	if err == nil {
		ts = timestamppb.New(oTime)
	} else {
		fmt.Printf("Faied to parse date: %s with err %v", value, err)
		ts = timestamppb.New(time.Now())
	}
	return ts
}

func handleCommonData(evt *pb.Event, o CommonData, productResponseMap string) {
	evt.Traits = append(evt.Traits, model.NewEventTrait("visitor_id", o.VisitorId))
	evt.Traits = append(evt.Traits, model.NewEventTrait("session_id", o.SessionId))
	var products = make([]string, len(o.ProductDetails))
	for i, product := range o.ProductDetails {
		bytes, err := json.Marshal(product)
		if err == nil {
			products[i] = string(bytes)
		} else {
			log.Default().Printf("Failed to marshal object to JSON: %v", err)
		}
	}
	evt.Traits = append(evt.Traits, model.NewEventTrait(productResponseMap, products...))
}

func ReadFromJson(line string, target interface{}) bool {
	err := json.Unmarshal([]byte(line), target)
	if err != nil {
		log.Default().Printf("Error Unmarshalling: %v", err)
		return false
	}
	return true
}

func Send(evt *pb.Event) {
	if err := clientPkg.Emit(context.Background(), client, nil, evt); err != nil {
		log.Default().Printf("Failed to persist: %v", err)
	}
}

func handleSearch(wg *sync.WaitGroup, line string) {
	wg.Add(1)
	defer wg.Done()
	var o = &Search{}
	if ok := ReadFromJson(line, o); ok {
		evt := &pb.Event{
			Created: getTimeFromBQ(o.EventTime),
			Name:    "web.search",
			Traits:  make([]*pb.Event_Trait, 0),
		}
		evt.Traits = append(evt.Traits, model.NewEventTrait("query", o.SearchQuery))
		handleCommonData(evt, o.CommonData, "search_results")
		Send(evt)
	}
}

func handleDetailPageView(wg *sync.WaitGroup, line string) {
	wg.Add(1)
	defer wg.Done()
	var o = &DetailPageView{}
	if ok := ReadFromJson(line, o); ok {
		evt := &pb.Event{
			Created: getTimeFromBQ(o.EventTime),
			Name:    "web.page.view",
			Traits:  make([]*pb.Event_Trait, 0),
		}
		evt.Traits = append(evt.Traits, model.NewEventTrait("page_name", "detail_page"))
		handleCommonData(evt, o.CommonData, "products")
		Send(evt)
	}
}

func handleAddToCart(wg *sync.WaitGroup, line string) {
	wg.Add(1)
	defer wg.Done()
	var o = &AddToCart{}
	if ok := ReadFromJson(line, o); ok {
		evt := &pb.Event{
			Created: getTimeFromBQ(o.EventTime),
			Name:    "web.action.cart.add",
			Traits:  make([]*pb.Event_Trait, 0),
		}
		handleCommonData(evt, o.CommonData, "products")
		Send(evt)
	}
}

func handlePurchaseComplete(wg *sync.WaitGroup, line string) {
	wg.Add(1)
	defer wg.Done()
	var o = &PurchaseComplete{}
	if ok := ReadFromJson(line, o); ok {
		evt := &pb.Event{
			Created: getTimeFromBQ(o.EventTime),
			Name:    "web.action.cart.tender",
			Traits:  make([]*pb.Event_Trait, 0),
		}
		handleCommonData(evt, o.CommonData, "products")
		evt.Traits = append(evt.Traits, model.NewEventTrait("transaction_id", o.PurchaseTransaction.Id))
		evt.Traits = append(evt.Traits, model.NewEventTrait("currency_code", o.PurchaseTransaction.CurrencyCode))
		evt.Traits = append(evt.Traits, model.NewEventTrait("total_tax", fmt.Sprintf("%f", o.PurchaseTransaction.Tax)))
		evt.Traits = append(evt.Traits, model.NewEventTrait("total_revenue", fmt.Sprintf("%f", o.PurchaseTransaction.Tax)))
		Send(evt)
	}
}

func ReadAllEvents() {
	fil, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	scanner := bufio.NewScanner(fil)
	scanner.Split(bufio.ScanLines)

	var wg sync.WaitGroup

	var latch = model2.NewLatch(32, "s")

	for scanner.Scan() {
		line := scanner.Text()
		ReadLine(latch, &wg, line)
	}
	wg.Wait()
	log.Default().Print("Complete")
}

func ReadLine(latch *model2.Latch, wg *sync.WaitGroup, line string) {
	if latch.IsClosed() {
		time.Sleep(latch.GetCurrentDuration())
		ReadLine(latch, wg, line)
	} else {
		if strings.Contains(line, BuildV1Discriminator(DiscriminatorV1Search)) {
			go handleSearch(wg, line)
		} else if strings.Contains(line, BuildV1Discriminator(DiscriminatorV1DetailPageView)) {
			go handleDetailPageView(wg, line)
		} else if strings.Contains(line, BuildV1Discriminator(DiscriminatorV1AddToCart)) {
			go handleAddToCart(wg, line)
		} else if strings.Contains(line, BuildV1Discriminator(DiscriminatorV1PurchaseComplete)) {
			go handlePurchaseComplete(wg, line)
		} else {
			log.Default().Println("No Match")
		}
	}
}

func TestDataPaser(t *testing.T) {
	ReadAllEvents()
}
