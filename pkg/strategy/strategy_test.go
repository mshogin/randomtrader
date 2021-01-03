package strategy

import (
	"testing"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/stretchr/testify/assert"
)

var processContextStub = func(ctx *bidcontext.BidContext) error {
	ctx.Event = config.BuyEvent
	ctx.Strategy = "test"
	return nil
}

func TestProcessContext(t *testing.T) {
	s := assert.New(t)
	ctx := bidcontext.NewBidContext()

	s.Empty(ctx.Strategy)
	s.Empty(ctx.Event)

	pluginsOrig := plugins
	defer func() {
		plugins = pluginsOrig
	}()
	plugins["test"] = processContextStub

	s.NoError(ProcessContext(ctx))

	s.NotEmpty(ctx.Strategy)
	s.NotEmpty(ctx.Event)
}
