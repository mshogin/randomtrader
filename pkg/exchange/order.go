package exchange

import (
	"context"
	"fmt"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/thrasher-corp/gocryptotrader/gctrpc"
)

func ExecuteContext(ctx *bidcontext.BidContext) error {
	if ctx.HasError() {
		return nil
	}

	c, err := setupClient()
	if err != nil {
		return fmt.Errorf("cannot create broker: %s", err)
	}
	defer c.Close()

	result, err := c.SubmitOrder(
		context.Background(),
		&gctrpc.SubmitOrderRequest{
			Exchange: ctx.Exchange,
			Pair: &gctrpc.CurrencyPair{
				Delimiter: "-",
				Base:      ctx.CurrencyBase.String(),
				Quote:     ctx.CurrencyQuote.String(),
			},
			Side:      ctx.Event.String(),
			OrderType: "MARKET",
			Amount:    ctx.MinAmount,
			AssetType: "SPOT",
		},
	)
	if err != nil {
		return fmt.Errorf("cannot place order: %s", err)
	}

	ctx.OrderID = result.GetOrderId()
	return nil
}
