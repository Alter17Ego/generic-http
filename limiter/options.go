package limiter

import (
	"github.com/didip/tollbooth/limiter"
)

type setter func(*limiter.Limiter)

func UseProxyIpLookups() setter {
	return func(lmt *limiter.Limiter) {
		lmt.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr", "X-Real-IP"})
	}
}

func UseProductionConfig() setter {
	return func(lmt *limiter.Limiter) {
		UseProxyIpLookups()(lmt)
	}
}

func applySetters(lmt *limiter.Limiter, setters ...setter) {
	for _, apply := range setters {
		apply(lmt)
	}
}
