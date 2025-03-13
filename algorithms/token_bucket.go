package algorithms

import (
	"fmt"
	"time"

	"github.com/chhod-bhai/rate-limiter/model"
	"github.com/chhod-bhai/rate-limiter/store"
)

type tokenBucket struct{}

// IsVaidRequest implements Algorithm.
func (l *tokenBucket) IsVaidRequest() (*model.RateLimitResponse, error) {
	// decrement token bucket , if output is zero return failed, else succes
	storeInstance := store.Get()
	// set retry after time
	ttl, err := storeInstance.GetTTL(l.getKey())
	if err != nil {
		println(fmt.Sprintf("ERROR: fetching expiry ttl: %v", err))
		return nil, err
	}

	resp := model.RateLimitResponse{
		RemainingLimit: 0,
		TotalLimit:     100,
		RetryAfterSecs: int(ttl.Seconds()),
		LimitExpired:   false,
	}

	oldVal, err := storeInstance.Get(l.getKey())
	if err != nil {
		return nil, err
	}
	if oldVal == 0 {
		resp.LimitExpired = true
		return &resp, nil
	}
	newVal, err := storeInstance.Decrement(l.getKey())
	if err != nil {
		println(fmt.Sprintf("ERROR: error decrementing store value: %v", err))
		return nil, err
	}
	resp.RemainingLimit = newVal
	return &resp, nil
}

func (l *tokenBucket) getKey() string {
	return "token_bucket"
}

// Replenish implements Algorithm.
func (l *tokenBucket) Replenish() error {
	// set the bucket to full
	storeInstance := store.Get()
	return storeInstance.Set(l.getKey(), 100, time.Duration(time.Minute))
}

func NewTokenBucket() Algorithm {
	return &tokenBucket{}
}
