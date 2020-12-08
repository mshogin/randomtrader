package strategy

import (
	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/strategy/random"
)

// ProcessContext ...
func ProcessContext(ctx *bidcontext.BidContext) error {
	if ctx.HasError() {
		return nil
	}

	if len(ctx.Strategy) == 0 {
		return random.ProcessContext(ctx)
	}

	return nil
}
