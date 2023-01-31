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

import "fmt"

const (
	DiscriminatorPattern            = "event_type\":\"%s\""
	DiscriminatorV1Search           = "search"
	DiscriminatorV1DetailPageView   = "detail-page-view"
	DiscriminatorV1AddToCart        = "add-to-cart"
	DiscriminatorV1PurchaseComplete = "purchase-complete"
)

func BuildV1Discriminator(value string) string {
	return fmt.Sprintf(DiscriminatorPattern, value)
}

type Product struct {
	Id        string `json:"id,omitempty"`
	PriceInfo struct {
		Price        float32 `json:"price,omitempty"`
		CurrencyCode string  `json:"currency_code,omitempty"`
	} `json:"price_info,omitempty"`
}

type ProductDetail struct {
	Quantity int32   `json:"quantity,omitempty,string"`
	Product  Product `json:"product,omitempty"`
}

type CommonData struct {
	EventType      string          `json:"event_type"`
	EventTime      string          `json:"event_time"`
	VisitorId      string          `json:"visitor_id"`
	SessionId      string          `json:"session_id"`
	ProductDetails []ProductDetail `json:"product_details"`
}

type Search struct {
	CommonData
	SearchQuery string `json:"search_query"`
}

type DetailPageView struct {
	CommonData //
}

type AddToCart struct {
	CommonData
}

type PurchaseComplete struct {
	CommonData
	PurchaseTransaction struct {
		Id           string  `json:"id"`
		Revenue      float32 `json:"revenue"`
		Tax          float32 `json:"tax"`
		CurrencyCode string  `json:"currency_code"`
	} `json:"purchase_transaction"`
}
