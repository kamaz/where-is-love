package user

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

const UserKey = "user"

// ServerHeader middleware adds a `Server` header to the response.
// todo: introduce a common error handling
func AppAuthorization(tg TokenGenerator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			token := req.Header.Get(echo.HeaderAuthorization)
			user, err := tg.Validate(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]string{})
				return err
			}
			*req = *req.WithContext(context.WithValue(req.Context(), UserKey, user))
			return next(c)
		}
	}
}
