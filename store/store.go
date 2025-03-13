package store

import (
	"sync"
	"time"
)

type Store interface {
	Set(key string, value int, expiry time.Duration) error
	Get(key string) (int, error)
	Increment(key string) (int, error)
	Decrement(key string) (int, error)
	Delete(key string) error
	SetEx(key string, expiry time.Duration) error
	GetTTL(key string) (time.Duration, error)
}

var once sync.Once
var db Store

func Get() Store {
	once.Do(func() {
		db = NewRedisStore()
	})
	return db
}
