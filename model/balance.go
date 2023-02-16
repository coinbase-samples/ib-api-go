/**
 * Copyright 2022-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import "time"

type AssetBalance struct {
	Asset     string  `json:"asset"`
	UserId    string  `json:"userId"`
	Amount    float32 `json:"amount"`
	Hold      float32 `json:"hold"`
	Available float32 `json:"available"`
}

type UserBalance struct {
	UserId   string            `json:"userId"`
	Interval string            `json:"interval"`
	History  []UserTimeBalance `json:"history"`
}

type UserTimeBalance struct {
	Amount float32   `json:"amount"`
	Time   time.Time `json:"date"`
}

type Balances struct {
}
