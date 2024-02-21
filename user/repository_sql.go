package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLUserRepository struct {
	db *pgxpool.Pool
}

func CreateSQLUserRepository(pool *pgxpool.Pool) *SQLUserRepository {
	return &SQLUserRepository{db: pool}
}

func (u *SQLUserRepository) CreateUser(ctx context.Context) (*UserEntity, error) {
	now := time.Now().Unix()

	gender := "male"
	if rand.Intn(100)%2 == 0 {
		gender = "female"
	}

	user := &UserEntity{
		Email:    fmt.Sprintf("test-%d@example.com", now),
		Password: "secret password",
		Name:     "First Last",
		Gender:   gender,
		Age:      20,
	}

	var userId int

	err := u.db.QueryRow(
		context.Background(),
		"INSERT INTO app_user(email, password, name, gender, age) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Email,
		user.Password,
		user.Name,
		user.Gender,
		user.Age,
	).Scan(&userId)

	user.Id = userId
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (u *SQLUserRepository) GetUserByEmailAndPassword(
	ctx context.Context,
	criteria *EmailAndPasswordCriteria,
) (*UserEntity, error) {
	var user UserEntity
	err := u.db.QueryRow(
		context.Background(),
		"SELECT id, email, name, gender, age FROM app_user WHERE email = $1 AND password = $2",
		criteria.Email,
		criteria.Password,
	).Scan(&user.Id, &user.Email, &user.Name, &user.Gender, &user.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}
