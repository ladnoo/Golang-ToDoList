package database

import (
	"ToDoList/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
)

// code repetition... needs to be fixed...

func InsertTask(ctx context.Context, conn *pgx.Conn, task models.TaskModel) error {
	sqlQuery := `
		INSERT INTO tasks (title, description, completed, created_at)
		VALUES ($1, $2, $3, $4);
	`

	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CreatedAt)
	return err
}

func UpdateTask(ctx context.Context, conn *pgx.Conn, task models.TaskModel) error {
	sqlQuery := `
		UPDATE tasks
		SET title=$1, description=$2, completed=$3, completed_at=$4
		Where id = $5;
	`

	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CompletedAt)
	return err
}

func SelectUncompletedTasks(ctx context.Context, conn *pgx.Conn) ([]models.TaskModel, error) {
	sqlQuery := `
		SELECT id, title, description, completed, created_at
		WHERE completed = FALSE
		ORDER BY created_at DESC
		`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks = make([]models.TaskModel, 0)
	for rows.Next() {
		var task models.TaskModel

		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func DeleteTasks(ctx context.Context, conn *pgx.Conn, tasksId []models.TaskModel) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = ANY($1);
`

	_, err := conn.Exec(ctx, sqlQuery, tasksId)
	return err

}
