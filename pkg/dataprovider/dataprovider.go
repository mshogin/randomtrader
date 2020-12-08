package dataprovider

import (
	"fmt"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
)

const precision = 100000000

// ProcessContext ...
func ProcessContext(ctx *bidcontext.BidContext) error {
	if ctx.HasError() {
		return nil
	}
	if err := setMinAmount(ctx); err != nil {
		return fmt.Errorf("cannot set minimal amount: %w", err)
	}
	return nil
}

func setMinAmount(ctx *bidcontext.BidContext) error {
	if ctx.TickerBid <= 0 {
		return fmt.Errorf("invalid 'TickerBid' value: %v", ctx.TickerBid)
	}
	bidAmount := (ctx.MinOrderSize + 1) / ctx.TickerBid
	ctx.MinAmount = float64(int(bidAmount*precision)) / precision

	ctx.AskPrice = ctx.MinAmount * ctx.TickerAsk
	ctx.BidPrice = ctx.MinAmount * ctx.TickerBid
	return nil
}
