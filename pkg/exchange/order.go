package exchange

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mshogin/randomtrader/pkg/bidcontext"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/utils"
	"github.com/thrasher-corp/gocryptotrader/gctrpc"
)

const orderBookHistorySize = 60

type (
	orderBookItem struct {
		Amount float64
		Price  float64
	}

	// OrderBook ...
	OrderBook struct {
		Asks     []orderBookItem
		Bids     []orderBookItem
		DateTime time.Time
	}
)

var orderBookHistory = make([]*OrderBook, 0)
var orderBookHistorySync sync.Mutex

// ExecuteContext ...
func ExecuteContext(ctx *bidcontext.BidContext) error {
	if ctx.HasError() {
		return nil
	}

	c, err := setupClient()
	if err != nil {
		return fmt.Errorf("cannot create broker: %s", err)
	}
	defer c.Close()

	result, err := c.SubmitOrder(
		context.Background(),
		&gctrpc.SubmitOrderRequest{
			Exchange: ctx.Exchange,
			Pair: &gctrpc.CurrencyPair{
				Delimiter: "-",
				Base:      ctx.CurrencyBase.String(),
				Quote:     ctx.CurrencyQuote.String(),
			},
			Side:      ctx.Event.String(),
			OrderType: "MARKET",
			Amount:    ctx.MinAmount,
			AssetType: "SPOT",
		},
	)
	if err != nil {
		return fmt.Errorf("cannot place order: %s", err)
	}

	ctx.OrderID = result.GetOrderId()
	return nil
}

// GetOrderBook ...
func GetOrderBook() (*OrderBook, error) {
	c, err := setupClient()
	if err != nil {
		return nil, fmt.Errorf("cannot create broker: %s", err)
	}
	defer c.Close()

	result, err := c.GetOrderbook(
		context.Background(),
		&gctrpc.GetOrderbookRequest{
			Exchange: config.GetExchange(),
			Pair: &gctrpc.CurrencyPair{
				Delimiter: "-",
				Base:      config.GetCurrencyBase().String(),
				Quote:     config.GetCurrencyQuote().String(),
			},
			AssetType: "SPOT",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("cannot place order: %s", err)
	}

	ob := OrderBook{
		DateTime: utils.GetCurrentTime(),
	}
	for _, ask := range result.GetAsks() {
		ob.Asks = append(ob.Asks,
			orderBookItem{Amount: ask.Amount, Price: ask.Price})
	}
	for _, bid := range result.GetBids() {
		ob.Bids = append(ob.Bids,
			orderBookItem{Amount: bid.Amount, Price: bid.Price})
	}

	addOrderBookItemToHistory(&ob)

	return &ob, nil
}

func addOrderBookItemToHistory(ob *OrderBook) {
	orderBookHistorySync.Lock()
	defer orderBookHistorySync.Unlock()
	if len(orderBookHistory) == orderBookHistorySize {
		orderBookHistory = orderBookHistory[1:]
	}
	orderBookHistory = append(orderBookHistory, ob)
}

// GetOrderBookHistory ...
func GetOrderBookHistory() []*OrderBook {
	orderBookHistorySync.Lock()
	defer orderBookHistorySync.Unlock()
	h := make([]*OrderBook, len(orderBookHistory))
	copy(h, orderBookHistory)
	return h
}
