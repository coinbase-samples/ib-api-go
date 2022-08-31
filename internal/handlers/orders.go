package handlers

import (
	"github.com/cfluke-cb/ib-client-api/model"
)

func PlaceOrder(req model.OrderRequest) (model.OrderResponse, error) {
	body := model.OrderResponse{
		OrderId: "12345",
		Order:   req,
	}
	return body, nil
}
