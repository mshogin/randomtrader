package storage

import (
	"fmt"
	"time"
)

const (
	saveObjectTimeout     = time.Second * 50
	orderBookBucketPrefix = "order-book"
	layoutISO             = "2006-01-02"
)

// Storage ...
type Storage interface {
	SaveObject(string, string) error
	DownloadObjects(string, string) error
}

// SaveOrderBookLog ...
func SaveOrderBookLog(fpath string) error {
	c, err := GetGCEClient()
	if err != nil {
		return fmt.Errorf("cannot create gce client: %w", err)
	}
	return c.SaveObject(
		fmt.Sprintf("%s/%s/", orderBookBucketPrefix, time.Now().Format(layoutISO)),
		fpath,
	)
}

func DownloadLogs(prefix string, dest string) error {
	c, err := GetGCEClient()
	if err != nil {
		return fmt.Errorf("cannot create gce client: %w", err)
	}
	return c.DownloadObjects(prefix, dest)
}
