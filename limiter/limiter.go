package limiter

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

type Setting struct {
	MaxRequestsPerSecond float64

	// How frequently expire job triggers
	ExpirationIn      time.Duration
	ExpireJobInterval time.Duration
}

func New(setting *Setting, setters ...setter) *limiter.Limiter {
	lmt := tollbooth.NewLimiter(setting.MaxRequestsPerSecond,
		&limiter.ExpirableOptions{
			DefaultExpirationTTL: setting.ExpirationIn,
			ExpireJobInterval:    setting.ExpireJobInterval,
		})
	applySetters(lmt, setters...)
	return lmt
}
