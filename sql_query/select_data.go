package sqlquery

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, con *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at 
	FROM tasks
	`

	data, err := con.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	tasks := make([]TaskModel, 0)

	for data.Next() {
		var task TaskModel

		data.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)

		tasks = append(tasks, task)

	}

	return tasks, nil

}
