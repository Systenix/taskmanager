package usecase

import (
	"context"
	"github.com/Systenix/taskmanager/task_queue_system/domain/model"
)

// TaskManagerUseCase defines the interface for task management.
type TaskManagerUseCase interface {
	SubmitTask(ctx context.Context, taskType string, payload string) (string, error)
	GetTaskStatus(ctx context.Context, taskID string) (*model.Task, error)
	CancelTask(ctx context.Context, taskID string) error
}
