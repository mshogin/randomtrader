package trader

import (
	"math/rand"
	"time"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
)

var (
	doneEventLoop chan bool = make(chan bool)
)

// Run ...
func Run() {
	rand.Seed(time.Now().UTC().UnixNano())
	go func() {
		bidAskTimer := time.NewTicker(config.GetEventsRaiseInterval())
		runnig := true
		for runnig {
			select {
			case <-doneEventLoop:
				runnig = false
			case <-bidAskTimer.C:
				go processContext(bidcontext.NewBidContext())
			}
		}
		<-doneEventLoop
	}()
}

// Shutdown ...
func Shutdown() {
	doneEventLoop <- true
	doneEventLoop <- true
}
