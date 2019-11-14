package implementaion

import (
	"github.com/gin-gonic/gin"
	"github.com/halapastefan/microservice/userservice"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type service struct {
	repo userservice.Repository
}

func New(repository userservice.Repository) userservice.Service {
	return &service{repo: repository}
}

func (s *service) GetAllUsers(context *gin.Context) {

	users, err := s.repo.GetAllUsers()

	if err != nil {
		context.Status(http.StatusNotFound)
		return
	}

	log.Info("Users found")
	context.JSON(http.StatusOK, users)
}

func (s *service) GetUser(ctx *gin.Context) {

	userId := ctx.Param("id")

	usr, err := s.repo.GetUser(userId)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	log.Info("User with user id: ", userId, " found")
	ctx.JSON(http.StatusOK, usr)
}

func (s *service) CreateUser(ctx *gin.Context) {

	var user userservice.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Info(user)

	err := s.repo.CreateUser(&user)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *service) DeleteUser(ctx *gin.Context) {

	userId := ctx.Param("id")
	log.Info("Try to delete user with id: ", userId)
	err := s.repo.DeleteUser(userId)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusOK)
}

func (*service) UpdateUser(ctx *gin.Context) {

}
