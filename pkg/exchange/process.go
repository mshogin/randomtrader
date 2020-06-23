package exchange

import (
	"context"
	"errors"
	"fmt"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/thrasher-corp/gocryptotrader/gctrpc"
)

// ProcessContext ...
func ProcessContext(ctx *bidcontext.BidContext) error {
	var err error

	if err = updateBalanceInfo(ctx); err != nil {
		return fmt.Errorf("cannot update balance: %w", err)
	}

	if err = updateTickerInfo(ctx); err != nil {
		return fmt.Errorf("cannot update ticker: %w", err)
	}

	return nil
}

func updateTickerInfo(ctx *bidcontext.BidContext) error {
	c, err := setupClient()
	if err != nil {
		return fmt.Errorf("cannot setup gct client: %w", err)
	}
	defer c.Close()

	result, err := c.GetTicker(context.Background(), &gctrpc.GetTickerRequest{
		Exchange: ctx.Exchange,
		Pair: &gctrpc.CurrencyPair{
			Delimiter: "-",
			Base:      ctx.CurrencyBase.String(),
			Quote:     ctx.CurrencyQuote.String(),
		},
		AssetType: "spot",
	})
	if err != nil {
		return fmt.Errorf("cannot get ticker info: %w", err)
	}

	ctx.TickerBid = result.Bid
	ctx.TickerAsk = result.Ask
	return nil
}

func updateBalanceInfo(ctx *bidcontext.BidContext) error {
	c, err := setupClient()
	if err != nil {
		return fmt.Errorf("cannot setup gct client: %s", err)
	}
	defer c.Close()

	r, err := c.UpdateAccountInfo(
		context.Background(),
		&gctrpc.GetAccountInfoRequest{
			Exchange: ctx.Exchange,
		},
	)
	if err != nil {
		return fmt.Errorf("cannot get account info: %s", err)
	}

	if len(r.Accounts) == 0 {
		return errors.New("cannot determine account: len == 0")
	}

	for _, c := range r.Accounts[0].Currencies {
		switch c.Currency {
		case ctx.CurrencyBase.String():
			ctx.BalanceBase = c.TotalValue
		case ctx.CurrencyQuote.String():
			ctx.BalanceQuote = c.TotalValue
		default:
			continue
		}
	}

	return nil
}
