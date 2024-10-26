package taskmanager

import (
	"context"
	"errors"
	"time"

	"github.com/Systenix/taskmanager/task_queue_system/domain/model"
	"github.com/Systenix/taskmanager/task_queue_system/infrastructure/dataservice"
	"github.com/Systenix/taskmanager/task_queue_system/usecase"
	"github.com/google/uuid"
)

// taskManagerUseCase implements TaskManagerUseCase.
type taskManagerUseCase struct {
	TaskDataService dataservice.TaskQueueDataService
}

func NewTaskManagerUseCase(taskDataService dataservice.TaskQueueDataService) usecase.TaskManagerUseCase {
	return &taskManagerUseCase{
		TaskDataService: taskDataService,
	}
}

func (uc *taskManagerUseCase) SubmitTask(ctx context.Context, taskType string, payload string) (string, error) {
	// Create a new task
	task := &model.Task{
		ID:          uuid.New().String(),
		Type:        taskType,
		Payload:     payload,
		Status:      model.StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Attempts:    0,
		MaxAttempts: 5, // Could be configurable
	}

	// Save the task status
	err := uc.TaskDataService.SaveTask(ctx, task)
	if err != nil {
		return "", err
	}

	// Enqueue the task in the message broker
	err = uc.TaskDataService.EnqueueTask(ctx, task)
	if err != nil {
		return "", err
	}

	return task.ID, nil
}

func (uc *taskManagerUseCase) GetTaskStatus(ctx context.Context, taskID string) (*model.Task, error) {
	return uc.TaskDataService.GetTask(ctx, taskID)
}

func (uc *taskManagerUseCase) CancelTask(ctx context.Context, taskID string) error {
	task, err := uc.TaskDataService.GetTask(ctx, taskID)
	if err != nil {
		return err
	}
	if task.Status != model.StatusPending && task.Status != model.StatusRunning {
		return errors.New("task cannot be canceled")
	}

	task.Status = model.StatusCanceled
	task.UpdatedAt = time.Now()

	// Update the task status
	return uc.TaskDataService.UpdateTask(ctx, task)
}
