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
        "LogsRoot": "/tmp",
        "ConfigsRoot": "/etc",
        "GCEBucket": "randomtrader-datacollector",
        "ServiceKeyFilename": "gce-bucket-service-key.json",
        "DataCollector": {
            "OrderBook": [
                {
                    "Filename": "orderbook10min.log",
                    "DumpInterval": 600,
                    "RotateInterval": 20
                }
            ]
        }
}`)
	s.NoError(ioutil.WriteFile(f.Name(), content, os.FileMode(644)))

	configOrig := config
	defer func() {
		config = configOrig
	}()

	_, err = Init(f.Name())
	s.NoError(err)

	s.Equal(time.Duration(1)*time.Second, GetEventsRaiseInterval())
	s.Equal("BTC-USD", GetCurrencyPair())
	s.Equal(30., GetOrderSize())
	s.Equal("kraken", GetExchange())

	dc := GetDataCollector()
	s.NotNil(dc)
	s.Equal(600, dc.OrderBook[0].DumpInterval)
	s.Equal("orderbook10min.log", dc.OrderBook[0].Filename)
	s.Equal(path.Join("/tmp", "orderbook10min.log"), dc.OrderBook[0].GetFilepath())
	s.Equal(path.Join("/etc", "gce-bucket-service-key.json"), GetGCEServiceKeyFilepath())
	s.Equal(20*time.Second, dc.OrderBook[0].GetRotateInterval())
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

	_, err = Init(f.Name())
	s.NoError(err)

	dc := GetDataCollector()
	s.NotNil(dc)
	s.Nil(dc.OrderBook)
}

func (s *ConfigTestSuite) TestStrategies() {
	f, err := ioutil.TempFile("", "")
	s.NoError(err)
	s.NoError(f.Close())

	content := []byte(`{
        "Strategies": {
            "archimedes": {
                "ProcessingEnabled": true,
                "RoutineEnabled": true
            }
        }
}`)
	s.NoError(ioutil.WriteFile(f.Name(), content, os.FileMode(644)))

	configOrig := config
	defer func() {
		config = configOrig
	}()

	_, err = Init(f.Name())
	s.NoError(err)

	archConf, ok := GetStrategyConfig("archimedes")
	s.True(ok)
	s.NotNil(archConf)
	s.True(archConf.ProcessingEnabled)
	s.True(archConf.RoutineEnabled)
}
