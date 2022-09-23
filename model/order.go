package model

import "time"

type OrderRequest struct {
	ProductId   string           `json:"productId"`
	Side        OrderSide        `json:"side"`
	Type        OrderType        `json:"type"`
	Quantity    float32          `json:"quantity"`
	LimitPrice  float32          `json:"limitPrice"`
	TimeInForce OrderTimeInForce `json:"timeInForce"`
}

type Order struct {
	OrderId            string           `json:"clientOrderId" dynamodbav:"orderId"`
	OwnerId            string           `json:"ownerId" dynamodbav:"userId"`
	VenueOrderId       string           `json:"venueOrderId" dynamodbav:",omitempty"`
	ProductId          string           `json:"productId"`
	Side               OrderSide        `json:"side"`
	Type               OrderType        `json:"type"`
	Quantity           float32          `json:"quantity"`
	LimitPrice         float32          `json:"limitPrice" dynamodbav:",omitempty"`
	TimeInForce        OrderTimeInForce `json:"timeInForce" dynamodbav:",omitempty"`
	Status             OrderStatus      `json:"status"`
	CreatedAt          time.Time        `json:"createdAt"`
	UpdatedAt          time.Time        `json:"updatedAt"`
	FilledQuantity     float32          `json:"filledQuantity"`
	FilledValue        float32          `json:"filledValue"`
	AverageFilledPrice float32          `json:"averageFilledPrice"`
	Commission         float32          `json:"commission"`
	ExchangeFee        float32          `json:"exchangeFee" dynamodbav:",omitempty"`
}
