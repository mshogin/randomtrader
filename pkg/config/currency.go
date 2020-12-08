package config

// Currency ...
type Currency string

// String ...
func (c Currency) String() string {
	return string(c)
}

// BTC ...
var BTC = Currency("BTC")

// USD ...
var USD = Currency("USD")
