package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getIdFromQueryString(ctx *gin.Context) uint {
	id64, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return 0
	}
	id := uint(id64)
	return id
}
