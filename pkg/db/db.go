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
