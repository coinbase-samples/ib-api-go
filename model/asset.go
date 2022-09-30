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
