package main

import (
	"log"
	"os"

	"github.com/Systenix/taskmanager/task_queue_system/app"
	"github.com/Systenix/taskmanager/task_queue_system/app/container/servicecontainer"
	"github.com/Systenix/taskmanager/task_queue_system/app/logger"
	interfacehttp "github.com/Systenix/taskmanager/task_queue_system/interfaces/http"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func setupRouter(sc *servicecontainer.ServiceContainer) *gin.Engine {

	logger.Log.Info("starting server...")

	router := gin.Default()
	err := router.SetTrustedProxies(sc.AppConfig.ServiceListener.TrustedProxies)
	if err != nil {
		return nil
	}

	taskService := &interfacehttp.TaskManagerService{Container: sc}

	// Submit a new task
	router.POST("/api/v1/tasks", taskService.SubmitTask)

	// Get task status
	router.GET("/api/v1/tasks/:task_id", taskService.GetTaskStatus)

	// Cancel a task
	router.DELETE("/api/v1/tasks/:task_id", taskService.CancelTask)

	return router
}

func buildContainer(filename string) (*servicecontainer.ServiceContainer, error) {

	container, err := app.InitApp("task-queue-system", filename)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	sc := container.(*servicecontainer.ServiceContainer)

	return sc, nil
}

func main() {

	filename := os.Getenv("SERVICE_CONFIG")
	container, err := buildContainer(filename)
	if err != nil {
		log.Println("error while building container:", err.Error())
		panic(err)
	}

	router := setupRouter(container)
	if err = router.Run(container.AppConfig.ServiceListener.Address); err != nil {
		panic(err)
	}
}
