package responses_test

import (
	"net/http"
	"testing"

	"github.com/Alter17Ego/generic-http/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockResponse struct {
	mock.Mock
}

func (res *MockResponse) CustomRender(r *http.Request, status int) {
	res.Called(r, status)
}

type ResponseTestSuite struct {
	suite.Suite
}

func (s ResponseTestSuite) TestNewShouldCreateResponseWithResponseCode() {
	statusCode := 200
	response := responses.New(statusCode)
	assert.Equal(s.T(), statusCode, response.StatusCode)
}

func (s ResponseTestSuite) TestRenderShouldCallCustomRenderer() {
	mockObj := new(MockResponse)
	statusCode := 200
	stub := mockObj.On("CustomRender", mock.Anything, 200)
	response := responses.New(statusCode, mockObj.CustomRender)
	response.Render(nil, nil)
	stub.Times(1)
}

func TestResponses(t *testing.T) {
	suite.Run(t, new(ResponseTestSuite))
}
