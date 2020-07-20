package server

import "sync"

type CacheDB struct {
	items map[string]string
	mu    sync.RWMutex
}

//NewCacheDB creates new cacheDB instance and returns them
func NewCacheDB() *CacheDB {

}
