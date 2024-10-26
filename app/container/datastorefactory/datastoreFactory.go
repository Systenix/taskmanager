package datastorefactory

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
)

var dsFbMap = map[string]dsFbInterface{
	config.Redis: &redisFactory{},
}

type DataStoreInterface interface{}

type dsFbInterface interface {
	Build(container.Container, *config.DataStoreConfig) (DataStoreInterface, error)
}

func GetDataStoreFb(key string) dsFbInterface {
	return dsFbMap[key]
}
