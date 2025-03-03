package store

type Store interface {
	Set(key string, value int) error
	Get(key string) (int, error)
	Increment(key string) (int, error)
	Decrement(key string) (int, error)
	Delete(key string) error
	SetEx(key string, expiry int) error
}
