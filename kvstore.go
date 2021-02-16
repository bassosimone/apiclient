package apiclient

import (
	"errors"
	"fmt"
	"sync"
)

// KVStore is a key-value store.
type KVStore interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

var errMemkvstoreNotFound = errors.New("apiclient: memkvstore: not found")

type memkvstore struct {
	m  map[string][]byte
	mu sync.Mutex
}

func (kvs *memkvstore) Get(key string) ([]byte, error) {
	defer kvs.mu.Unlock()
	kvs.mu.Lock()
	out, good := kvs.m[key]
	if !good {
		return nil, fmt.Errorf("%w: %s", errMemkvstoreNotFound, key)
	}
	return out, nil
}

func (kvs *memkvstore) Set(key string, value []byte) error {
	defer kvs.mu.Unlock()
	kvs.mu.Lock()
	kvs.m[key] = value
	return nil
}

var defaultKVStore KVStore = &memkvstore{}

// kvstore returns the configured KVStore or a default
// memory-based, ephemeral KVStore instance.
func (c *Client) kvstore() KVStore {
	kvstore := defaultKVStore
	if c.KVStore != nil {
		kvstore = c.KVStore
	}
	return kvstore
}
