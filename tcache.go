// Package cache implements a simple goroutine safe cache with expiration time
// accepts any value but it only gives back valid values ( not expired once ).
// Provides only 3 methods, Put, Get and Stop ( stop must be called when done with
// the timed cache to avoid memory leakage ) and a constructor New.

package tcache

import (
	"sync"
	"time"
)

//Cache is the structure the implements the time cache
type Cache struct {
	values *sync.Map
	tick   *time.Ticker
	done   chan bool
	exp    time.Duration
}

type cacheObject struct {
	Value interface{}
	Valid time.Time
}

// NewWithDefault creates a default valued timed cache with value purging set
// to run every 3 minutes
func NewWithDefault(exp time.Duration) *Cache {
	return New(3, exp)
}

// New creates a new cache with minutes, which rappresent a interval at which
// old values are deleted and exp ( in minutes as well ) which sets the expiration
// time for values ( cannot be change once the cache has been instantiated ).
func New(minutes time.Duration, exp time.Duration) *Cache {
	cache := &Cache{
		&sync.Map{},
		time.NewTicker(time.Minute * minutes),
		make(chan bool, 1),
		exp,
	}
	cache.purger()

	return cache
}

// Put adds a values to the cache
func (c *Cache) Put(key string, v interface{}) {
	c.values.Store(key, &cacheObject{v, time.Now()})
}

// Get give you back the value assuming it hasn't be purged yet
func (c *Cache) Get(key string) (interface{}, bool) {

	cacheEntry, ok := c.values.Load(key)
	if !ok {
		return nil, ok
	}

	return cacheEntry.(*cacheObject).Value, ok
}

// Stop Must be called otherwise the cache will leak
func (c *Cache) Stop() {
	c.done <- true
}

func (c *Cache) purger() {
	go func() {
		for {
			select {
			case <-c.done:
				c.tick.Stop()
				return

			case now := <-c.tick.C:
				c.cleaner(now)
			}
		}
	}()
}

func (c *Cache) cleaner(now time.Time) {
	c.values.Range(func(key, value interface{}) bool {
		t := value.(*cacheObject)

		if now.After(t.Valid.Add(c.exp)) {
			c.values.Delete(key)
		}

		return true
	})
}
