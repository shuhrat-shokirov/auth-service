package add

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

type NewUser struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var ErrAddUser = errors.New("login has haven't")

func (s * Service) AddNewUser(context context.Context, request NewUser) ( err error) {
	log.Print(request)
	exec, err := s.pool.Exec(context, `INSERT INTO users (login, password) VALUES ($1, $2);`, request.Login, request.Password)
	if err != nil {
		log.Print(err)
		err = ErrAddUser
		return
	}
	log.Print(exec)
	return
}
