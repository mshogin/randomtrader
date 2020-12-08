package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

const (
	defaultEventRaiseInterval = 1
	defaultCurrencyPair       = "BTC-USD"
	defaultOrderSize          = 27.
	defaultExchange           = "bitstamp"
	defaultMinimumOrderSize   = 25.
	defaultLogsRoot           = "/var/log/randomtrader"
	defaultConfigsRoot        = "/etc/randomtrader"
	// BuyEvent ...
	BuyEvent = "BUY"
	// SellEvent ...
	SellEvent = "SELL"
)

// Event ...
type Event string

// String ...
func (m Event) String() string {
	return string(m)
}

// EnabledEvents ...
var EnabledEvents = []Event{BuyEvent, SellEvent}

var defaultConfigsFilename = "/etc/randomtrader/config.json"
var config = Configuration{}

// Configuration ...
type Configuration struct {
	EnableDebug    bool
	TestBrokerMode bool

	EventRaiseInterval int
	LogsRoot           string
	ConfigsRoot        string
	Exchange           string
	CurrencyPair       string
	OrderSize          float64
	MinimumOrderSize   float64

	DataCollector DataCollector

	GCEBucket          string
	ServiceKeyFilename string
}

// OrderBookLog ...
type OrderBookLog struct {
	Filename       string
	DumpInterval   int
	RotateInterval int
}

// DataCollector ...
type DataCollector struct {
	OrderBook []OrderBookLog
}

// Init ...
func Init(configPath string) error {
	setDefaults()

	if len(configPath) == 0 {
		configPath = path.Join(GetConfigsRoot(), defaultConfigsFilename)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("can't open configuration file %q: %s", configPath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var c Configuration
	if err := decoder.Decode(&c); err != nil {
		return fmt.Errorf("can't parse configuration file %q: %s", configPath, err)
	}

	SwapConfig(c)

	return nil
}

// SwapConfig ...
func SwapConfig(c Configuration) Configuration {
	oldConfig := config
	config = c
	return oldConfig
}

// GetEventsRaiseInterval ...
func GetEventsRaiseInterval() time.Duration {
	return time.Duration(config.EventRaiseInterval) * time.Second
}

// IsDebugEnabled ...
func IsDebugEnabled() bool {
	return config.EnableDebug
}

// GetCurrencyPair ...
func GetCurrencyPair() string {
	if len(config.CurrencyPair) == 0 {
		return defaultCurrencyPair
	}
	return config.CurrencyPair
}

// GetCurrencyBase ...
func GetCurrencyBase() Currency {
	c := strings.Split(GetCurrencyPair(), "-")
	return Currency(c[0])
}

// GetCurrencyQuote ...
func GetCurrencyQuote() Currency {
	c := strings.Split(GetCurrencyPair(), "-")
	return Currency(c[1])
}

// GetOrderSize ...
func GetOrderSize() float64 {
	return config.OrderSize
}

// GetExchange ...
func GetExchange() string {
	return config.Exchange
}

// GetLogsRoot ...
func GetLogsRoot() string {
	return config.LogsRoot
}

// GetConfigsRoot ...
func GetConfigsRoot() string {
	return config.ConfigsRoot
}

// GetGCEServiceKeyFilepath ...
func GetGCEServiceKeyFilepath() string {
	return path.Join(GetConfigsRoot(), config.ServiceKeyFilename)
}

// GetGCEBucket ...
func GetGCEBucket() string {
	return config.GCEBucket
}

func setDefaults() {
	config.EventRaiseInterval = defaultEventRaiseInterval
	config.EnableDebug = false
	config.CurrencyPair = defaultCurrencyPair
	config.OrderSize = defaultOrderSize
	config.Exchange = defaultExchange
	config.MinimumOrderSize = defaultMinimumOrderSize
	config.LogsRoot = defaultLogsRoot
	config.ConfigsRoot = defaultConfigsRoot
}

// GetDataCollector ...
func GetDataCollector() DataCollector {
	return config.DataCollector
}

// GetFilepath ...
func (m OrderBookLog) GetFilepath() string {
	return path.Join(GetLogsRoot(), m.Filename)
}

// GetRotateInterval ...
func (m OrderBookLog) GetRotateInterval() time.Duration {
	return time.Duration(m.RotateInterval) * time.Second
}
