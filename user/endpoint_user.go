package user

import (
	"fmt"
	"net/http"

	"github.com/kamaz/where-is-love/types"
	"github.com/labstack/echo"
)

type UserResponse struct {
	Id        uint    `json:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Name      string  `json:"name"`
	Gender    string  `json:"gender"`
	Age       uint    `json:"age"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	City      string  `json:"city"`
}

type UserResult struct {
	Result *UserResponse `json:"result"`
}

// /user/create
// Generate and store user
//
//	{
//	    "result": {
//	        "id": <integer>,
//	        "email": <string>,
//	        "password": <string>,
//	        "name": <string>,
//	        "gender": <string>,
//	        "age": <integer>
//	    }
//	}
type userEndpoint struct {
	repository  UserRepository
	middlewares []echo.MiddlewareFunc
}

func CreateUserEndpoint(repo UserRepository) types.Endpoint {
	return &userEndpoint{
		repository: repo,
	}
}

func (u *userEndpoint) Process(e echo.Context) error {
	user, err := u.repository.CreateUser(e.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to create user %w", err)
	}
	userResponse := UserResponse(*user)
	result := UserResult{Result: &userResponse}

	e.JSON(http.StatusCreated, result)
	return nil
}

func (u *userEndpoint) Method() string {
	return "POST"
}

func (u *userEndpoint) Path() string {
	return "/user/create"
}

func (u *userEndpoint) Middlewares() []echo.MiddlewareFunc {
	return u.middlewares
}
