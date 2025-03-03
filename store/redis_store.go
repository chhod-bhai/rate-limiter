package store

type redisStore struct{}

func NewRedisStore() Store {
	return &redisStore{}
}

func (rs *redisStore) Set(key string, value int) error {
	return nil
}

func (rs *redisStore) Get(key string) (int, error) {
	return 0, nil
}

func (rs *redisStore) Increment(key string) (int, error) {
	return 0, nil
}

func (rs *redisStore) Decrement(key string) (int, error) {
	return 0, nil
}

func (rs *redisStore) Delete(key string) error {
	return nil
}

func (rs *redisStore) SetEx(key string, expiry int) error {
	return nil
}
