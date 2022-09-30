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
	Filter             string           `json:"filter" dynamodbav:"filter"`
	VenueOrderId       string           `json:"venueOrderId" dynamodbav:"venueOrderId,omitempty"`
	ProductId          string           `json:"productId" dynamodbav:"productId"`
	Side               OrderSide        `json:"side" dynamodbav:"side"`
	Type               OrderType        `json:"type" dynamodbav:"type"`
	Quantity           string           `json:"quantity" dynamodbav:"quantity"`
	LimitPrice         string           `json:"limitPrice" dynamodbav:"limitPrice,omitempty"`
	TimeInForce        OrderTimeInForce `json:"timeInForce" dynamodbav:"timeInForce,omitempty"`
	Status             OrderStatus      `json:"status" dynamodbav:"orderStatus"`
	CreatedAt          time.Time        `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt          time.Time        `json:"updatedAt" dynamodbav:"updatedAt,omitempty"`
	FilledQuantity     string           `json:"filledQuantity" dynamodbav:"filledQuantity,omitempty"`
	FilledValue        string           `json:"filledValue" dynamodbav:"filledValue,omitempty"`
	AverageFilledPrice string           `json:"averageFilledPrice" dynamodbav:"averageFilledPrice,omitempty"`
	Commission         string           `json:"commission" dynamodbav:"commission,omitempty"`
	ExchangeFee        string           `json:"exchangeFee" dynamodbav:"exchangeFee,omitempty"`
}
