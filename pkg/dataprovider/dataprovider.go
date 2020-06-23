package dataprovider

import (
	"fmt"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
)

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
	bid_amount := (ctx.MinOrderSize + 1) / ctx.TickerBid
	ctx.MinAmount = float64(int(bid_amount*100000000)) / 100000000
	return nil
}
