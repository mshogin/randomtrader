package config

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) TestConfigFile() {
	f, err := ioutil.TempFile("", "")
	s.NoError(err)
	s.NoError(f.Close())

	content := []byte(`{
	"EventRaiseInterval": 1,
	"CurrencyPair": "BTC-USD",
        "OrderSize": 30,
        "Exchange": "kraken",
        "MinimumOrderSize": 10,
        "DataCollector": {
            "OrderBook": [
                {
                    "filename": "orderbook10min.log",
                    "interval": 600
                }
            ]
        }
}`)
	s.NoError(ioutil.WriteFile(f.Name(), content, os.FileMode(644)))

	configOrig := config
	defer func() {
		config = configOrig
	}()

	s.NoError(Init(f.Name()))

	s.Equal(time.Duration(1)*time.Second, GetEventsRaiseInterval())
	s.Equal("BTC-USD", GetCurrencyPair())
	s.Equal(30., GetOrderSize())
	s.Equal("kraken", GetExchange())

	dc := GetDataCollector()
	s.NotNil(dc)
	s.Equal(600, dc.OrderBook[0].Interval)
	s.Equal("orderbook10min.log", dc.OrderBook[0].Filename)
	s.Equal(path.Join(defatulLogsRoot, "orderbook10min.log"), dc.OrderBook[0].GetFilepath())
}

func (s *ConfigTestSuite) TestDefaultDataCollector() {
	f, err := ioutil.TempFile("", "")
	s.NoError(err)
	s.NoError(f.Close())

	content := []byte(`{
}`)
	s.NoError(ioutil.WriteFile(f.Name(), content, os.FileMode(644)))

	configOrig := config
	defer func() {
		config = configOrig
	}()

	s.NoError(Init(f.Name()))

	dc := GetDataCollector()
	s.NotNil(dc)
	s.Nil(dc.OrderBook)
}
