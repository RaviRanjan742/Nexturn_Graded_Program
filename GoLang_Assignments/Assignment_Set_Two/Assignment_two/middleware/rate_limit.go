package middleware

import (
	"net/http"
	"sync"
	"time"
)


type RateLimiter struct {
	requests      map[string]*ClientRequest
	mu            sync.Mutex
	requestLimit  int
	timeWindow    time.Duration
}


type ClientRequest struct {
	timestamp time.Time
	count     int
}


func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests:     make(map[string]*ClientRequest),
		requestLimit: limit,
		timeWindow:   window,
	}
}


func (rl *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr 

		rl.mu.Lock()
		defer rl.mu.Unlock()

		now := time.Now()
		clientRequest, exists := rl.requests[clientIP]

		if !exists || now.Sub(clientRequest.timestamp) > rl.timeWindow {
			
			rl.requests[clientIP] = &ClientRequest{
				timestamp: now,
				count:     1,
			}
		} else {
			
			clientRequest.count++
			if clientRequest.count > rl.requestLimit {
				
				w.Header().Set("Retry-After", rl.timeWindow.String())
				http.Error(w, "Rate limit exceeded. Try again later.", http.StatusTooManyRequests)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
