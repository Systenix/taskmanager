package usecasefactory

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
)

var UseCaseFactoryBuilderMap = map[string]UseCaseFbInterface{
	config.TaskQueueSystemUseCase: &taskQueueFactory{},
}

type UseCaseInterface interface{}

type UseCaseFbInterface interface {
	Build(c container.Container, appConfig *config.AppConfig) (UseCaseInterface, error)
}

func GetUseCaseFb(key string) UseCaseFbInterface {
	return UseCaseFactoryBuilderMap[key]
}
