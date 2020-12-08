package datacollector

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/logger"
	"github.com/mshogin/randomtrader/pkg/storage"
)

const (
	orderBookBucketPrefix = "order-book"
	layoutISO             = "2006-01-02"
	rotationSuffix        = ".%Y%m%d%H%M%S"
)

type logFile struct {
	rl *rotatelogs.RotateLogs
}

func createLogFile(conf config.OrderBookLog) (*logFile, error) {
	rl, err := rotatelogs.New(
		conf.GetFilepath()+rotationSuffix,
		rotatelogs.WithLinkName(conf.GetFilepath()),
		rotatelogs.WithRotationTime(conf.GetRotateInterval()),
		rotatelogs.WithHandler(
			rotatelogs.HandlerFunc(func(e rotatelogs.Event) {
				if e.Type() != rotatelogs.FileRotatedEventType {
					return
				}

				prevFilePath := e.(*rotatelogs.FileRotatedEvent).PreviousFile()
				if len(prevFilePath) != 0 {
					prefix := fmt.Sprintf("%s/%s/", orderBookBucketPrefix, time.Now().Format(layoutISO))
					if err := storage.SaveObject(prefix, prevFilePath); err != nil {
						logger.Errorf("cannot send %q to gce bucket: %w", prevFilePath, err)
					}

					if err := os.Remove(prevFilePath); err != nil {
						logger.Errorf("cannot remove %q: %w", prevFilePath, err)
					}
				}
			}),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot create rotatelog %q: %w", conf.GetFilepath(), err)
	}

	return &logFile{rl}, nil
}

// Write ...
func (m *logFile) Write(item interface{}) error {
	buf, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("cannot marshal order book to JSON: %w", err)
	}

	if _, err := m.rl.Write(buf); err != nil {
		return fmt.Errorf("cannot write buffer to file: %w", err)
	}

	if _, err := m.rl.Write([]byte("\n")); err != nil {
		return fmt.Errorf("cannot write buffer to file: %w", err)
	}

	return nil
}
