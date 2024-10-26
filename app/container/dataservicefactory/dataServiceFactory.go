package dataservicefactory

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
)

// DataServiceFactoryBuilderMap is a map of string to DataServiceFbInterface
// This part should be generated
var dsFbMap = map[string]dataServiceFbInterface{
	config.TaskQueueSystemData: &taskQueueDataServiceFactoryWrapper{},
}

type DataServiceInterface interface{}

type dataServiceFbInterface interface {
	Build(container.Container, *config.DataConfig) (DataServiceInterface, error)
}

func GetDataServiceFb(key string) dataServiceFbInterface {
	return dsFbMap[key]
}
