package user

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

var _ TokenGenerator = (*SimpleTokenGenerator)(nil)

type UserToken struct {
	Id        uint    `json:"id"`
	Email     string  `json:"email"`
	Name      string  `json:"name"`
	Gender    string  `json:"gender"`
	Age       uint    `json:"age"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
}

func toToken(entity *UserEntity) *UserToken {
	return &UserToken{
		Id:        entity.Id,
		Email:     entity.Email,
		Name:      entity.Name,
		Gender:    entity.Gender,
		Age:       entity.Age,
		Latitude:  entity.Latitude,
		Longitude: entity.Longitude,
		City:      entity.City,
	}
}

// TokenGenerator is an interface for token generation and validation of token.
type TokenGenerator interface {
	Generate(user *UserToken) (string, error)
	Validate(token string) (*UserToken, error)
}

// SimpleTokenGenerator is a simple token generator
// which uses a json to marshal and unmarshal the user entity
// and then convert the generated string to base64.
type SimpleTokenGenerator struct{}

func (t *SimpleTokenGenerator) Generate(user *UserToken) (string, error) {
	userContent, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user: %w", err)
	}

	return base64.StdEncoding.EncodeToString(userContent), nil
}

func (t *SimpleTokenGenerator) Validate(token string) (*UserToken, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token: %w", err)
	}
	var user UserToken
	if err := json.Unmarshal(decoded, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	if user.Id == 0 || user.Age == 0 || user.Email == "" || user.Name == "" || user.Gender == "" {
		return nil, nil
	}

	return &user, nil
}
