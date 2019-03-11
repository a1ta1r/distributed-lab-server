package controller

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/service"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	teamService service.TeamService
}

func NewTeamController(teamService service.TeamService) TeamController {
	return TeamController{teamService: teamService}
}

func (uc TeamController) AddTeam(ctx *gin.Context) {
	var team entity.Team
	err := ctx.ShouldBindJSON(&team)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to team"})
		return
	}
	team = uc.teamService.AddTeam(team)
	ctx.JSON(201, gin.H{"team": team})
}

func (uc TeamController) GetTeams(ctx *gin.Context) {
	teams := uc.teamService.GetTeams()
	ctx.JSON(200, gin.H{"teams": teams})
}

func (uc TeamController) GetTeam(ctx *gin.Context) {
	var team entity.Team
	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}
	team.ID = id
	team, err := uc.teamService.GetTeam(team)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "team not found"})
		return
	}
	ctx.JSON(200, gin.H{"team": team})
}

func (uc TeamController) UpdateTeam(ctx *gin.Context) {
	var team entity.Team

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	err := ctx.ShouldBindJSON(&team)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to team"})
		return
	}

	team.ID = id
	team, err = uc.teamService.UpdateTeam(team)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "team not found"})
		return
	}
	ctx.JSON(200, gin.H{"team": team})
}

func (uc TeamController) DeleteTeam(ctx *gin.Context) {
	var team entity.Team

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	team.ID = id
	team, err := uc.teamService.GetTeam(team)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "team not found"})
		return
	}
	uc.teamService.DeleteTeam(team)
	ctx.JSON(200, gin.H{"team": "deleted"})
}
