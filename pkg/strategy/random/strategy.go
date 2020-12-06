package random

import (
	"math/rand"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
)

const strategyName = "random"

// ProcessContext ...
func ProcessContext(ctx *bidcontext.BidContext) error {
	ctx.Strategy = strategyName
	ctx.Event = config.EnabledEvents[rand.Intn(len(config.EnabledEvents))]
	return nil
}
