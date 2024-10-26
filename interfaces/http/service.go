package http

import (
	"github.com/Systenix/taskmanager/task_queue_system/app/container/containerhelper"
	"github.com/Systenix/taskmanager/task_queue_system/app/logger"
	"go.uber.org/zap"
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/Systenix/taskmanager/task_queue_system/app/container/servicecontainer"
	"github.com/gin-gonic/gin"
)

type TaskManagerService struct {
	Container *servicecontainer.ServiceContainer
}

func catchPanic() {

	if p := recover(); p != nil {
		switch err := p.(type) {
		case string:
			logger.Log.Error("unrecoverable panic occurred", zap.String("err", err))
		case runtime.Error:
			logger.Log.Error("unrecoverable panic occurred", zap.String("err", err.Error()))
		case error:
			logger.Log.Error("unrecoverable panic occurred", zap.String("err", err.Error()))
		default:
			logger.Log.Error("unrecoverable error occurred", zap.String("err", err.(string)))
		}
		logger.Log.Debug("DEBUG TRACE", zap.String("stack", string(debug.Stack())))
	}
}

func (s *TaskManagerService) SubmitTask(ctx *gin.Context) {
	defer catchPanic()

	var req struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	tmuc, err := containerhelper.GetTaskQueueUseCase(s.Container)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	taskID, err := tmuc.SubmitTask(ctx, req.Type, req.Payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to submit task"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"task_id": taskID})
}

func (s *TaskManagerService) GetTaskStatus(ctx *gin.Context) {
	taskID := ctx.Param("task_id")

	tmuc, err := containerhelper.GetTaskQueueUseCase(s.Container)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	task, err := tmuc.GetTaskStatus(ctx, taskID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task": task})
}

func (s *TaskManagerService) CancelTask(ctx *gin.Context) {
	taskID := ctx.Param("task_id")

	tmuc, err := containerhelper.GetTaskQueueUseCase(s.Container)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	err = tmuc.CancelTask(ctx, taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "task canceled"})
}
