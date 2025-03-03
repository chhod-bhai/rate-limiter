package algorithms

import "github.com/chhod-bhai/rate-limiter/model"

type slidingWindowCounter struct{}

// IsVaidRequest implements Algorithm.
func (l *slidingWindowCounter) IsVaidRequest() (*model.RateLimitResponse, error) {
	panic("unimplemented")
}

// Replenish implements Algorithm.
func (l *slidingWindowCounter) Replenish() error {
	panic("unimplemented")
}

func NewSlidingWindowCounter() Algorithm {
	return &slidingWindowCounter{}
}
