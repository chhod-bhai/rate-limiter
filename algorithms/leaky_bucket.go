package algorithms

import "github.com/chhod-bhai/rate-limiter/model"

type leakyBucket struct{}

// IsVaidRequest implements Algorithm.
func (l *leakyBucket) IsVaidRequest() (*model.RateLimitResponse, error) {
	panic("unimplemented")
}

// Replenish implements Algorithm.
func (l *leakyBucket) Replenish() error {
	panic("unimplemented")
}

func NewLeakyBucket() Algorithm {
	return &leakyBucket{}
}
