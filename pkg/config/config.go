package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
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

var config = Configuration{}
var configSync sync.Mutex

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
func Init(configPath string) (Configuration, error) {
	configSync.Lock()
	defer configSync.Unlock()

	setDefaults()

	file, err := os.Open(configPath)
	if err != nil {
		return config, fmt.Errorf("can't open configuration file %q: %s", configPath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var c Configuration
	if err := decoder.Decode(&c); err != nil {
		return config, fmt.Errorf("can't parse configuration file %q: %s", configPath, err)
	}

	return swapConfig(c), nil
}

// SwapConfig ...
func SwapConfig(c Configuration) Configuration {
	configSync.Lock()
	defer configSync.Unlock()
	return swapConfig(c)
}

func swapConfig(c Configuration) Configuration {
	oldConfig := config
	config = c
	return oldConfig
}

// GetEventsRaiseInterval ...
func GetEventsRaiseInterval() time.Duration {
	configSync.Lock()
	defer configSync.Unlock()
	return time.Duration(config.EventRaiseInterval) * time.Second
}

// IsDebugEnabled ...
func IsDebugEnabled() bool {
	configSync.Lock()
	defer configSync.Unlock()
	return config.EnableDebug
}

// GetCurrencyPair ...
func GetCurrencyPair() string {
	configSync.Lock()
	defer configSync.Unlock()
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
	configSync.Lock()
	defer configSync.Unlock()
	return config.OrderSize
}

// GetExchange ...
func GetExchange() string {
	configSync.Lock()
	defer configSync.Unlock()
	return config.Exchange
}

// getLogsRoot ...
func getLogsRoot() string {
	return config.LogsRoot
}

// getConfigsRoot ...
func getConfigsRoot() string {
	return config.ConfigsRoot
}

// GetGCEServiceKeyFilepath ...
func GetGCEServiceKeyFilepath() string {
	configSync.Lock()
	defer configSync.Unlock()
	return path.Join(getConfigsRoot(), config.ServiceKeyFilename)
}

// GetGCEBucket ...
func GetGCEBucket() string {
	configSync.Lock()
	defer configSync.Unlock()
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
	configSync.Lock()
	defer configSync.Unlock()
	return config.DataCollector
}

// GetFilepath ...
func (m OrderBookLog) GetFilepath() string {
	configSync.Lock()
	defer configSync.Unlock()
	return path.Join(getLogsRoot(), m.Filename)
}

// GetRotateInterval ...
func (m OrderBookLog) GetRotateInterval() time.Duration {
	configSync.Lock()
	defer configSync.Unlock()
	return time.Duration(m.RotateInterval) * time.Second
}
