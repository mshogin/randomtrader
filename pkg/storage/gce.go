package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"github.com/mshogin/randomtrader/pkg/config"
	"github.com/mshogin/randomtrader/pkg/logger"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type gceClientImpl struct {
	cli *storage.Client
}

var gceClient Storage
var gceClientSync sync.Mutex

func SwapGCEClient(newClient Storage) Storage {
	gceClientSync.Lock()
	defer gceClientSync.Unlock()
	gceClientPrev := gceClient
	gceClient = newClient
	return gceClientPrev
}

// GetGCEClient ...
func GetGCEClient() (Storage, error) {
	gceClientSync.Lock()
	defer gceClientSync.Unlock()

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

func (m *gceClientImpl) DownloadObjects(prefix, targetDir string) error {
	ctx, cancel := context.WithTimeout(context.Background(), saveObjectTimeout)
	defer cancel()

	q := &storage.Query{Prefix: prefix}
	bucket := m.cli.Bucket(config.GetGCEBucket())
	it := bucket.Objects(ctx, q)

	var wg sync.WaitGroup

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects(): %v", "", err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := downloadFile(bucket, attrs.Name, targetDir); err != nil {
				logger.Errorf("cannot download file %q: %w", attrs.Name, err)
			}
		}()
	}

	wg.Wait()
	return nil
}

func downloadFile(bucket *storage.BucketHandle, objectName, targetDir string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	rc, err := bucket.Object(objectName).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("Object(%q).NewReader: %v", objectName, err)
	}
	defer rc.Close()

	fileName := strings.Split(objectName, "/")[2]
	fh, err := os.Create(filepath.Join(targetDir, fileName))
	if err != nil {
		return fmt.Errorf("cannot open file to store object from the bucket: %w", err)
	}

	if _, err = io.Copy(fh, rc); err != nil {
		return fmt.Errorf("cannot copy file from the bucket: %w", err)
	}

	if err := fh.Close(); err != nil {
		return fmt.Errorf("cannot close local object: %v", err)
	}

	return nil
}
