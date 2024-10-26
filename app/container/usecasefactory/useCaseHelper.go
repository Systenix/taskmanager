package usecasefactory

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
	"github.com/Systenix/taskmanager/task_queue_system/app/container/dataservicefactory"
	"github.com/Systenix/taskmanager/task_queue_system/infrastructure/dataservice"
)

func buildTaskQueueDataService(c container.Container, dc *config.DataConfig) (dataservice.TaskQueueDataService, error) {
	dataServiceInterface, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, err
	}
	udi := dataServiceInterface.(dataservice.TaskQueueDataService)
	return udi, nil
}
