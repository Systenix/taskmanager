package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Systenix/taskmanager/task_queue_system/domain/model"
	"github.com/Systenix/taskmanager/task_queue_system/infrastructure/dataservice"
	"github.com/redis/go-redis/v9"
)

type taskDataService struct {
	redisClient *redis.Client
}

func NewTaskDataService(redisClient *redis.Client) dataservice.TaskQueueDataService {
	return &taskDataService{redisClient: redisClient}
}

func (ds *taskDataService) SaveTask(ctx context.Context, task *model.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	// Save task status in Redis (could also be a database)
	return ds.redisClient.HSet(ctx, "tasks", task.ID, data).Err()
}

func (ds *taskDataService) GetTask(ctx context.Context, taskID string) (*model.Task, error) {
	data, err := ds.redisClient.HGet(ctx, "tasks", taskID).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	var task model.Task
	err = json.Unmarshal([]byte(data), &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (ds *taskDataService) UpdateTask(ctx context.Context, task *model.Task) error {
	return ds.SaveTask(ctx, task)
}

func (ds *taskDataService) EnqueueTask(ctx context.Context, task *model.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	// Push task to Redis list (or use a message queue)
	return ds.redisClient.LPush(ctx, "task_queue", data).Err()
}
