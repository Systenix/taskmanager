package taskmanagerdataservicefactory

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
	"github.com/Systenix/taskmanager/task_queue_system/infrastructure/dataservice"
)

var taskQueuedsFbMap = map[string]taskQueueDataServiceFbInterface{
	config.NoSqlDB: &redisDataServiceFactory{},
}

type taskQueueDataServiceFbInterface interface {
	Build(container.Container, *config.DataConfig) (dataservice.TaskQueueDataService, error)
}

func GetTaskQueueDataServiceFb(key string) taskQueueDataServiceFbInterface {
	return taskQueuedsFbMap[key]
}
