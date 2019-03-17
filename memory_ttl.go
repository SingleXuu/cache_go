package cache

import (
	"fmt"
	"sync"
	"time"
)

type MemoryTTL struct {
	sync.Mutex
	cache      *MemoryNoTs
	ttl        time.Duration
	setAts     map[string]time.Time
	gcInterval time.Duration
}

var zeroTTL = time.Duration(0)

func NewMemoryWithTTL(ttl time.Duration) *MemoryTTL {
	return &MemoryTTL{
		cache:      NewMemoryNoTs(),
		setAts:     map[string]time.Time{},
		gcInterval: ttl,
	}
}

func (r *MemoryTTL) StartGc(ttl time.Duration) {
	r.gcInterval = ttl

	go func() {
		for _ = range time.Tick(ttl) {
			fmt.Println("************gc*************")
			for key := range r.cache.items {
				if !r.vaildate(key) {
					r.delete(key)
				}
			}
		}
	}()
}

func (r *MemoryTTL) vaildate(key string) bool {
	at, ok := r.setAts[key]
	if !ok {
		return false
	}

	if r.ttl == zeroTTL {
		return true
	}

	return time.Now().Before(at)
}

func (r *MemoryTTL) delete(key string) error {
	err := r.Delete(key)
	delete(r.setAts, key)
	return err
}

func (r *MemoryTTL) Get(key string) (interface{}, error) {
	r.Lock()
	defer r.Unlock()

	return r.cache.Get(key)
}

func (r *MemoryTTL) Set(key string, value interface{}) error {

	r.Lock()
	defer r.Unlock()

	err := r.cache.Set(key, value)
	r.setAts[key] = time.Now().Add(r.ttl)
	return err
}
func (r *MemoryTTL) Delete(key string) error {

	r.Lock()
	defer r.Unlock()

	return r.cache.Delete(key)
}
