package random

import (
	"testing"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestProcessContext(t *testing.T) {
	s := assert.New(t)

	ctx := bidcontext.NewBidContext()
	s.NoError(ProcessContext(ctx))
	s.Equal(strategyName, ctx.Strategy)
	s.Contains(config.EnabledEvents, ctx.Event)
}
