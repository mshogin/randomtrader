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
	"github.com/mshogin/randomtrader/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestOrderBookCollector(t *testing.T) {
	s := assert.New(t)

	tmpDir, err := ioutil.TempDir("", "")
	s.NoError(err)

	logFilename := "orderbook1s.log"
	oldConfig := config.SwapConfig(config.Configuration{
		LogsRoot: tmpDir,
		DataCollector: config.DataCollector{
			OrderBook: []config.OrderBookLog{
				{
					Filename:       logFilename,
					DumpInterval:   1,
					RotateInterval: 1,
				},
			},
		},
	})
	defer func() {
		config.SwapConfig(oldConfig)
	}()

	exchange.SetupTestGRPCClient()

	GetGCEClientOrig := storage.SwapGCEClient(storage.GetGCETestClient())
	defer storage.SwapGCEClient(GetGCEClientOrig)

	s.NoError(Start())
	defer Stop()

	time.Sleep(3 * time.Second) // give collector the time to collect at least once
	Stop()

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

func TestReload(t *testing.T) {
	s := assert.New(t)
	startOrig := Start
	stopOrig := Stop
	defer func() {
		Start = startOrig
		Stop = stopOrig
	}()

	startCount, stopCount := 0, 0
	Start = func() error {
		startCount++
		return nil
	}
	Stop = func() {
		stopCount++
	}

	f, err := ioutil.TempFile("", "")
	s.NoError(err)
	s.NoError(f.Close())

	s.NoError(ioutil.WriteFile(f.Name(), []byte("{}"), os.FileMode(644)))

	Reload(f.Name())
	defer Stop()

	s.Equal(1, startCount)
	s.Equal(1, stopCount)
}
