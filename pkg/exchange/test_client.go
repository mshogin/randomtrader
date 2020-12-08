package exchange

import (
	"context"

	"github.com/thrasher-corp/gocryptotrader/gctrpc"
	"google.golang.org/grpc"
)

var defaultTestValue = 10.

type testClient struct{}

// GetTicker ...
func (m *testClient) GetTicker(ctx context.Context, in *gctrpc.GetTickerRequest, opts ...grpc.CallOption) (*gctrpc.TickerResponse, error) {
	return &gctrpc.TickerResponse{
		Pair: &gctrpc.CurrencyPair{
			Delimiter: "-",
			Base:      "BTC",
			Quote:     "USD",
		},
		Bid: defaultTestValue,
		Ask: defaultTestValue,
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
						TotalValue: defaultTestValue,
					},
					{
						Currency:   "USD",
						TotalValue: defaultTestValue,
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

// GetOrderbook ...
func (m *testClient) GetOrderbook(ctx context.Context, in *gctrpc.GetOrderbookRequest, opts ...grpc.CallOption) (*gctrpc.OrderbookResponse, error) {
	return &gctrpc.OrderbookResponse{
		Asks: []*gctrpc.OrderbookItem{
			{
				Price:  12.,
				Amount: 0.01,
			},
		},
		Bids: []*gctrpc.OrderbookItem{
			{
				Price:  12.,
				Amount: 0.01,
			},
		},
	}, nil
}

// Close ...
func (m *testClient) Close() error { return nil }

// SetupTestGRPCClient ...
func SetupTestGRPCClient() {
	setupClient = func() (grpcClient, error) {
		return &testClient{}, nil
	}
}
