package config

type Currency string

// String ...
func (c Currency) String() string {
	return string(c)
}

var BTC = Currency("BTC")
var USD = Currency("USD")
