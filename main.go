package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kamaz/where-is-love/discover"
	"github.com/kamaz/where-is-love/server"
	"github.com/kamaz/where-is-love/swipe"
	"github.com/kamaz/where-is-love/user"
	"github.com/rs/zerolog/log"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		os.Exit(1)
	}
	defer dbpool.Close()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	stopper := make(chan struct{})
	go func() {
		<-done
		close(stopper)
	}()

	// todo: configure enviornment variables
	server := server.CreateServer(5000,
		user.CreateUserEndpoint(user.CreateSQLUserRepository(dbpool)),
		user.CreateLoginEndpoint(
			user.CreateSQLUserRepository(dbpool),
			&user.SimpleTokenGenerator{},
		),
		&discover.DiscoverEndpoint{},
		&swipe.SwipeEndpoint{},
	)
	server.Run()
	<-stopper
	server.Stop(context.Background())
}
