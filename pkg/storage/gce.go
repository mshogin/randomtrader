package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"

	"cloud.google.com/go/storage"
	"github.com/mshogin/randomtrader/pkg/config"
	"google.golang.org/api/option"
)

type gceClientImpl struct {
	cli *storage.Client
}

var gceClient *gceClientImpl

// GetGCEClient ...
var GetGCEClient = func() (Storage, error) {
	if gceClient != nil {
		return gceClient, nil
	}
	ctx := context.Background()
	cli, err := storage.NewClient(ctx, option.WithCredentialsFile(config.GetGCEServiceKeyFilepath()))
	if err != nil {
		return nil, fmt.Errorf("cannot create gce storage client: %w", err)
	}
	gceClient = &gceClientImpl{cli: cli}
	return gceClient, nil
}

// SaveObject ...
func (m *gceClientImpl) SaveObject(prefix, fpath string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return fmt.Errorf("cannot open file %q: %w", fpath, err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(context.Background(), saveObjectTimeout)
	defer cancel()

	_, objectName := path.Split(fpath)
	wc := m.cli.Bucket(config.GetGCEBucket()).Object(prefix + objectName).NewWriter(ctx)

	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("cannot copy file to the bucket: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("cannot close remote object: %v", err)
	}

	return nil
}
