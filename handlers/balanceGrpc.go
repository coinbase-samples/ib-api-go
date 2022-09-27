package handlers

import (
	"context"

	"github.com/coinbase-samples/ib-api-go/config"
	"github.com/coinbase-samples/ib-api-go/conversions"
	"github.com/coinbase-samples/ib-api-go/model"
	ledger "github.com/coinbase-samples/ib-api-go/pkg/pbs/ledger"
	balance "github.com/coinbase-samples/ib-api-go/pkg/pbs/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

type BalanceServer struct {
	balance.UnimplementedBalanceServiceServer
	Tracer      trace.Tracer
	OrderClient ledger.LedgerClient
	appConfig   config.AppConfig
}

func (o *BalanceServer) ListBalances(ctx context.Context, req *balance.ListBalancesRequest) (*balance.ListBalancesResponse, error) {
	l := ctxlogrus.Extract(ctx)
	authedUser := ctx.Value(model.UserCtxKey).(model.User)
	if err := req.Validate(); err != nil {
		l.Debugln("invalid request", err)
		return nil, err
	}
	l.Debugln("starting trace", authedUser.Id)
	_, span := o.Tracer.Start(ctx, "listBalance",
		trace.WithAttributes(attribute.String("UserId", authedUser.Id), attribute.String("BalanceUserId", req.Id)))
	defer span.End()

	md := metadata.New(map[string]string{"x-route-id": o.appConfig.OrderRouteId})

	balances, err := o.OrderClient.GetAccounts(
		metadata.NewOutgoingContext(ctx, md),
		&ledger.GetAccountsRequest{
			UserId: req.Id,
		},
	)

	if err != nil {
		l.Warn("error listing balances", err)
		return nil, err
	}
	finalBalances := conversions.ConvertListBalancesLedgerToProto(balances)

	return &finalBalances, nil
}
