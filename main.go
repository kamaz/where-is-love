package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("Running Happily!")

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	_, err = dbpool.Exec(
		context.Background(),
		"INSERT INTO app_user(email, password, name, gender, age) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		"test@example.com",
		"secret",
		"Test User",
		"male",
		25,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	time.Sleep(10 * time.Minute)
}
