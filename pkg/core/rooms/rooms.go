package rooms

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Service struct {
	pool *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

type Rooms struct {
	Id     int64 `json:"id"`
	Status bool `json:"status"`
	TimeInFour int `json:"time_in_four"`
	TimeInMinutes int `json:"time_in_minutes"`
	TimeOutFour int `json:"time_out_four"`
	TimeOutMinutes int `json:"time_out_minutes"`
	FileName string `json:"file_name"`
}

func (s *Service) AllRooms (context context.Context) ([]Rooms , error)  {
	list := make([]Rooms, 0)
	rows, err := s.pool.Query(context, `SELECT id, status, timeinhour, timeinminutes, timeouthour, timeoutminutes, filename FROM mitings;`)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	log.Print(rows)
	for rows.Next() {
		item := Rooms{}
		err := rows.Scan(&item.Id, &item.Status, &item.TimeInFour, &item.TimeInMinutes, &item.TimeOutFour, &item.TimeOutMinutes, &item.FileName)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	log.Print(list)
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}