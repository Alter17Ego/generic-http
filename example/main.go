package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"github.com/Alter17Ego/generic-http/limiter"
	"github.com/Alter17Ego/generic-http/responses"
	"github.com/Alter17Ego/generic-http/responses/errorresp"
)

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	*responses.StatusCodedResponse
	Message string `json:"message"`
}

func newHelloResponse(msg string) *HelloResponse {
	return &HelloResponse{Message: msg, StatusCodedResponse: responses.New(200)}
}

func HelloEndpoint(w http.ResponseWriter, r *http.Request) {
	var requestBody HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		wrappedErr := codederrors.Wrap(err, "-1", "something went wrong while decoding request")
		resp := errorresp.InternalServerError(wrappedErr)
		render.Render(w, r, resp)
		fmt.Println(err.Error())
		return
	}
	render.Render(w, r, newHelloResponse(fmt.Sprint("Hello", " ", requestBody.Name)))
	return
}

func main() {
	router := chi.NewRouter()

	router.Use(limiter.NewChiMiddleware(limiter.
		New(&limiter.
			Setting{
			ExpirationIn:         time.Hour,
			MaxRequestsPerSecond: 1,
		},
			limiter.UseProductionConfig())))

	router.Post("/hello", HelloEndpoint)
	router.Post("/hello2", HelloEndpoint)
	http.ListenAndServe(":8080", router)
}
