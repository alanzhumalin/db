package sqlquery

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteTable(ctx context.Context, con *pgx.Conn, id int) error {
	sqlquery := `
	DELETE FROM tasks WHERE id = $1
	`

	_, err := con.Exec(ctx, sqlquery, id)

	return err
}
