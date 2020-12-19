package datacollector

import (
	"context"
	"fmt"
	"time"

	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/exchange"
	"github.com/mshogin/randomtrader/pkg/logger"
)

var shutdown context.CancelFunc

// Start ...
var Start = func() error {
	var ctx context.Context
	ctx, shutdown = context.WithCancel(context.Background())
	c := config.GetDataCollector()

	if err := startOrderBookDumper(ctx, c.OrderBook); err != nil {
		return fmt.Errorf("cannot start order book collector: %w", err)
	}

	return nil
}

// Stop ...
var Stop = func() {
	shutdown()
}

// Reload ...
func Reload(configPath string) error {
	oldConfig, err := config.Init(configPath)
	if err != nil {
		config.SwapConfig(oldConfig)
		return fmt.Errorf("cannot reload config: %w", err)
	}

	Stop()
	if err := Start(); err != nil {
		return fmt.Errorf("cannot start datacollector: %w", err)
	}

	return nil
}

func startOrderBookDumper(ctx context.Context, logConfigs []config.OrderBookLog) error {
	if logConfigs == nil {
		return fmt.Errorf("cannot run order book collector: empty config")
	}
	for _, logConfig := range logConfigs {
		go runOrderBookDumper(ctx, logConfig)
	}
	return nil
}

func runOrderBookDumper(ctx context.Context, conf config.OrderBookLog) {
	log, err := createLogFile(conf)
	if err != nil {
		panic(fmt.Errorf("cannot create log file %q: %w", conf.GetFilepath(), err))
	}
	ticker := time.NewTicker(time.Duration(conf.DumpInterval) * time.Second)

	running := true
	for running {
		select {
		case <-ctx.Done():
			running = false
		case <-ticker.C:
			ob, err := exchange.GetOrderBook()
			if err != nil {
				logger.Errorf("cannot get order book: %w", err)
			} else {
				if err := log.Write(ob); err != nil {
					logger.Errorf("cannot write order book: %w", err)
				}
			}
		}
	}
}
