package bidcontext

import (
	"github.com/mshogin/randomtrader/pkg/config"
)

// BidContext ...
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

	AskPrice float64
	BidPrice float64

	Strategy string

	OrderID string
}

// NewBidContext ...
func NewBidContext() *BidContext {
	return &BidContext{
		CurrencyBase:  config.GetCurrencyBase(),
		CurrencyQuote: config.GetCurrencyQuote(),
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
