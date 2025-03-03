package algorithms

import "github.com/chhod-bhai/rate-limiter/model"

type slidingWindowLog struct{}

// IsVaidRequest implements Algorithm.
func (l *slidingWindowLog) IsVaidRequest() (*model.RateLimitResponse, error) {
	panic("unimplemented")
}

// Replenish implements Algorithm.
func (l *slidingWindowLog) Replenish() error {
	panic("unimplemented")
}

func NewSlidingWindowLog() Algorithm {
	return &slidingWindowLog{}
}
