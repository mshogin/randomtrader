package validator

import (
	"fmt"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
)

// ProcessContext ...
func ProcessContext(ctx *bidcontext.BidContext) error {
	if ctx.HasError() {
		return nil
	}

	if err := validateQuoteBalance(ctx); err != nil {
		return fmt.Errorf("cannot pass the quote balance validation: %w", err)
	}

	if err := validateBaseBalance(ctx); err != nil {
		return fmt.Errorf("cannot pass the base balance validation: %w", err)
	}

	if err := validateOrderSize(ctx); err != nil {
		return fmt.Errorf("cannot pass the order_size validation: %w", err)
	}

	return nil
}

func validateQuoteBalance(ctx *bidcontext.BidContext) error {
	if ctx.Event != config.BuyEvent {
		return nil
	}

	if ctx.BalanceQuote < ctx.MinAmount*ctx.TickerBid {
		return fmt.Errorf("inefficient quote balance %.8f want %.8f", ctx.BalanceQuote, ctx.TickerBid*ctx.MinAmount)
	}

	return nil
}

func validateBaseBalance(ctx *bidcontext.BidContext) error {
	if ctx.Event != config.SellEvent {
		return nil
	}

	if ctx.BalanceBase < ctx.MinAmount {
		return fmt.Errorf("inefficient base balance %.8f want %.8f", ctx.BalanceBase, ctx.MinAmount)
	}

	return nil
}

func validateOrderSize(ctx *bidcontext.BidContext) error {
	// TODO Implement conversion to the base currency
	if ctx.CurrencyQuote != config.USD {
		panic("Unsupported currency pair")
	}

	if ctx.MinOrderSize == 0 || ctx.MinAmount*ctx.TickerBid < ctx.MinOrderSize {
		return fmt.Errorf("inefficient order size: act=%v, min: %v", ctx.MinAmount*ctx.TickerBid, ctx.MinOrderSize)
	}

	return nil
}
