package config

import (
	"io/ioutil"
	"os"
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
        "MinimumOrderSize": 10
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
	s.Equal(BTC, GetCurrencyBase())
	s.Equal(USD, GetCurrencyQuote())
	s.Equal("kraken", GetExchange())
	s.Equal(10., GetMinimumOrderSize())
}
