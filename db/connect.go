package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectDb(ctx context.Context) (*pgx.Conn, error) {

	return pgx.Connect(ctx, "postgres://alanzhumalin:@localhost:5432/alanzhumalin")
}
