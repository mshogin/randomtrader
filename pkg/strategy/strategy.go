package strategy

import (
	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/logger"
)

// ProcessContext ...
func ProcessContext(ctx *bidcontext.BidContext) error {
	if ctx.HasError() {
		return nil
	}

	for name, processContextCallback := range plugins {
		if err := processContextCallback(ctx); err != nil {
			logger.Errorf("strategy %q finished with error: %w", name, err)
		}
	}

	return nil
}
