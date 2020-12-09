package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/logger"
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

	if err := config.Init(*configPath); err != nil {
		logger.Errorf("can't initialise configuration: %s", err)
		os.Exit(1)
	}

	if config.IsDebugEnabled() {
		logger.EnableDebug()
	}

	trader.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	logger.Infof("Random trader has been started")
	<-c

	logger.Infof("Shutting down randomtrader...")
	trader.Shutdown()
	logger.Infof("Random trader has been stopped")
}
