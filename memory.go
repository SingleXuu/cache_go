package cache

import "sync"

type Memory struct {
	sync.Mutex
	cache Cache
}

func NewMemory() Cache {
	return &Memory{
		cache: NewMemoryNoTs(),
	}
}

func (r *Memory) Get(key string) (interface{}, error) {
	r.Lock()
	defer r.Unlock()
	return r.cache.Get(key)
}

//
func (r *Memory) Set(key string, value interface{}) error {

	r.Lock()
	defer r.Unlock()
	return r.cache.Set(key, value)
}

//
func (r *Memory) Delete(key string) error {

	r.Lock()
	defer r.Unlock()
	return r.cache.Delete(key)
}
