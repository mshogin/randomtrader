package strategy

import (
	"testing"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/stretchr/testify/assert"
)

func TestProcessContext(t *testing.T) {
	s := assert.New(t)
	ctx := bidcontext.NewBidContext()

	s.Empty(ctx.Strategy)
	s.Empty(ctx.Event)

	s.NoError(ProcessContext(ctx))

	s.NotEmpty(ctx.Strategy)
	s.NotEmpty(ctx.Event)
}
