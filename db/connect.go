package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func ConnectDb() (*pgx.Conn, error) {
	connect, err := pgx.Connect(context.Background(), "postgres://alanzhumalin:@localhost:5432/alanzhumalin")

	if err != nil {
		msg := "Error occured while connection with db:" + err.Error()
		return nil, errors.New(msg)
	}

	return connect, nil
}
