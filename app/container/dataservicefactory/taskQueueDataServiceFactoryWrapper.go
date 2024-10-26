package dataservicefactory

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
	"github.com/Systenix/taskmanager/task_queue_system/app/container/dataservicefactory/taskmanagerdataservicefactory"
)

type taskQueueDataServiceFactoryWrapper struct{}

func (usdsfw *taskQueueDataServiceFactoryWrapper) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	udsi, err := taskmanagerdataservicefactory.GetTaskQueueDataServiceFb(config.NoSqlDB).Build(c, dataConfig)
	if err != nil {
		return nil, err
	}
	return udsi, nil
}
