package cache

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"cloud.google.com/go/storage"
)

const (
	jsonFile = "data.json"
	bucket   = "lucasrod16-github-data"
)

type Cache struct {
	data      []byte
	timestamp string
	mutex     sync.RWMutex
}

func New() *Cache {
	return &Cache{}
}

func (c *Cache) Set(data []byte) {
	c.mutex.Lock()
	c.data = data
	c.timestamp = time.Now().UTC().Format(http.TimeFormat)
	c.mutex.Unlock()
}

func (c *Cache) Get() ([]byte, string) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.data, c.timestamp
}

func (c *Cache) RepoData(ctx context.Context) error {
	gcsClient, err := storage.NewClient(ctx, storage.WithJSONReads())
	if err != nil {
		return fmt.Errorf("failed to create GCS client: %w", err)
	}
	defer gcsClient.Close()

	r, err := gcsClient.Bucket(bucket).Object(jsonFile).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("error creating GCS reader: %w", err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("error reading GitHub data from GCS bucket %q: %w", bucket, err)
	}

	c.Set(data)
	log.Println("Successfully loaded GitHub data from GCS bucket into the in-memory cache.")
	return nil
}
