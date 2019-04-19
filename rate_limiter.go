package tatsumaki_go

import (
	"sync"
	"time"
)

// RateLimiter is the Tatsumaki API rate limiter.
// THE API allows 1 request every 0.2s.
type rateLimiter struct {
	lastRequest   time.Time
	resetInterval time.Duration
	sync.Mutex
}

// NewRateLimiter creates a new instance of the rate limiter.
func newRateLimiter() *rateLimiter {
	return &rateLimiter{
		resetInterval: 200 * time.Millisecond,
	}
}

func (rl *rateLimiter) acquire() {
	// Get current time.
	currentTime := time.Now()

	// If current time is greater than the last request + reset interval, allow request.
	if currentTime.After(rl.lastRequest.Add(rl.resetInterval)) {
		// Allow request.
		return
	}

	// Sleep for the time difference.
	time.Sleep((rl.lastRequest.Add(rl.resetInterval)).Sub(currentTime))
}
