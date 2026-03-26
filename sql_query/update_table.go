package sqlquery

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, con *pgx.Conn, id int) error {
	sqlQuery := `
	UPDATE tasks 
	SET completed = FALSE
	where id = $1
	`
	_, err := con.Exec(ctx, sqlQuery, id)

	return err
}
