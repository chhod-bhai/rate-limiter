package model

type AlgorithmType int

const (
	TokenBucket AlgorithmType = iota
	LeakyBucket
	SlidingWindowCounter
	SlidingWindowLog
)
