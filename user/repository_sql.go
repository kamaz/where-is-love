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
	now := time.Now().UnixNano()

	gender := "male"
	if rand.Intn(100)%2 == 0 {
		gender = "female"
	}
	age := 20 + rand.Intn(30)

	cities := [][]any{
		{"London", 51.5072, -0.1275},
		{"Birmingham", 52.4800, -1.9025},
		{"Manchester", 53.4794, -2.2453},
		{"Liverpool", 53.4075, -2.9919},
		{"Portsmouth", 50.8058, -1.0872},
		{"Southampton", 50.9025, -1.4042},
		{"Nottingham", 52.9533, -1.1500},
		{"Bristol", 51.4536, -2.5975},
		{"Leicester", 52.6344, -1.1319},
		{"Coventry", 52.4081, -1.5106},
		{"Belfast", 54.5964, -5.9300},
		{"Stockport", 53.4083, -2.1494},
		{"Bradford", 53.8000, -1.7500},
		{"Plymouth", 50.3714, -4.1422},
		{"Derby", 52.9217, -1.4767},
		{"Westminster", 51.4947, -0.1353},
		{"Wolverhampton", 52.5833, -2.1333},
		{"Norwich", 52.6286, 1.2928},
		{"Luton", 51.8783, -0.4147},
	}
	city := cities[rand.Intn(len(cities))]
	latitude := city[1].(float64)
	longitude := city[2].(float64)

	user := &UserEntity{
		Email:     fmt.Sprintf("test-%d@example.com", now),
		Password:  "secret password",
		Name:      fmt.Sprintf("First-%d Last", now),
		Gender:    gender,
		Age:       uint(age),
		Latitude:  latitude,
		Longitude: longitude,
		City:      city[0].(string),
	}

	var userId uint

	err := u.db.QueryRow(
		ctx,
		"INSERT INTO app_user(email, password, name, gender, age, latitude, longitude, city) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		user.Email,
		user.Password,
		user.Name,
		user.Gender,
		user.Age,
		user.Latitude,
		user.Longitude,
		user.City,
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
		ctx,
		"SELECT id, email, name, gender, age, latitude, longitude, city FROM app_user WHERE email = $1 AND password = $2",
		criteria.Email,
		criteria.Password,
	).Scan(&user.Id, &user.Email, &user.Name, &user.Gender, &user.Age, &user.Latitude, &user.Longitude, &user.City)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}
