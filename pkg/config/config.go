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
	defatulLogsRoot           = "/var/log/randomtrader"
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

var defaultConfigFilePath = "/etc/randomtrader/config.json"
var config = Configuration{}
var configSync sync.Mutex

// Configuration ...
type Configuration struct {
	EnableDebug    bool
	TestBrokerMode bool

	EventRaiseInterval int
	LogsRoot           string
	Exchange           string
	CurrencyPair       string
	OrderSize          float64
	MinimumOrderSize   float64

	DataCollector DataCollectorConfiguration
}

type DataCollectorConfiguration struct {
	OrderBook []DataCollectorOrderBook
}

type DataCollectorOrderBook struct {
	Filename string
	Interval int
}

// Init ...
func Init(configPath string) error {
	setDefaults()

	if len(configPath) == 0 {
		configPath = defaultConfigFilePath
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

// SetConfig ...
func SwapConfig(c Configuration) Configuration {
	configSync.Lock()
	defer configSync.Unlock()
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

// GetLogsRoot ...
func GetLogsRoot() string {
	configSync.Lock()
	defer configSync.Unlock()
	return config.LogsRoot
}

func setDefaults() {
	configSync.Lock()
	defer configSync.Unlock()
	config.EventRaiseInterval = defaultEventRaiseInterval
	config.EnableDebug = false
	config.CurrencyPair = defaultCurrencyPair
	config.OrderSize = defaultOrderSize
	config.Exchange = defaultExchange
	config.MinimumOrderSize = defaultMinimumOrderSize
	config.LogsRoot = defatulLogsRoot
}

// GetDataCollector ...
func GetDataCollector() DataCollectorConfiguration {
	configSync.Lock()
	defer configSync.Unlock()
	return config.DataCollector
}

// GetFilepath ...
func (m DataCollectorOrderBook) GetFilepath() string {
	return path.Join(GetLogsRoot(), m.Filename)
}
