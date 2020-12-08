package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"cloud.google.com/go/storage"
	"github.com/mshogin/randomtrader/pkg/config"
	"google.golang.org/api/option"
)

const saveObjectTimeout = time.Second * 50

var client *storage.Client

// Init ...
func Init() error {
	var err error
	ctx := context.Background()
	client, err = storage.NewClient(ctx, option.WithCredentialsFile(config.GetGCEServiceKeyFilepath()))
	if err != nil {
		return fmt.Errorf("cannot create gce storage client: %w", err)
	}

	return nil
}

// SaveObject ...
func SaveObject(prefix, fpath string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return fmt.Errorf("cannot open file %q: %w", fpath, err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(context.Background(), saveObjectTimeout)
	defer cancel()

	_, objectName := path.Split(fpath)
	wc := client.Bucket(config.GetGCEBucket()).Object(prefix + objectName).NewWriter(ctx)

	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("cannot copy file to the bucket: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("cannot close remote object: %v", err)
	}

	return nil
}
