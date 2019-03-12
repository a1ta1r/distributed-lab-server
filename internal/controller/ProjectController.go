package controller

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/service"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	ProjectService service.ProjectService
}

func NewProjectController(ProjectService service.ProjectService) ProjectController {
	return ProjectController{ProjectService: ProjectService}
}

func (uc ProjectController) AddProject(ctx *gin.Context) {
	var Project entity.Project
	err := ctx.ShouldBindJSON(&Project)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to Project"})
		return
	}
	Project = uc.ProjectService.AddProject(Project)
	ctx.JSON(201, gin.H{"Project": Project})
}

func (uc ProjectController) GetProjects(ctx *gin.Context) {
	Projects := uc.ProjectService.GetProjects()
	ctx.JSON(200, gin.H{"Projects": Projects})
}

func (uc ProjectController) GetProject(ctx *gin.Context) {
	var Project entity.Project
	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}
	Project.ID = id
	Project, err := uc.ProjectService.GetProject(Project)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Project not found"})
		return
	}
	ctx.JSON(200, gin.H{"Project": Project})
}

func (uc ProjectController) UpdateProject(ctx *gin.Context) {
	var Project entity.Project

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	err := ctx.ShouldBindJSON(&Project)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to Project"})
		return
	}

	Project.ID = id
	Project, err = uc.ProjectService.UpdateProject(Project)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Project not found"})
		return
	}
	ctx.JSON(200, gin.H{"Project": Project})
}

func (uc ProjectController) DeleteProject(ctx *gin.Context) {
	var Project entity.Project

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	Project.ID = id
	Project, err := uc.ProjectService.GetProject(Project)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Project not found"})
		return
	}
	uc.ProjectService.DeleteProject(Project)
	ctx.JSON(200, gin.H{"Project": "deleted"})
}
