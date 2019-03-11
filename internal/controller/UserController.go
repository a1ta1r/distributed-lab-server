package controller

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService: userService}
}

func (uc UserController) AddUser(ctx *gin.Context) {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to user"})
		return
	}
	user = uc.userService.AddUser(user)
	ctx.JSON(201, gin.H{"user": user})
}

func (uc UserController) GetUsers(ctx *gin.Context) {
	users := uc.userService.GetUsers()
	ctx.JSON(200, gin.H{"users": users})
}

func (uc UserController) GetUser(ctx *gin.Context) {
	var user entity.User
	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}
	user.ID = id
	user, err := uc.userService.GetUser(user)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(200, gin.H{"user": user})
}

func (uc UserController) UpdateUser(ctx *gin.Context) {
	var user entity.User

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot map given fields to user"})
		return
	}

	user.ID = id
	user, err = uc.userService.UpdateUser(user)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(200, gin.H{"user": user})
}

func (uc UserController) DeleteUser(ctx *gin.Context) {
	var user entity.User

	id := getIdFromQueryString(ctx)
	if id == 0 {
		ctx.AbortWithStatusJSON(422, gin.H{"error": "cannot parse id, please use uint"})
		return
	}

	user.ID = id
	user, err := uc.userService.GetUser(user)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	uc.userService.DeleteUser(user)
	ctx.JSON(200, gin.H{"user": "deleted"})
}
