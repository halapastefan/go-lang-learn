package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/halapastefan/microservice/userservice/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type User struct {
}

func (u User) GetAllUsers(context *gin.Context) {

	usr1 := model.User{FirstName: "st", LastName: "te"}
	usr2 := model.User{FirstName: "st", LastName: "te"}

	users := []model.User{usr1, usr2}

	log.Info("Users found")
	context.JSON(http.StatusOK, users)
}

func (User) GetUser(ctx *gin.Context) {

	userId := ctx.Param("id")
	usr := model.User{
		LastName:  "Stefan",
		FirstName: "Halapa",
	}

	log.Info("User with user id: ", userId, " found")
	ctx.JSON(http.StatusOK, usr)
}

func (User) CreateUser(ctx *gin.Context) {

	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Info(user)
	ctx.JSON(http.StatusOK, user)
}

func (User) DeleteUser(ctx *gin.Context) {

	userId := ctx.Param("id")
	log.Info("Try to delete user with id: ", userId)

	ctx.Status(http.StatusOK)
}

func (User) UpdateUser(ctx *gin.Context) {

}
