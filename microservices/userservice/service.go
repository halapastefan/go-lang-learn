package userservice

import "github.com/gin-gonic/gin"

type Service interface {
	GetAllUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}
