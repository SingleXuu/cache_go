package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestTTL(t *testing.T) {

	cache := NewMemoryWithTTL(time.Second * 10)
	cache.StartGc(time.Millisecond * 10)

	cache.Set("name", "冬天")
	time.Sleep(time.Second * 5)

	va, err := cache.Get("name")
	fmt.Println("v=", va, "err=", err)
}
