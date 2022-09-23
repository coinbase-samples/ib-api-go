package conversions

import (
	"time"

	"github.com/coinbase-samples/ib-api-go/model"
	order "github.com/coinbase-samples/ib-api-go/pkg/pbs/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertCreateOrderToProto(o model.Order) order.CreateOrderResponse {
	return order.CreateOrderResponse{
		OrderId:     o.OrderId,
		ProductId:   o.ProductId,
		OwnerId:     o.OwnerId,
		Side:        ProtoOrderSide(o.Side),
		Type:        ProtoOrderType(o.Type),
		Quantity:    o.Quantity,
		LimitPrice:  o.LimitPrice,
		TimeInForce: ProtoOrderTimeInForce(o.TimeInForce),
		Status:      ProtoOrderStatus(o.Status),
		CreatedAt:   timestamppb.New(o.CreatedAt),
		UpdatedAt:   timestamppb.New(o.UpdatedAt),
	}
}

func ConvertReadOrderToProto(o model.Order) order.ReadOrderResponse {
	return order.ReadOrderResponse{
		OrderId:            o.OrderId,
		ProductId:          o.ProductId,
		OwnerId:            o.OwnerId,
		Side:               ProtoOrderSide(o.Side),
		Type:               ProtoOrderType(o.Type),
		Quantity:           o.Quantity,
		LimitPrice:         o.LimitPrice,
		TimeInForce:        ProtoOrderTimeInForce(o.TimeInForce),
		Status:             ProtoOrderStatus(o.Status),
		CreatedAt:          timestamppb.New(o.CreatedAt),
		UpdatedAt:          timestamppb.New(o.UpdatedAt),
		FilledQuantity:     o.FilledQuantity,
		FilledValue:        o.FilledValue,
		AverageFilledPrice: o.AverageFilledPrice,
		Commission:         o.Commission,
		ExchangeFee:        o.ExchangeFee,
	}
}

func ConvertUpdateOrderToProto(o model.Order) order.UpdateOrderResponse {
	return order.UpdateOrderResponse{
		OrderId:            o.OrderId,
		ProductId:          o.ProductId,
		OwnerId:            o.OwnerId,
		Side:               ProtoOrderSide(o.Side),
		Type:               ProtoOrderType(o.Type),
		Quantity:           o.Quantity,
		LimitPrice:         o.LimitPrice,
		TimeInForce:        ProtoOrderTimeInForce(o.TimeInForce),
		Status:             ProtoOrderStatus(o.Status),
		CreatedAt:          timestamppb.New(o.CreatedAt),
		UpdatedAt:          timestamppb.New(o.UpdatedAt),
		FilledQuantity:     o.FilledQuantity,
		FilledValue:        o.FilledValue,
		AverageFilledPrice: o.AverageFilledPrice,
		Commission:         o.Commission,
		ExchangeFee:        o.ExchangeFee,
	}
}

func ConvertListOrdersToProto(orders []model.Order, cursor string, hasNext bool) order.ListOrdersResponse {
	var data []*order.ReadOrderResponse

	for _, o := range orders {
		data = append(data,
			&order.ReadOrderResponse{OrderId: o.OrderId,
				ProductId:          o.ProductId,
				OwnerId:            o.OwnerId,
				Side:               ProtoOrderSide(o.Side),
				Type:               ProtoOrderType(o.Type),
				Quantity:           o.Quantity,
				LimitPrice:         o.LimitPrice,
				TimeInForce:        ProtoOrderTimeInForce(o.TimeInForce),
				Status:             ProtoOrderStatus(o.Status),
				CreatedAt:          timestamppb.New(o.CreatedAt),
				UpdatedAt:          timestamppb.New(o.UpdatedAt),
				FilledQuantity:     o.FilledQuantity,
				FilledValue:        o.FilledValue,
				AverageFilledPrice: o.AverageFilledPrice,
				Commission:         o.Commission,
				ExchangeFee:        o.ExchangeFee,
			})
	}

	return order.ListOrdersResponse{
		Data:       data,
		Count:      int32(len(data)),
		NextCursor: cursor,
		HasNext:    hasNext,
	}
}

