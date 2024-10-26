package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Systenix/taskmanager/task_worker/model"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	for {
		result, err := redisClient.BRPop(ctx, 0*time.Second, "task_queue").Result()
		if err != nil {
			log.Println("Error fetching task:", err)
			continue
		}

		// result[0] is the list name, result[1] is the value
		var task model.Task
		err = json.Unmarshal([]byte(result[1]), &task)
		if err != nil {
			log.Println("Error unmarshaling task:", err)
			continue
		}

		// Process the task
		processTask(ctx, &task, redisClient)
	}
}

func processTask(ctx context.Context, task *model.Task, redisClient *redis.Client) {
	// Simulate task processing
	log.Printf("Processing task %s of type %s", task.ID, task.Type)
	time.Sleep(2 * time.Second)

	// Update task status
	task.Status = model.StatusCompleted
	task.UpdatedAt = time.Now()

	data, err := json.Marshal(task)
	if err != nil {
		log.Println("Error marshaling task:", err)
		return
	}

	// Update task status in Redis
	err = redisClient.HSet(ctx, "tasks", task.ID, data).Err()
	if err != nil {
		log.Println("Error updating task status:", err)
	}
}
