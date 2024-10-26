package containerhelper

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
	"github.com/Systenix/taskmanager/task_queue_system/usecase"
)

func GetTaskQueueUseCase(c container.Container) (usecase.TaskManagerUseCase, error) {

	key := config.TaskQueueSystemUseCase
	useCase, err := c.BuildUseCase(key)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, err
	}
	return useCase.(usecase.TaskManagerUseCase), nil
}
