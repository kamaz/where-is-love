package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := UserEndpoint{repository: &MockUserRepository{}}

	if assert.NoError(h.Process(c)) {
		assert.Equal(http.StatusCreated, rec.Code)
		assert.JSONEq(
			`{"result": {"age":22, "email": "test@email.com", "gender": "male", "id": 1, "name": "name", "password": "password" }}`,
			rec.Body.String(),
		)
	}
}
