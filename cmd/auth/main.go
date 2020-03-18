package main

import (
	"auth-service/cmd/auth/app"
	"auth-service/pkg/core/token"
	"auth-service/pkg/core/user"
	"auth-service/pkg/di"
	"auth-service/pkg/jwt"
	"auth-service/pkg/mux"
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"net/http"
)

// flag - max priority, env - lower priority

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9999", "Server port")
	dsn  = flag.String("dsn", "postgres://user:pass@localhost:5401/auth", "Postgres DSN")
)

type DSN string

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	// get secret from env
	// обычно для приложения создаётся отдельный пользователь ОС с урезанными правами
	secret := jwt.Secret("secret")
	start(addr, *dsn, secret)
}

func start(addr string, dsn string, secret jwt.Secret) {
	// DI - Martin Fowler
	container := di.NewContainer()

	container.Provide(
		app.NewServer,
		mux.NewExactMux,
		func() jwt.Secret { return secret },
		func() DSN { return DSN(dsn) },
		func(dsn DSN) *pgxpool.Pool {
			pool, err := pgxpool.Connect(context.Background(), string(dsn))
			if err != nil {
				panic(fmt.Errorf("can't create pool: %w", err))
			}
			return pool
		},
		token.NewService,
		user.NewService,
	)

	container.Start()
	// IoC - inversion of control (программа определяет, куда вы можете встроиться)
	// StartListener, StopListener
	// см. Errors.As
	var appServer *app.Server
	container.Component(&appServer)

	panic(http.ListenAndServe(addr, appServer))
}

