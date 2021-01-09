package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestSaveOrderBookLog(t *testing.T) {
	s := assert.New(t)

	gceClientOrig := SwapGCEClient(GetGCETestClient())
	defer func() { SwapGCEClient(gceClientOrig) }()

	s.NoError(SaveOrderBookLog("/some/path"))
}

func TestDownloadLogs(t *testing.T) {
	s := assert.New(t)

	cliOrig := SwapGCEClient(GetGCETestClient())
	defer SwapGCEClient(cliOrig)

	// co := setupTestConfig()
	// defer config.SwapConfig(co)

	for _, t := range []time.Time{
		time.Now(),
		time.Now().Add(-time.Hour * 24),
	} {
		prefix := fmt.Sprintf(
			"%s/%s/orderbook1h",
			orderBookBucketPrefix,
			t.Format(layoutISO))

		err := DownloadLogs(prefix, "./tmp")
		s.NoError(err)
	}
}

func setupTestConfig() config.Configuration {
	cfg := config.Configuration{
		ConfigsRoot:        "/home/mshogin/randomtrader/dcf",
		ServiceKeyFilename: "gce-bucket-service-key.json",
		GCEBucket:          "randomtrader-datacollector",
	}

	return config.SwapConfig(cfg)
}
