package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	defaultEventRaiseInterval = 1
	defaultCurrencyPair       = "BTC-USD"
	defaultOrderSize          = 27.
	defaultExchange           = "bitstamp"
	defaultMinimumOrderSize   = 25.

	BuyEvent  = "BUY"
	SellEvent = "SELL"
)

type Event string

// String ...
func (m Event) String() string {
	return string(m)
}

var EnabledEvents = []Event{BuyEvent, SellEvent}

type Configuration struct {
	EnableDebug    bool
	TestBrokerMode bool

	EventRaiseInterval int

	Exchange         string
	CurrencyPair     string
	OrderSize        float64
	MinimumOrderSize float64
}

var defaultConfigFilePath = "/etc/randomtrader/config.json"
var config = Configuration{}

func Init(configPath string) error {
	SetDefaults()

	if len(configPath) == 0 {
		configPath = defaultConfigFilePath
	}

	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("can't open configuration file %q: %s", configPath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return fmt.Errorf("can't parse configuration file %q: %s", configPath, err)
	}

	return nil
}

func GetEventsRaiseInterval() time.Duration {
	return time.Duration(config.EventRaiseInterval) * time.Second
}

func IsDebugEnabled() bool {
	return config.EnableDebug
}

func GetCurrencyPair() string {
	if len(config.CurrencyPair) == 0 {
		return defaultCurrencyPair
	}
	return config.CurrencyPair
}

func GetOrderSize() float64 {
	return config.OrderSize
}

func GetMinimumOrderSize() float64 {
	return config.MinimumOrderSize
}

func GetCurrencyBase() Currency {
	p := strings.Split(config.CurrencyPair, "-")
	return Currency(p[0])
}

func GetCurrencyQuote() Currency {
	p := strings.Split(config.CurrencyPair, "-")
	return Currency(p[1])
}

func GetExchange() string {
	return config.Exchange
}

func SetDefaults() {
	config.EventRaiseInterval = defaultEventRaiseInterval
	config.EnableDebug = false
	config.CurrencyPair = defaultCurrencyPair
	config.OrderSize = defaultOrderSize
	config.Exchange = defaultExchange
	config.MinimumOrderSize = defaultMinimumOrderSize
}
