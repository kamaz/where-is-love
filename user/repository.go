package user

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound                = errors.New("user not found")
	_               UserRepository = (*sqlUserRepository)(nil)
	_               UserRepository = (*mockUserRepository)(nil)
)

// UserEntity represents a user entity
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

// EmailAndPasswordCriteria represents criteria for getting user by email and password
type EmailAndPasswordCriteria struct {
	Email    string
	Password string
}

// UserRepository is an interface for user repository
type UserRepository interface {
	CreateUser(ctx context.Context) (*UserEntity, error)
	GetUserByEmailAndPassword(
		ctx context.Context,
		criteria *EmailAndPasswordCriteria,
	) (*UserEntity, error)
}
