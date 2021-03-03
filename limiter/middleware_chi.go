package limiter

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

func NewChiMiddleware(lmt *limiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return tollbooth.LimitHandler(lmt, next)
	}
}
