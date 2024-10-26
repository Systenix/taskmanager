package datastorefactory

import (
	"context"

	"github.com/Systenix/taskmanager/task_queue_system/app/config"
	"github.com/Systenix/taskmanager/task_queue_system/app/container"
	"github.com/Systenix/taskmanager/task_queue_system/app/logger"
	redisv9 "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type redisFactory struct{}

// Build creates a new instance of the Redis client
func (rF *redisFactory) Build(c container.Container, dsc *config.DataStoreConfig) (DataStoreInterface, error) {
	key := dsc.Code
	if !dsc.Tx {
		if value, found := c.Get(key); found {
			logger.Log.Debug("found connPool in container for", zap.String("found key", key))
			return value, nil
		}
	}
	connPool := redisv9.NewClient(&redisv9.Options{
		Addr: dsc.Address,
	})
	if err := connPool.Ping(context.Background()).Err(); err != nil {
		logger.Log.Error("Error while building Redis client", zap.String("err", err.Error()))
		return nil, err
	}
	if !dsc.Tx {
		c.Put(key, connPool)
	}
	// Return the Redis client instead of the data service
	return connPool, nil
}
