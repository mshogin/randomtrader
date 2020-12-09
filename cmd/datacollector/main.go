package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/datacollector"
	"github.com/mshogin/randomtrader/pkg/logger"
)

func main() {
	logger.Infof("starting data collector")

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

	cancel, err := datacollector.Start()
	if err != nil {
		logger.Errorf("cannot run datacollector")
		cancel()
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	logger.Infof("data collector has been started")
	<-c

	logger.Infof("shutting down data collector...")
	cancel()
	logger.Infof("data collector has been stopped")
}
