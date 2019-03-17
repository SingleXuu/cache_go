package cache

import "sync"

type LRU struct {
	sync.Mutex
	cache Cache
}

func (r *LRU) Get(key string) (interface{}, error) {
	r.Lock()
	defer r.Unlock()
	return r.cache.Get(key)
}

func (r *LRU) Set(key string, value interface{}) error {
	r.Lock()
	defer r.Unlock()
	return r.cache.Set(key, value)
}

func (r *LRU) Delete(key string) error {
	r.Lock()
	defer r.Unlock()
	return r.cache.Delete(key)
}
