# Generic-HTTP

## Overview
This library will improve your http transport layer.

## How To Use

### Rendering HTTP response

#### Rendering HTTP Error Message
```
import (
    ...
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"github.com/Alter17Ego/generic-http/responses/errorresp"
)

/// Render CodedError as an error response
func HelloEndpoint(w http.ResponseWriter, r *http.Request) {
   ...
   if err := r.userService.Modify(userSession,profile); err != nil {
     if err == user.MissMatchUserErr {
         resp := errorresp.Forbidden(err) // 403
         render.Render(w,r,resp)
      }else if err == user.UnkownErr {
         resp := errorresp.InternalServerError(err) // 500
         render.Render(w,r,resp)
      } else if err == user.InvalidParamterErr {
         resp := errorresp.BadRequest(err) // 400
         render.Render(w,r,resp)
      } else if err == user.SessionExpiredErr {
         resp := errorresp.Unauthorized(err) // 401
         render.Render(w,r,resp)
      }
   }
   ...
}

/// Wrap go error type then render as an error response
func HelloEndpoint(w http.ResponseWriter, r *http.Request) {
 if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		wrappedErr := codederrors.Wrap(err, "-1", "something went wrong while decoding request")
		resp := errorresp.InternalServerError(wrappedErr)
		render.Render(w, r, resp)
		return
	}
}
```

#### Rendering Responses
```
import (
    ...
    "github.com/Alter17Ego/generic-http/responses"
)

type HelloResponse struct {
	*responses.StatusCodedResponse
	Message string `json:"message"`
}

func newHelloResponse(msg string) *HelloResponse {
	return &HelloResponse{Message: msg, StatusCodedResponse: responses.New(200)}
}

func HelloEndpoint(w http.ResponseWriter, r *http.Request) {
	var requestBody HelloRequest
    ...
	render.Render(w, r, newHelloResponse(fmt.Sprint("Hello", " ", requestBody.Name)))
	return
}

```

### Rate limiter

#### Use pre defined config
```
import (
	...
	"github.com/Alter17Ego/generic-http/limiter"
)

func main() {
	router := chi.NewRouter()

	router.Use(limiter.NewChiMiddleware(limiter.
		New(&limiter.
			Setting{
			ExpirationIn:         time.Hour,
			MaxRequestsPerSecond: 1, // limit request to 1 rps
		},
			limiter.UseProductionConfig())))

	router.Post("/hello", HelloEndpoint)
	router.Post("/hello2", HelloEndpoint)
	http.ListenAndServe(":8080", router)
}
```

#### Customizing via setter
```
import (
	...
	"github.com/Alter17Ego/generic-http/limiter"
)

func main() {
	router := chi.NewRouter()

	router.Use(limiter.NewChiMiddleware(limiter.
		New(&limiter.
			Setting{
			ExpirationIn:         time.Hour,
			MaxRequestsPerSecond: 1, // limit request to 1 rps
		},
            //manually config via setter
			limiter.UseProxyIpLookups(),
            limiter.Use...
            )))

	router.Post("/hello", HelloEndpoint)
	router.Post("/hello2", HelloEndpoint)
	http.ListenAndServe(":8080", router)
}
```

#### Injecting tollbolt
```
import (
	...
	"github.com/Alter17Ego/generic-http/limiter"
)

...
    router := chi.NewRouter()
    tollboltLimitter := limitter.New(...)
    ...
	router.Use(limiter.NewChiMiddleware(tollboltLimitter))

	router.Post("/hello", HelloEndpoint)
	router.Post("/hello2", HelloEndpoint)
	http.ListenAndServe(":8080", router)
```

