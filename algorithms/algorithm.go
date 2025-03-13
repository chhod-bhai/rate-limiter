package algorithms

import (
	"github.com/chhod-bhai/rate-limiter/model"
)

type Algorithm interface {
	IsVaidRequest() (*model.RateLimitResponse, error)
	Replenish() error
}

func New(algorithmType model.AlgorithmType) Algorithm {
	switch algorithmType {
	case model.LeakyBucket:
		return NewLeakyBucket()
	case model.TokenBucket:
		return NewTokenBucket()
	case model.SlidingWindowCounter:
		return NewSlidingWindowCounter()
	case model.SlidingWindowLog:
		return NewSlidingWindowLog()
	default:
		return NewTokenBucket()
	}
}
