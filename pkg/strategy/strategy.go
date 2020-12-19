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

	// if err := lisa.ProcessContext(ctx); err != nil {
	// 	logger.Errorf("strategy lisa finished with error: %w", err)
	// }

	if len(ctx.Strategy) == 0 {
		return random.ProcessContext(ctx)
	}

	return nil
}
