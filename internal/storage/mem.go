package storage

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
)

// MemStorage data store mem
type MemStorage struct {
	sync.RWMutex
	cache map[string]string
}

// GetFeedData get feed from mem
func (m *MemStorage) GetFeedData(key string) (string, error) {
	m.RLock()
	data, exist := m.cache[key]
	m.RUnlock()
	if !exist {
		return "", fmt.Errorf("can not fin key %s", key)
	}
	return data, nil
}

// SaveFeedData save feed to mem
func (m *MemStorage) SaveFeedData(key string, data string) error {
	log.Debug().Msgf("store %s data", key)
	m.Lock()
	m.cache[key] = data
	m.Unlock()
	return nil
}

// NewMemStorage init mem storage
func NewMemStorage() FeedStorager {
	return &MemStorage{
		cache: make(map[string]string),
	}
}
