package user

import (
	"fmt"
	"net/http"

	"github.com/kamaz/where-is-love/types"
	"github.com/labstack/echo"
)

// /login
// {
//     "results": {
//          "token": <string>
//      }
// }

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginResult struct {
	Result *LoginResponse `json:"result"`
}

type loginEndpoint struct {
	repository     UserRepository
	tokenGenerator TokenGenerator
	middlewares    []echo.MiddlewareFunc
}

func CreateLoginEndpoint(repo UserRepository, tokenGenerator TokenGenerator) types.Endpoint {
	return &loginEndpoint{
		repository:     repo,
		tokenGenerator: tokenGenerator,
	}
}

func (u *loginEndpoint) Process(e echo.Context) error {
	var request LoginRequest
	if err := e.Bind(&request); err != nil {
		return fmt.Errorf("failed to bind payload %w", err)
	}

	userEntity, err := u.repository.GetUserByEmailAndPassword(
		e.Request().Context(),
		&EmailAndPasswordCriteria{
			Email:    request.Email,
			Password: request.Password,
		},
	)
	if err != nil {
		if err == ErrUserNotFound {
			e.JSON(http.StatusBadRequest, map[string]string{})
			return nil
		}
	}

	token, err := u.tokenGenerator.Generate(toToken(userEntity))
	if err != nil {
		return fmt.Errorf("failed to generate token %w", err)
	}

	result := LoginResult{Result: &LoginResponse{
		Token: token,
	}}
	e.JSON(http.StatusOK, result)
	return nil
}

func (u *loginEndpoint) Method() string {
	return "POST"
}

func (u *loginEndpoint) Path() string {
	return "/login"
}

func (u *loginEndpoint) Middlewares() []echo.MiddlewareFunc {
	return u.middlewares
}
