package controller

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/service"
	"github.com/gin-gonic/gin"
)

type StatusController struct {
	statusService service.StatusService
}

func NewStatusController(statusService service.StatusService) StatusController {
	return StatusController{statusService: statusService}
}

func (uc StatusController) AddStatus(ctx *gin.Context) {
	var status entity.Status
	err := ctx.ShouldBindJSON(&status)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to status"})
		return
	}
	status = uc.statusService.AddStatus(status)
	ctx.JSON(201, gin.H{"status": status})
}

func (uc StatusController) GetStatuses(ctx *gin.Context) {
	statuses := uc.statusService.GetStatuses()
	ctx.JSON(200, gin.H{"statuses": statuses})
}

func (uc StatusController) GetStatus(ctx *gin.Context) {
	var status entity.Status
	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}
	status.ID = id
	status, err := uc.statusService.GetStatus(status)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "status not found"})
		return
	}
	ctx.JSON(200, gin.H{"status": status})
}

func (uc StatusController) UpdateStatus(ctx *gin.Context) {
	var status entity.Status

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	err := ctx.ShouldBindJSON(&status)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to status"})
		return
	}

	status.ID = id
	status, err = uc.statusService.UpdateStatus(status)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "status not found"})
		return
	}
	ctx.JSON(200, gin.H{"status": status})
}

func (uc StatusController) DeleteStatus(ctx *gin.Context) {
	var status entity.Status

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	status.ID = id
	status, err := uc.statusService.GetStatus(status)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "status not found"})
		return
	}
	uc.statusService.DeleteStatus(status)
	ctx.JSON(200, gin.H{"status": "deleted"})
}
