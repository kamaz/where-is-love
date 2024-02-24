package user

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound                = errors.New("user not found")
	_               UserRepository = (*SQLUserRepository)(nil)
	_               UserRepository = (*MockUserRepository)(nil)
)

type UserEntity struct {
	Id        uint
	Email     string
	Password  string
	Name      string
	Gender    string
	Age       uint
	Longitude float64
	Latitude  float64
	City      string
}

type EmailAndPasswordCriteria struct {
	Email    string
	Password string
}

type UserRepository interface {
	CreateUser(ctx context.Context) (*UserEntity, error)
	GetUserByEmailAndPassword(
		ctx context.Context,
		criteria *EmailAndPasswordCriteria,
	) (*UserEntity, error)
}
