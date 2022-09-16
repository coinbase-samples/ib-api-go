package handlers

import (
	"context"

	"github.com/coinbase-samples/ib-api-go/model"
	order "github.com/coinbase-samples/ib-api-go/pkg/pbs/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type OrderServer struct {
	order.UnimplementedOrderServiceServer
	Tracer     trace.Tracer
	ClientConn *grpc.ClientConn
}

func (o *OrderServer) ListOrders(ctx context.Context, req *order.ListOrdersRequest) (*order.ListOrdersResponse, error) {
	l := ctxlogrus.Extract(ctx)
	authedUser := ctx.Value(model.UserCtxKey).(model.User)
	if err := req.Validate(); err != nil {
		l.Debugln("invalid request", err)
		return nil, err
	}
	l.Debugln("starting trace", authedUser.Id)
	_, span := o.Tracer.Start(ctx, "listBalance",
		trace.WithAttributes(attribute.String("UserId", authedUser.Id)))
	defer span.End()

	client := order.NewOrderServiceClient(o.ClientConn)
	orders, err := client.ListOrders(ctx, req)

	if err != nil {
		l.Warn("error listing orders", err)
		return nil, err
	}

	return orders, nil
}
