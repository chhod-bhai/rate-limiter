package model

type RateLimitResponse struct {
	RemainingLimit int
	TotalLimit     int
	RetryAfterSecs int
	LimitExpired   bool
}
