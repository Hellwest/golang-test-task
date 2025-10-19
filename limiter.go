package main

import (
	"time"
)

// RateLimiter limits client to send n items in p period
type RateLimiter struct {
	n uint64
	p time.Duration

	count    uint64
	lastTick time.Time
}

func NewRateLimiter(n uint64, p time.Duration) *RateLimiter {
	return &RateLimiter{
		n:        n,
		p:        p,
		lastTick: time.Now(),
	}
}

// Wait until it's safe to send next batch
func (r *RateLimiter) Wait(batchSize uint64) {
	for {
		now := time.Now()

		if now.Sub(r.lastTick) >= r.p {
			r.count = 0
			r.lastTick = now
		}

		if r.count+batchSize <= r.n {
			r.count += batchSize
			return
		}

		wait := r.p - now.Sub(r.lastTick)
		time.Sleep(wait)
	}
}
