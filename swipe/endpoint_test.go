package swipe

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSwipe(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := SwipeEndpoint{}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusOK, rec.Code)
		// ideally we should have here 3 states
		// 1. matched
		// 2. no
		// 3. maybe (especially when second person did not responded)
		assert.JSONEq(
			`{"result": {}}`,
			rec.Body.String(),
		)
	}
}
