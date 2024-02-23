package match

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kamaz/where-is-love/user"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestDiscover(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(context.WithValue(req.Context(), user.UserKey, &user.UserToken{Id: 1}))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := DiscoverEndpoint{
		repository: &MockMatchRepository{},
	}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.JSONEq(
			`{"results": [{"id": 1, "name": "Mark", "gender": "male", "age":23}, {"id": 2, "name": "Joanna", "gender": "female", "age":25}]}`,
			rec.Body.String(),
		)
	}
}

func TestDiscoverWithGenderFilter(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?gender=male", nil)
	req = req.WithContext(context.WithValue(req.Context(), user.UserKey, &user.UserToken{Id: 1}))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := DiscoverEndpoint{
		repository: &MockMatchRepository{},
	}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.JSONEq(
			`{"results": [{"id": 1, "name": "Mark", "gender": "male", "age":23}]}`,
			rec.Body.String(),
		)
	}
}

func TestDiscoverWithAgeFilter(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?age=25", nil)
	req = req.WithContext(context.WithValue(req.Context(), user.UserKey, &user.UserToken{Id: 1}))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := DiscoverEndpoint{
		repository: &MockMatchRepository{},
	}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.JSONEq(
			`{"results": [{"id": 2, "name": "Joanna", "gender": "female", "age":25}]}`,
			rec.Body.String(),
		)
	}
}

func TestDiscoverWithAgeAndGenderFilters(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?age=25&gender=female", nil)
	req = req.WithContext(context.WithValue(req.Context(), user.UserKey, &user.UserToken{Id: 1}))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := DiscoverEndpoint{
		repository: &MockMatchRepository{},
	}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		assert.JSONEq(
			`{"results": [{"id": 2, "name": "Joanna", "gender": "female", "age":25}]}`,
			rec.Body.String(),
		)
	}
}
