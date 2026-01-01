package internal

import (
	"context"
	"database/sql"
	"encoding/json"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) GetAll(ctx context.Context) ([]Task, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, job_name, start_at, stop_at, is_close, priority, tags FROM tasks`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		var tags string

		if err := rows.Scan(
			&t.ID,
			&t.JobName,
			&t.StartAt,
			&t.StopAt,
			&t.IsClose,
			&t.Priority,
			&tags,
		); err != nil {
			return nil, err
		}

		_ = json.Unmarshal([]byte(tags), &t.Tags)
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *SQLiteRepository) Create(ctx context.Context, task *Task) error {
	tags, _ := json.Marshal(task.Tags)

	res, err := r.db.ExecContext(ctx,
		`INSERT INTO tasks (job_name, start_at, stop_at, is_close, priority, tags)
		 VALUES (?, ?, ?, ?, ?, ?)`,
		task.JobName,
		task.StartAt,
		task.StopAt,
		task.IsClose,
		task.Priority,
		string(tags),
	)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	task.ID = id
	return nil
}
