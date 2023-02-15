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

type OrderStatus string

const (
	PendingInternal OrderStatus = "pendingInternal"
	PendingVenue    OrderStatus = "pendingVenue"
	Open            OrderStatus = "open"
	Filled          OrderStatus = "filled"
	Cancelled       OrderStatus = "cancelled"
	Expired         OrderStatus = "expired"
	Failed          OrderStatus = "failed"
)

type OrderType string

const (
	Market OrderType = "market"
	Limit  OrderType = "limit"
)

type OrderSide string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)

type OrderTimeInForce string

const (
	UntilDateTime     OrderTimeInForce = "untilDateTime"
	UntilCancelled    OrderTimeInForce = "untilCancelled"
	ImmediateOrCancel OrderTimeInForce = "immediateOrCancel"
	FillOrKill        OrderTimeInForce = "fillOrKill"
)

var (
	orderStatusMap = map[string]OrderStatus{
		"pendingInternal": PendingInternal,
		"pendingVenue":    PendingVenue,
		"open":            Open,
		"filled":          Filled,
		"cancelled":       Cancelled,
		"expired":         Expired,
		"failed":          Failed,
	}
	SidePrefix   = "ORDER_SIDE_"
	orderSideMap = map[string]OrderSide{
		"buy":  Buy,
		"sell": Sell,
	}
	TypePrefix   = "ORDER_TYPE_"
	OrderTypeMap = map[string]OrderType{
		"market": Market,
		"limit":  Limit,
	}
	TimeInForcePrefix = "ORDER_TIME_IN_FORCE_"
	OrderTifMap       = map[string]OrderTimeInForce{
		"untilDateTime":     UntilDateTime,
		"untilCancelled":    UntilCancelled,
		"immediateOrCancel": ImmediateOrCancel,
		"fillOrKill":        FillOrKill,
	}
)

func ParseOrderStatus(o string) OrderStatus {
	return orderStatusMap[o]
}

func ParseOrderSide(o string) OrderSide {
	return orderSideMap[o]
}
