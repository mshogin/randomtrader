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

	if _, err := config.Init(*configPath); err != nil {
		logger.Errorf("can't initialise configuration: %s", err)
		os.Exit(1)
	}

	if err := datacollector.Start(); err != nil {
		logger.Errorf("cannot run datacollector")
		datacollector.Stop()
		os.Exit(1)
	}

	go processReload(*configPath)

	Run()
}

func Run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	logger.Infof("data collector has been started")
	<-c

	logger.Infof("shutting down data collector...")
	datacollector.Stop()
	logger.Infof("data collector has been stopped")
}

// processReload ...
func processReload(configPath string) {
	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP)
		<-c
		if err := datacollector.Reload(configPath); err != nil {
			logger.Fatalf("cannot reload datacollector: %w", err)
		} else {
			logger.Infof("datacollector reloaded successfully")
		}
	}
}
