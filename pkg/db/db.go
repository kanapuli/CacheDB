package db

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

//Delete removes a key from CacheDB
func (c *CacheDB) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

//Save saves the in-memory key value pairs to a persistent file
func (c *CacheDB) Save(file string) error {
	file, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening persistent file %v to save", file)
		return error
	}
	err = json.NewEncoder(f).Encode(c.items)
	if err != nil {
		fmt.Printf("Got error while saving key values to persistent file %v. Err: %v", file, err)
	}
	return err
}
