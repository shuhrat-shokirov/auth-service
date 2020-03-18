package token

import (
	"auth-service/pkg/jwt"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	secret jwt.Secret
	pool *pgxpool.Pool
}

func NewService(secret jwt.Secret, pool *pgxpool.Pool) *Service {
	return &Service{secret: secret, pool: pool}
}


type Payload struct {
	Id    int64    `json:"id"`
	Exp   int64    `json:"exp"`
	Roles []string `json:"roles"`
}

type RequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Token string `json:"token"`
}

var ErrInvalidLogin = errors.New("invalid password")
var ErrInvalidPassword = errors.New("invalid password")

func (s *Service) Generate(context context.Context, request *RequestDTO) (response ResponseDTO, err error) {
	var pass string
	err = s.pool.QueryRow(context, `SELECT password FROM users WHERE login = $1;
`, request.Username).Scan(&pass)
	if err != nil {
		err = ErrInvalidLogin
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	//hashInDb, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(request.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		err = ErrInvalidPassword
		return
	}

	response.Token, err = jwt.Encode(Payload{
		Id:    1,
		Exp:   time.Now().Add(time.Hour).Unix(),
		Roles: []string{"ROLE_USER"},
	}, s.secret)
	return
}
