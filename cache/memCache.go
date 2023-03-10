package cache

import "time"

type memCache struct {
}

func (mc *memCache) SetMaxMemory(size string) bool {

	return false
}

func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	return false
}

func (mc *memCache) Get(key string) (interface{}, bool)

func (mc *memCache) Del(key string) bool

func (mc *memCache) Exists(key string) bool

func (mc *memCache) Flush() bool

func (mc *memCache) Keys() int64
