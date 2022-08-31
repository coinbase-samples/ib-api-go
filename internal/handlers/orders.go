package handlers

import (
	"github.com/cfluke-cb/ib-client-api/data"
)

func PlaceOrder(req data.OrderRequest) (data.OrderResponse, error) {
	body := data.OrderResponse{
		OrderId: "12345",
		Order:   req,
	}
	return body, nil
}
