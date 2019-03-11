package controller

import (
	"github.com/a1ta1r/distributed-lab-server/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService: userService}
}

func (uc UserController) GetUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"response": "Users"})
}
