package exchange

import (
	"context"

	"github.com/thrasher-corp/gocryptotrader/gctrpc"
	"github.com/thrasher-corp/gocryptotrader/gctrpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	rpcHost     = "localhost:9052"
	rpcUsername = "admin"
	rpcPassword = "test"
	certPath    = "/home/gocryptotrader/.gocryptotrader/tls/cert.pem"
)

type grpcClient interface {
	GetTicker(ctx context.Context, in *gctrpc.GetTickerRequest, opts ...grpc.CallOption) (*gctrpc.TickerResponse, error)
	UpdateAccountInfo(ctx context.Context, in *gctrpc.GetAccountInfoRequest, opts ...grpc.CallOption) (*gctrpc.GetAccountInfoResponse, error)
	SubmitOrder(ctx context.Context, in *gctrpc.SubmitOrderRequest, opts ...grpc.CallOption) (*gctrpc.SubmitOrderResponse, error)
	GetOrderbook(ctx context.Context, in *gctrpc.GetOrderbookRequest, opts ...grpc.CallOption) (*gctrpc.OrderbookResponse, error)

	Close() error
}

type gctClient struct {
	gctrpc.GoCryptoTraderClient
	conn *grpc.ClientConn
}

func (m *gctClient) Close() error {
	return m.conn.Close()
}

var setupClient = func() (grpcClient, error) {
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(
			auth.BasicAuth{
				Username: rpcUsername,
				Password: rpcPassword,
			},
		),
	}
	conn, err := grpc.Dial(rpcHost, opts...)
	if err != nil {
		return nil, err
	}

	c := gctrpc.NewGoCryptoTraderClient(conn)

	return &gctClient{
		GoCryptoTraderClient: c,
		conn:                 conn,
	}, nil
}
