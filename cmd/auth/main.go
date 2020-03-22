package main

import (
	"auth-service/cmd/auth/app"
	"auth-service/pkg/core/add"
	"auth-service/pkg/core/token"
	"auth-service/pkg/core/user"
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shuhrat-shokirov/jwt/pkg/cmd"
	"github.com/shuhrat-shokirov/mux/pkg/mux"
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
	pool, err := pgxpool.Connect(context.Background(), string(dsn))
	if err != nil {
		panic(fmt.Errorf("can't create pool: %w", err))
	}
	service := token.NewService(secret, pool)
	exactMux := mux.NewExactMux()
	newService := add.NewService(pool)
	userSvc := user.NewService()
	server := app.NewServer(exactMux, pool, secret, service, userSvc, newService)
	server.Start()
	panic(http.ListenAndServe(addr, server))
}