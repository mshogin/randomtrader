package dataprovider

import (
	"testing"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/stretchr/testify/assert"
)

func TestProcessContext(t *testing.T) {
	s := assert.New(t)
	ctx := &bidcontext.BidContext{TickerBid: 0.}

	s.Error(ProcessContext(ctx))

	ctx = &bidcontext.BidContext{TickerBid: 1., MinOrderSize: 1.}
	s.NoError(ProcessContext(ctx))

	s.Equal(1., ctx.MinAmount)
}
