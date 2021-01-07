package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/datacollector"
	"github.com/mshogin/randomtrader/pkg/exchange"
	"github.com/mshogin/randomtrader/pkg/logger"
	"github.com/mshogin/randomtrader/pkg/storage"
	"github.com/mshogin/randomtrader/pkg/strategy"
	"github.com/mshogin/randomtrader/pkg/trader"
)

func main() {
	logger.Infof("Starting Random trader")

	configPath := flag.String(
		"config",
		"/etc/randomtrader/config.json",
		"Path to configuration file")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	if _, err := config.Init(*configPath); err != nil {
		logger.Errorf("can't initialize configuration: %s", err)
		os.Exit(1)
	}

	if config.IsTestModeEnabled() {
		exchange.SetupTestGRPCClient()
		storage.SwapGCEClient(storage.GetGCETestClient())
	}

	if err := strategy.Init(); err != nil {
		logger.Errorf("can't initialize strategy: %s", err)
		os.Exit(1)
	}

	if config.IsDataCollectorEnabled() {
		if err := datacollector.Start(); err != nil {
			logger.Errorf("cannot run datacollector")
			datacollector.Stop()
			os.Exit(1)
		}
	}

	if config.IsTraderEnabled() {
		trader.Run()
	}

	go processReload(*configPath)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	logger.Infof("Random trader has been started")
	<-c
	logger.Infof("Shutting down randomtrader...")

	if config.IsTraderEnabled() {
		trader.Shutdown()
	}

	if config.IsDataCollectorEnabled() {
		datacollector.Stop()
	}

	logger.Infof("Random trader has been stopped")
}

// processReload ...
func processReload(configPath string) {
	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP)

		<-c

		if config.IsDataCollectorEnabled() {
			if err := datacollector.Reload(configPath); err != nil {
				logger.Fatalf("cannot reload datacollector: %w", err)
			} else {
				logger.Infof("datacollector reloaded successfully")
			}
		}
	}
}
