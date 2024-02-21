package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type UserResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
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
type UserEndpoint struct {
	repository  UserRepository
	middlewares []echo.MiddlewareFunc
}

func CreateUserEndpoint(repo UserRepository) *UserEndpoint {
	return &UserEndpoint{
		repository: repo,
	}
}

func (u *UserEndpoint) Process(e echo.Context) error {
	user, err := u.repository.CreateUser(e.Request().Context())
	// check if there is an log and mask error
	if err != nil {
		return fmt.Errorf("failed to create user")
	}
	userResponse := UserResponse(*user)
	result := UserResult{Result: &userResponse}

	e.JSON(http.StatusCreated, result)
	return nil
}

func (u *UserEndpoint) Method() string {
	return "POST"
}

func (u *UserEndpoint) Path() string {
	return "/user/create"
}

func (u *UserEndpoint) Middlewares() []echo.MiddlewareFunc {
	return u.middlewares
}
