package exchange

import (
	"testing"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestExecuteContext(t *testing.T) {
	s := assert.New(t)
	SetupTestGRPCClient()

	for _, e := range []config.Event{config.BuyEvent, config.SellEvent} {
		ctx := bidcontext.NewBidContext()
		ctx.Event = e
		s.NoError(ExecuteContext(ctx))
		s.NotEmpty(ctx.OrderID)
	}
}
