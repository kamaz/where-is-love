package types

import "github.com/labstack/echo"

// Endpoint is an interface for all server endpoints
type Endpoint interface {
	Process(echo.Context) error
	Method() string
	Path() string
	Middlewares() []echo.MiddlewareFunc
}

// todo: Return an appropriate error if login fails.
// todo: Ensure that all other endpoints are appropriately authenticated.
