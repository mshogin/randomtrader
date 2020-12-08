package datacollector

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/exchange"
	"github.com/stretchr/testify/assert"
)

func TestOrderBookCollector(t *testing.T) {
	s := assert.New(t)

	tmpDir, err := ioutil.TempDir("", "")
	s.NoError(err)

	logFilename := "orderbook1s.log"
	oldConfig := config.SwapConfig(config.Configuration{
		LogsRoot: tmpDir,
		DataCollector: config.DataCollectorConfiguration{
			OrderBook: []config.DataCollectorOrderBook{
				{
					Filename: logFilename,
					Interval: 1,
				},
			},
		},
	})
	defer func() {
		config.SwapConfig(oldConfig)
	}()

	exchange.SetupTestGRPCClient()

	cancelDataCollector := Run()
	defer cancelDataCollector()
	time.Sleep(2 * time.Second) // give collector the time to collect at least once
	cancelDataCollector()

	fpath := path.Join(tmpDir, logFilename)
	file, err := os.Open(fpath)
	s.NoError(err)
	defer func() { s.NoError(file.Close()) }()

	lineNo := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lineNo++
		var ob map[string]interface{}
		s.NoError(json.Unmarshal(sc.Bytes(), &ob))
	}
	s.NoError(sc.Err())
	s.Greater(lineNo, 0)
}
