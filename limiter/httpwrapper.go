package limiter

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"net/http"
)

func WrapHttpFunc(lmt *limiter.Limiter, httpFunc http.HandlerFunc) http.HandlerFunc {
	return tollbooth.LimitFuncHandler(lmt, httpFunc).ServeHTTP
}
