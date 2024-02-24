package match

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kamaz/where-is-love/user"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSwipe_WithoutMatch(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	matchJSON := `{"userID":2, "preference": "YES"}`
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(matchJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req = req.WithContext(context.WithValue(req.Context(), user.UserKey, &user.UserToken{Id: 1}))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := swipeEndpoint{repository: CreateMockMatchRepository()}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		// ideally we should have here 3 states
		// 1. matched
		// 2. no
		// 3. maybe (especially when second person did not responded)
		assert.JSONEq(
			`{"result": {"matched": false}}`,
			rec.Body.String(),
		)
	}
}

func TestSwipe_WithMatch(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	matchJSON := `{"userID":3, "preference": "YES"}`
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(matchJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req = req.WithContext(context.WithValue(req.Context(), user.UserKey, &user.UserToken{Id: 1}))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := swipeEndpoint{repository: CreateMockMatchRepository()}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		// ideally we should have here 3 states
		// 1. matched
		// 2. no
		// 3. maybe (especially when second person did not responded)
		assert.JSONEq(
			`{"result": {"matched": true, "matchID": 3}}`,
			rec.Body.String(),
		)
	}
}
