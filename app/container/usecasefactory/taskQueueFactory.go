package usecasefactory

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
	"github.com/Systenix/taskmanager/task_queue_system/usecase/taskmanager"
)

type taskQueueFactory struct {
}

func (usf *taskQueueFactory) Build(c container.Container, appConfig *config.AppConfig) (UseCaseInterface, error) {
	uc := appConfig.UseCaseConfig.TaskQueue

	taskQueueDS, err := buildTaskQueueDataService(c, &uc.TaskQueueDataConfig)
	if err != nil {
		return nil, err
	}
	taskQueueUseCase := taskmanager.NewTaskManagerUseCase(taskQueueDS)

	return taskQueueUseCase, nil
}
