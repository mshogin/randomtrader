package trader

import (
	"fmt"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/dataprovider"
	"github.com/mshogin/randomtrader/pkg/exchange"
	"github.com/mshogin/randomtrader/pkg/logger"
	"github.com/mshogin/randomtrader/pkg/strategy"
	"github.com/mshogin/randomtrader/pkg/validator"
)

func processContext(ctx *bidcontext.BidContext) {
	if err := exchange.ProcessContext(ctx); err != nil {
		ctx.SetError(fmt.Errorf("cannot process exchange: %w", err))
	}

	if err := dataprovider.ProcessContext(ctx); err != nil {
		ctx.SetError(fmt.Errorf("cannot process dataprovider: %w", err))
	}

	if err := strategy.ProcessContext(ctx); err != nil {
		ctx.SetError(fmt.Errorf("cannot process strategy: %w", err))
	}

	if err := validator.ProcessContext(ctx); err != nil {
		ctx.SetError(fmt.Errorf("cannot process validator: %w", err))
	}

	if err := exchange.ExecuteContext(ctx); err != nil {
		ctx.SetError(fmt.Errorf("cannot execute context: %w", err))
	}

	if err := logger.ProcessContext(ctx); err != nil {
		panic(fmt.Errorf("cannot execute context: %w", err))
	}
}
