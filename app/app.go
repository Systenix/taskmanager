package app

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
	"github.com/Systenix/taskmanager/task_queue_system/app/container/servicecontainer"
	"github.com/Systenix/taskmanager/task_queue_system/app/logger"
	"go.uber.org/zap"
)

func InitApp(prefix string, filename string) (container.Container, error) {
	config.LoadConfig(prefix, filename)
	// config.ValidateConfig(filename)
	appConfig := config.GetConfig[config.AppConfig]()
	if err := initLogger(); err != nil {
		return nil, err
	}
	return initContainer(appConfig)
}

func initLogger() error {
	log, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	logger.SetLogger(log)
	return nil
}

func initContainer(config *config.AppConfig) (container.Container, error) {
	factoryMap := make(map[string]interface{})
	c := servicecontainer.ServiceContainer{FactoryMap: factoryMap, AppConfig: config}
	return &c, nil
}
