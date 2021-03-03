package errorresp_test

import (
	"testing"

	"github.com/Alter17Ego/generic-app/errors/codederrors"
	"github.com/Alter17Ego/generic-http/responses/errorresp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ErrorResponseTestSuite struct {
	suite.Suite
}

func (s ErrorResponseTestSuite) TestNewShouldCreateErrorResponseWithCodedError() {
	err := codederrors.New("err.token.exp", "your token are expired")
	errResponse := errorresp.New(401, err)
	assert.Equal(s.T(), err, errResponse.CodedError)
}

func (s ErrorResponseTestSuite) TestInternalServerErrorShouldHaveStatus500() {
	err := codederrors.New("err.token.parse", "unable to parse token")
	errResponse := errorresp.InternalServerError(err)
	assert.Equal(s.T(), 500, errResponse.StatusCode)
}

func (s ErrorResponseTestSuite) TestForbiddenErrorShouldHaveStatus403() {
	err := codederrors.New("err.forbidden.access", "forbiden.access.manager")
	errResponse := errorresp.Forbidden(err)
	assert.Equal(s.T(), 403, errResponse.StatusCode)
}

func (s ErrorResponseTestSuite) TestUnauthorizedShouldHaveStatus401() {
	err := codederrors.New("err.token.expr", "token was expired")
	errResponse := errorresp.Unauthorized(err)
	assert.Equal(s.T(), 401, errResponse.StatusCode)
}

func (s ErrorResponseTestSuite) TestBadRequestShouldHaveStatus400() {
	err := codederrors.New("balance.must.gt", "balance must greater then 10.000")
	errResponse := errorresp.BadRequest(err)
	assert.Equal(s.T(), 400, errResponse.StatusCode)
}

func TestErrResp(t *testing.T) {
	suite.Run(t, new(ErrorResponseTestSuite))
}
