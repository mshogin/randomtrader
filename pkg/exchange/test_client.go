package exchange

import (
	"context"

	"github.com/thrasher-corp/gocryptotrader/gctrpc"
	"google.golang.org/grpc"
)

type testClient struct{}

// GetTicker ...
func (m *testClient) GetTicker(ctx context.Context, in *gctrpc.GetTickerRequest, opts ...grpc.CallOption) (*gctrpc.TickerResponse, error) {
	return &gctrpc.TickerResponse{
		Pair: &gctrpc.CurrencyPair{
			Delimiter: "-",
			Base:      "BTC",
			Quote:     "USD",
		},
		Bid: 10.,
		Ask: 10.,
	}, nil
}

// UpdateAccountInfo ...
func (m *testClient) UpdateAccountInfo(ctx context.Context, in *gctrpc.GetAccountInfoRequest, opts ...grpc.CallOption) (*gctrpc.GetAccountInfoResponse, error) {
	return &gctrpc.GetAccountInfoResponse{
		Exchange: "test-exchange",
		Accounts: []*gctrpc.Account{
			{
				Currencies: []*gctrpc.AccountCurrencyInfo{
					{
						Currency:   "BTC",
						TotalValue: 10.,
					},
					{
						Currency:   "USD",
						TotalValue: 10.,
					},
				},
			},
		},
	}, nil
}

// SubmitOrder ...
func (m *testClient) SubmitOrder(ctx context.Context, in *gctrpc.SubmitOrderRequest, opts ...grpc.CallOption) (*gctrpc.SubmitOrderResponse, error) {
	return &gctrpc.SubmitOrderResponse{
		OrderId:     "1",
		OrderPlaced: true,
	}, nil
}

// Close ...
func (m *testClient) Close() error { return nil }

// setupTestClient ...
func setupTestClient() (grpcClient, error) {
	return &testClient{}, nil
}
