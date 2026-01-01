package internal

import "context"

type TaskRepository interface {
	GetAll(ctx context.Context) ([]Task, error)
	Create(ctx context.Context, task *Task) error
}
