package exchange

import (
	"testing"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestProcessContext(t *testing.T) {
	s := assert.New(t)

	clientOrig := setupClient
	defer func() {
		setupClient = clientOrig
	}()
	setupClient = setupTestClient

	ctx := bidcontext.NewBidContext(config.BuyEvent)

	s.NoError(ProcessContext(ctx))
	s.Greater(ctx.TickerBid, 0.)
	s.Greater(ctx.TickerAsk, 0.)
	s.Greater(ctx.BalanceBase, 0.)
	s.Greater(ctx.BalanceQuote, 0.)

}
