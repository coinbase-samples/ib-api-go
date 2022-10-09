package conversions

import (
	"github.com/coinbase-samples/ib-api-go/model"
	balance "github.com/coinbase-samples/ib-api-go/pkg/pbs/balance/v1"
	ledger "github.com/coinbase-samples/ib-api-go/pkg/pbs/ledger/v1"
)

func ConvertListBalancesToProto(o model.Balances) balance.ListBalancesResponse {
	var balances []*balance.AccountAndBalance

	return balance.ListBalancesResponse{
		Data: balances,
	}
}

func ConvertListBalancesLedgerToProto(o *ledger.GetAccountsResponse) balance.ListBalancesResponse {
	var balances []*balance.AccountAndBalance
	for _, b := range o.Accounts {
		balances = append(balances, &balance.AccountAndBalance{
			Available: b.Available,
		})
	}
	return balance.ListBalancesResponse{
		Data: balances,
	}
}
