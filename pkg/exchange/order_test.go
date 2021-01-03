package exchange

import (
	"testing"
	"time"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/utils"
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

func TestGetOrderBook(t *testing.T) {
	s := assert.New(t)
	SetupTestGRPCClient()

	GetCurrentTimeOrig := utils.GetCurrentTime
	defer func() { utils.GetCurrentTime = GetCurrentTimeOrig }()
	currentTime := time.Now()
	utils.GetCurrentTime = func() time.Time {
		return currentTime
	}

	ob, err := GetOrderBook()
	s.NoError(err)
	s.NotNil(ob)

	s.Equal(currentTime, ob.DateTime)
}

func TestGetOrderBookHistory(t *testing.T) {
	s := assert.New(t)

	size := len(GetOrderBookHistory())
	addOrderBookItemToHistory(&OrderBook{})
	s.Len(GetOrderBookHistory(), size+1)

	for i := 0; i < orderBookHistorySize+10; i++ {
		addOrderBookItemToHistory(&OrderBook{})
	}

	s.Len(GetOrderBookHistory(), orderBookHistorySize)
}
