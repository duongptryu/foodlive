package mycache

import (
	"errors"
	"github.com/bluele/gcache"
	"time"
)

type Cache interface {
	Set(key string, value interface{}) error
	SetWithExpire(key string, value interface{}, expire time.Duration) error
	Get(key string) (interface{}, error)
	Remove (key string) error
}

type cache struct {
	gCache gcache.Cache
}

func NewMyCache() *cache {
	gCache := gcache.New(20).
		LRU().
		Build()
	return &cache{
		gCache,
	}
}

func (mc *cache) Set(key string, value interface{}) error {
	err := mc.gCache.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (mc *cache) SetWithExpire(key string, value interface{}, expire time.Duration) error {
	err := mc.gCache.SetWithExpire(key, value, expire)
	if err != nil {
		return err
	}
	return nil
}

func (mc *cache) Get(key string) (interface{}, error) {
	data, err := mc.gCache.Get(key)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (mc *cache) Remove(key string) error {
	ok := mc.gCache.Remove(key)
	if ok {
		return nil
	}
	return errors.New("Cannot remove in cache")
}
