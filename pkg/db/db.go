package server

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

//CacheDB contains a map of string to string to store the incoming items
type CacheDB struct {
	items map[string]string
	mu    sync.RWMutex
}

const (
	persistentFile = "cachedb.json"
)

//NewCacheDB creates new cacheDB instance and returns them
func NewCacheDB() *CacheDB {
	// Try to open persistent file
	file, err := os.Open(persistentFile)
	if err != nil {
		fmt.Printf("Unable to open the persistent file %v", persistentFile)
		return &CacheDB{
			items: make(map[string]string),
		}
	}

	oldKVPairs := make(map[string]string)
	err = json.NewDecoder(file).Decode(&oldKVPairs)
	if err != nil {
		fmt.Printf("Failed to decode the persistent file %v", persistentFile)
		return &CacheDB{
			items: make(map[string]string),
		}
	}

	return &CacheDB{
		items: oldKVPairs,
	}
}

//Get retrieves a value stored in the CacheDB
func (c *CacheDB) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, found := c.items[key]
	return value, found
}

//Set adds a key value pair to Cachedb
func (c *CacheDB) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value

}
