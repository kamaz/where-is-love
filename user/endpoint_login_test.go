package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	loginJSON := `{"email":"test@email.com", "password": "secret"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := loginEndpoint{repository: &mockUserRepository{}, tokenGenerator: &simpleTokenGenerator{}}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.JSONEq(
			`{"result": {"token": "eyJpZCI6MSwiZW1haWwiOiJ0ZXN0QGVtYWlsLmNvbSIsIm5hbWUiOiJuYW1lIiwiZ2VuZGVyIjoibWFsZSIsImFnZSI6MjIsImxhdGl0dWRlIjowLCJsb25naXR1ZGUiOjAsImNpdHkiOiIifQ=="}}`,
			rec.Body.String(),
		)
	}
}
