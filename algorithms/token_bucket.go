package algorithms

import "github.com/chhod-bhai/rate-limiter/model"

type tokenBucket struct{}

// IsVaidRequest implements Algorithm.
func (l *tokenBucket) IsVaidRequest() (*model.RateLimitResponse, error) {
	panic("unimplemented")
}

// Replenish implements Algorithm.
func (l *tokenBucket) Replenish() error {
	panic("unimplemented")
}

func NewTokenBucket() Algorithm {
	return &tokenBucket{}
}
