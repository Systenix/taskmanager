package taskmanagerdataservicefactory

import (
	"fmt"

	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
	"github.com/Systenix/taskmanager/task_queue_system/app/container/datastorefactory"
	"github.com/Systenix/taskmanager/task_queue_system/infrastructure/dataservice"
	"github.com/Systenix/taskmanager/task_queue_system/infrastructure/dataservice/redis"
	redisv9 "github.com/redis/go-redis/v9"
)

type redisDataServiceFactory struct{}

func (sdsf *redisDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (dataservice.TaskQueueDataService, error) {
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, err
	}

	// Assert that dsi is a *redis.Client
	connPool, ok := dsi.(*redisv9.Client)
	if !ok {
		return nil, fmt.Errorf("expected *redis.Client, got %T", dsi)
	}

	// Create the TaskQueueDataService using the connPool
	ds := redis.NewTaskDataService(connPool)
	return ds, nil
}
