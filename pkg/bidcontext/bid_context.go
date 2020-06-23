package bidcontext

import (
	"strings"

	"github.com/mshogin/randomtrader/pkg/config"
)

type BidContext struct {
	Exchange      string
	MinOrderSize  float64
	CurrencyBase  config.Currency
	CurrencyQuote config.Currency

	Error string

	Event        config.Event
	TickerBid    float64
	TickerAsk    float64
	BalanceBase  float64
	BalanceQuote float64
	MinAmount    float64

	OrderID string
}

func NewBidContext(e config.Event) *BidContext {
	c := strings.Split(config.GetCurrencyPair(), "-")

	return &BidContext{
		Event:         e,
		CurrencyBase:  config.Currency(c[0]),
		CurrencyQuote: config.Currency(c[1]),
		Exchange:      config.GetExchange(),
		MinOrderSize:  config.GetOrderSize(),
	}
}

// SetError ...
func (m *BidContext) SetError(e error) {
	m.Error = e.Error()
}

// HasError ...
func (m *BidContext) HasError() bool {
	return len(m.Error) > 0
}
