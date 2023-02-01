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
