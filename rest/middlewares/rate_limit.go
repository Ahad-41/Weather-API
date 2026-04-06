package middleware

import (
	"net/http"
	"sync"
	"time"
	"weather-api/config"

	"golang.org/x/time/rate"
)

type IpRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

func NewIpRateLimiter(r rate.Limit, b int) *IpRateLimiter {
	return &IpRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
}

func (i *IpRateLimiter) AddIp(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

func (i *IpRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.RLock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.RUnlock()
		return i.AddIp(ip)
	}

	i.mu.RUnlock()

	return limiter
}

func RateLimitMiddleware(cnf *config.Config) Middleware {
	limiter := NewIpRateLimiter(rate.Every(time.Duration(cnf.RateLimitDurationSeconds)*time.Second/time.Duration(cnf.RateLimitRequests)), cnf.RateLimitRequests)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if !limiter.GetLimiter(ip).Allow() {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
