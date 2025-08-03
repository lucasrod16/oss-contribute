package cache

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	gistURL = "https://gist.githubusercontent.com/lucasrod16/dafa982abfa42982e02c75f1ddec46be/raw/data.json"
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
	resp, err := http.Get(gistURL)
	if err != nil {
		return fmt.Errorf("error fetching data from GitHub gist: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GitHub gist returned status %d: %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading GitHub data from gist: %w", err)
	}

	c.Set(data)
	log.Println("Successfully loaded GitHub data from gist into the in-memory cache.")
	return nil
}
