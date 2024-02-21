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
	Id       int
	Email    string
	Password string
	Name     string
	Gender   string
	Age      int
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