func ConvertNewOrderToModel(req *order.CreateOrderRequest, orderId, requestUserId string) model.Order {
	return model.Order{
		OrderId:     orderId,
		OwnerId:     requestUserId,
		ProductId:   req.ProductId,
		Side:        ModelOrderSide(req.Side),
		Type:        ModelOrderType(req.Type),
		Quantity:    req.Quantity,
		LimitPrice:  req.LimitPrice,
		TimeInForce: ModelOrderTimeInForce(req.TimeInForce),
		Status:      model.PendingInternal,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ConvertUpdateOrderToModel(req *order.UpdateOrderRequest, curr model.Order) model.Order {
	return model.Order{
		OrderId:            curr.OrderId,
		OwnerId:            curr.OwnerId,
		ProductId:          curr.ProductId,
		Side:               curr.Side,
		Type:               curr.Type,
		Quantity:           curr.Quantity,
		LimitPrice:         curr.LimitPrice,
		TimeInForce:        curr.TimeInForce,
		CreatedAt:          curr.CreatedAt,
		Status:             ModelOrderStatus(req.Status),
		UpdatedAt:          time.Now(),
		FilledQuantity:     req.FilledQuantity,
		FilledValue:        req.FilledValue,
		AverageFilledPrice: req.AverageFilledPrice,
		Commission:         req.Commission,
		ExchangeFee:        req.ExchangeFee,
	}
}

func ProtoOrderSide(os model.OrderSide) order.OrderSide {
	switch o := os; o {
	case model.Buy:
		return order.OrderSide_ORDER_SIDE_BUY
	case model.Sell:
		return order.OrderSide_ORDER_SIDE_SELL
	default:
		return order.OrderSide_ORDER_SIDE_UNSPECIFIED
	}
}

func ProtoOrderType(ot model.OrderType) order.OrderType {
	switch o := ot; o {
	case model.Limit:
		return order.OrderType_ORDER_TYPE_LIMIT
	case model.Market:
		return order.OrderType_ORDER_TYPE_MARKET
	default:
		return order.OrderType_ORDER_TYPE_UNSPECIFIED
	}
}

func ProtoOrderTimeInForce(ot model.OrderTimeInForce) order.OrderTimeInForce {
	switch o := ot; o {
	case model.UntilDateTime:
		return order.OrderTimeInForce_ORDER_TIME_IN_FORCE_GOOD_UNTIL_DATETIME
	case model.UntilCancelled:
		return order.OrderTimeInForce_ORDER_TIME_IN_FORCE_GOOD_UNTIL_CANCELLED
	case model.ImmediateOrCancel:
		return order.OrderTimeInForce_ORDER_TIME_IN_FORCE_IMMEDIATE_OR_CANCEL
	case model.FillOrKill:
		return order.OrderTimeInForce_ORDER_TIME_IN_FORCE_FILL_OR_KILL
	default:
		return order.OrderTimeInForce_ORDER_TIME_IN_FORCE_UNSPECIFIED
	}
}

func ProtoOrderStatus(os model.OrderStatus) order.OrderStatus {
	switch o := os; o {
	case model.PendingInternal:
		return order.OrderStatus_ORDER_STATUS_PENDING
	case model.Open:
		return order.OrderStatus_ORDER_STATUS_OPEN
	case model.Cancelled:
		return order.OrderStatus_ORDER_STATUS_CANCELLED
	case model.Filled:
		return order.OrderStatus_ORDER_STATUS_FILLED
	case model.Expired:
		return order.OrderStatus_ORDER_STATUS_EXPIRED
	case model.Failed:
		return order.OrderStatus_ORDER_STATUS_FAILED
	default:
		return order.OrderStatus_ORDER_STATUS_UNSPECIFIED
	}
}

func ModelOrderStatus(os order.OrderStatus) model.OrderStatus {
	switch o := os; o {
	case order.OrderStatus_ORDER_STATUS_PENDING:
		return model.PendingVenue
	case order.OrderStatus_ORDER_STATUS_OPEN:
		return model.Open
	case order.OrderStatus_ORDER_STATUS_CANCELLED:
		return model.Cancelled
	case order.OrderStatus_ORDER_STATUS_FILLED:
		return model.Filled
	case order.OrderStatus_ORDER_STATUS_EXPIRED:
		return model.Expired
	case order.OrderStatus_ORDER_STATUS_FAILED:
		return model.Failed
	default:
		return model.PendingInternal
	}
}

func ModelOrderSide(os order.OrderSide) model.OrderSide {
	switch o := os; o {
	case order.OrderSide_ORDER_SIDE_BUY:
		return model.Buy
	case order.OrderSide_ORDER_SIDE_SELL:
		return model.Sell
	}
	return "Unspecified"
}

func ModelOrderType(os order.OrderType) model.OrderType {
	switch o := os; o {
	case order.OrderType_ORDER_TYPE_LIMIT:
		return model.Limit
	case order.OrderType_ORDER_TYPE_MARKET:
		return model.Market
	}
	return "Unspecified"
}

func ModelOrderTimeInForce(os order.OrderTimeInForce) model.OrderTimeInForce {
	switch o := os; o {
	case order.OrderTimeInForce_ORDER_TIME_IN_FORCE_GOOD_UNTIL_DATETIME:
		return model.UntilDateTime
	case order.OrderTimeInForce_ORDER_TIME_IN_FORCE_GOOD_UNTIL_CANCELLED:
		return model.UntilCancelled
	case order.OrderTimeInForce_ORDER_TIME_IN_FORCE_IMMEDIATE_OR_CANCEL:
		return model.ImmediateOrCancel
	case order.OrderTimeInForce_ORDER_TIME_IN_FORCE_FILL_OR_KILL:
		return model.FillOrKill
	}
	return "Unspecified"
}
