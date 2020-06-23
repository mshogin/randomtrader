package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/logger"
	"github.com/mshogin/randomtrader/pkg/trader"
)

func main() {
	logger.Info("Starting Random trader")

	configPath := flag.String(
		"config",
		"/etc/randomtrader/config.json",
		"Path to configuration file")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	if err := config.Init(*configPath); err != nil {
		logger.Error("can't initialise configuration: %s", err)
		os.Exit(1)
	}

	if config.IsDebugEnabled() {
		logger.EnableDebug()
	}

	trader.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	logger.Debug("Random trader has been started")
	<-c

	logger.Debug("Shutting down randomtrader...")
	trader.Shutdown()
	logger.Debug("Random trader has been stopped")
}
