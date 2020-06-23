package trader

import (
	"math/rand"
	"time"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
)

var (
	doneEventLoop    chan bool = make(chan bool)
	startedEventLoop chan bool = make(chan bool)
)

func Run() {
	go func() {
		bidAskTimer := time.NewTicker(config.GetEventsRaiseInterval())
		runnig := true
		for runnig {
			select {
			case <-startedEventLoop:
				<-startedEventLoop
			case <-doneEventLoop:
				runnig = false
			case <-bidAskTimer.C:
				e := config.EnabledEvents[rand.Intn(len(config.EnabledEvents))]
				go processContext(bidcontext.NewBidContext(e))
			}
		}
		<-doneEventLoop
	}()

	startedEventLoop <- true
	startedEventLoop <- true
}

func Shutdown() {
	doneEventLoop <- true
	doneEventLoop <- true
}
