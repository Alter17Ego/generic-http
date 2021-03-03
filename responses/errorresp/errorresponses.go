package errorresp

import (
	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"github.com/Alter17Ego/generic-http/responses"
)

type CodedErrorResponse struct {
	*responses.StatusCodedResponse
	*codederrors.CodedError
}

func New(statusCode int, err *codederrors.CodedError, customRenderers ...responses.StatusRenderer) *CodedErrorResponse {
	return &CodedErrorResponse{StatusCodedResponse: responses.New(statusCode, customRenderers...), CodedError: err}
}

func Forbidden(err *codederrors.CodedError) *CodedErrorResponse {
	return New(403, err)
}

func Unauthorized(err *codederrors.CodedError) *CodedErrorResponse {
	return New(401, err)
}

func BadRequest(err *codederrors.CodedError) *CodedErrorResponse {
	return New(400, err)
}

func InternalServerError(err *codederrors.CodedError) *CodedErrorResponse {
	return New(500, err)
}
