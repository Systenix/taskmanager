package dataservice

import (
	"context"
	"github.com/Systenix/taskmanager/task_queue_system/domain/model"
)

// TaskQueueDataService defines the interface for task data service.
type TaskQueueDataService interface {
	SaveTask(ctx context.Context, task *model.Task) error
	GetTask(ctx context.Context, taskID string) (*model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) error
	EnqueueTask(ctx context.Context, task *model.Task) error
}
