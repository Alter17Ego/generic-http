package responses

import (
	"net/http"

	"github.com/go-chi/render"
)

type StatusRenderer func(r *http.Request, status int)

type StatusCodedResponse struct {
	StatusCode   int `json:"-"`
	renderStatus StatusRenderer
}

func (res StatusCodedResponse) Render(w http.ResponseWriter, r *http.Request) error {
	res.renderStatus(r, res.StatusCode)
	return nil
}

func New(statusCode int, statusRenderers ...StatusRenderer) *StatusCodedResponse {
	statusRenderer := render.Status
	if len(statusRenderers) > 0 {
		statusRenderer = statusRenderers[0]
	}
	return &StatusCodedResponse{StatusCode: statusCode, renderStatus: statusRenderer}
}
