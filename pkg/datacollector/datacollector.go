package datacollector

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/exchange"
	"github.com/mshogin/randomtrader/pkg/logger"
)

func Run() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	c := config.GetDataCollector()
	go collectOrderBook(ctx, c.OrderBook)
	return cancel
}

func collectOrderBook(ctx context.Context, conf []config.DataCollectorOrderBook) {
	if conf == nil {
		logger.Errorf("cannot run order book collector: empty config")
	}

	for _, logConfig := range conf {
		go runOrderBookCollector(ctx,
			logConfig.GetFilepath(),
			time.Duration(logConfig.Interval)*time.Second,
		)
	}
}

// runOrderBookCollector ...
func runOrderBookCollector(ctx context.Context, fpath string, d time.Duration) {
	ticker := time.NewTicker(d)

	done := false
	for !done {
		select {
		case <-ctx.Done():
			done = true
		case <-ticker.C:
			ob, err := exchange.GetOrderBook()
			if err != nil {
				logger.Errorf("cannot get orderbook: %w", err)
			}
			if err := createLogRecord(fpath, ob); err != nil {
				logger.Errorf("cannot create log record in %q: %w", fpath, err)
			}
		}
	}
}

var logFiles = map[string]io.Writer{}

// getLogFile ...
func getLogFile(fpath string) (io.Writer, error) {
	var f io.Writer
	var ok bool
	if f, ok = logFiles[fpath]; !ok {
		var err error
		f, err = rotatelogs.New(
			fpath+".%Y%m%d%H%M%S",
			rotatelogs.WithLinkName(fpath),
			rotatelogs.WithRotationTime(5*time.Second),
			rotatelogs.WithHandler(
				rotatelogs.HandlerFunc(func(e rotatelogs.Event) {
					if e.Type() != rotatelogs.FileRotatedEventType {
						return
					}

					prevFilePath := e.(*rotatelogs.FileRotatedEvent).PreviousFile()
					if len(prevFilePath) == 0 {
						return
					}
					if err := os.Remove(prevFilePath); err != nil {
						logger.Errorf("cannot remove %q: %w", prevFilePath, err)
					}
					fmt.Printf("removed %+v\n", prevFilePath) // output for debug
				}),
			),
		)
		if err != nil {
			return nil, fmt.Errorf("cannot create rotatelog %q: %w", fpath, err)
		}
		logFiles[fpath] = f
	}

	return f, nil
}

// createLogRecord ...
func createLogRecord(fpath string, ob *exchange.OrderBook) error {
	fh, err := getLogFile(fpath)
	if err != nil {
		return fmt.Errorf("cannot get log file %q: %w", fpath, err)
	}

	buf, err := json.Marshal(ob)
	if err != nil {
		return fmt.Errorf("cannot marshal order book to JSON: %w", err)
	}

	if _, err := fh.Write(buf); err != nil {
		return fmt.Errorf("cannot write buffer to file: %w", err)
	}

	if _, err := fh.Write([]byte("\n")); err != nil {
		return fmt.Errorf("cannot write buffer to file: %w", err)
	}

	return nil

}
