package controller

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return TaskController{taskService: taskService}
}

func (uc TaskController) AddTask(ctx *gin.Context) {
	var task entity.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to task"})
		return
	}
	task = uc.taskService.AddTask(task)
	ctx.JSON(201, gin.H{"task": task})
}

func (uc TaskController) GetTasks(ctx *gin.Context) {
	tasks := uc.taskService.GetTasks()
	ctx.JSON(200, gin.H{"tasks": tasks})
}

func (uc TaskController) GetTask(ctx *gin.Context) {
	var task entity.Task
	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}
	task.ID = id
	task, err := uc.taskService.GetTask(task)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "task not found"})
		return
	}
	ctx.JSON(200, gin.H{"task": task})
}

func (uc TaskController) UpdateTask(ctx *gin.Context) {
	var task entity.Task

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to task"})
		return
	}

	task.ID = id
	task, err = uc.taskService.UpdateTask(task)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "task not found"})
		return
	}
	ctx.JSON(200, gin.H{"task": task})
}

func (uc TaskController) DeleteTask(ctx *gin.Context) {
	var task entity.Task

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	task.ID = id
	task, err := uc.taskService.GetTask(task)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "task not found"})
		return
	}
	uc.taskService.DeleteTask(task)
	ctx.JSON(200, gin.H{"task": "deleted"})
}
