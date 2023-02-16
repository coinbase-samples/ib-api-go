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

type Asset struct {
	AssetId              string    `json:"assetId" dynamodbav:"productId"`
	Filter               string    `json:"filter" dynamodbav:"filter"`
	Ticker               string    `json:"ticker" dynamodbav:"ticker"`
	Name                 string    `json:"name" dynamodbav:"name"`
	MinTransactionAmount string    `json:"minTransactionAmount" dynamodbav:"minTransactionAmount"`
	MaxTransactionAmount string    `json:"maxTransactionAmount" dynamodbav:"maxTransactionAmount"`
	Slippage             string    `json:"slippage" dynamodbav:"slippage"`
	HighOffer            string    `json:"highOffer" dynamodbav:"highOffer"`
	LowBid               string    `json:"lowBid" dynamodbav:"lowBid"`
	Spread               string    `json:"spread" dynamodbav:"spread"`
	CreatedAt            time.Time `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt" dynamodbav:"updatedAt"`
	MarketCap            string    `json:"marketCap" dynamodbav:"marketCap"`
	Volume               string    `json:"volume" dynamodbav:"volume"`
	Supply               string    `json:"supply" dynamodbav:"supply"`
	Direction            string    `json:"direction" dynamodbav:"direction"`
}
