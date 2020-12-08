package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/datacollector"
	"github.com/mshogin/randomtrader/pkg/logger"
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

	// trader.Run()
	datacollector.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	logger.Infof("Random trader has been started")
	<-c

	logger.Infof("Shutting down randomtrader...")
	// trader.Shutdown()
	logger.Infof("Random trader has been stopped")
}
