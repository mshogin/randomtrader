package validator

import (
	"testing"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/stretchr/testify/assert"
)

const testStrategyName = "some strategy"

func TestProcessContextBuyEvent(t *testing.T) {
	s := assert.New(t)

	ctx := bidcontext.NewBidContext()
	ctx.Event = config.BuyEvent
	ctx.Strategy = testStrategyName
	ctx.MinAmount = 0.01
	ctx.TickerBid = 10
	ctx.BalanceQuote = 0.05

	err := ProcessContext(ctx)
	s.Error(err)
	s.Contains(err.Error(), "inefficient quote balance")

	ctx = bidcontext.NewBidContext()
	ctx.Event = config.BuyEvent
	ctx.Strategy = testStrategyName
	ctx.MinAmount = 0.01
	ctx.TickerBid = 10
	ctx.BalanceQuote = 15

	err = ProcessContext(ctx)
	s.Error(err)
	s.Contains(err.Error(), "inefficient order size")

	ctx = bidcontext.NewBidContext()
	ctx.Event = config.BuyEvent
	ctx.Strategy = testStrategyName
	ctx.MinAmount = 0.01
	ctx.TickerBid = 10
	ctx.BalanceQuote = 15
	ctx.MinOrderSize = ctx.TickerBid*ctx.MinAmount - 1

	s.NoError(ProcessContext(ctx))
}

func TestProcessContextSellEvent(t *testing.T) {
	s := assert.New(t)

	ctx := bidcontext.NewBidContext()
	ctx.Event = config.SellEvent
	ctx.Strategy = testStrategyName
	ctx.MinAmount = 0.05
	ctx.BalanceBase = 0.01

	err := ProcessContext(ctx)
	s.Error(err)
	s.Contains(err.Error(), "inefficient base balance")

	ctx = bidcontext.NewBidContext()
	ctx.Event = config.SellEvent
	ctx.Strategy = testStrategyName
	ctx.MinAmount = 0.05
	ctx.BalanceBase = 0.1

	err = ProcessContext(ctx)
	s.Error(err)
	s.Contains(err.Error(), "inefficient order size")

	ctx = bidcontext.NewBidContext()
	ctx.Event = config.SellEvent
	ctx.Strategy = testStrategyName
	ctx.MinAmount = 0.05
	ctx.BalanceBase = 0.1
	ctx.TickerBid = 10
	ctx.MinOrderSize = ctx.TickerBid*ctx.MinAmount - 1

	s.NoError(ProcessContext(ctx))
}
