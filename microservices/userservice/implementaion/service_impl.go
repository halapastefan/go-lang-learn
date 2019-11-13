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

	//usr1 := userservice.User{FirstName: "st", LastName: "te"}
	//usr2 := userservice.User{FirstName: "st", LastName: "te"}

	users := s.repo.GetAllUsers()

	log.Info("Users found")
	context.JSON(http.StatusOK, users)
}

func (*service) GetUser(ctx *gin.Context) {

	userId := ctx.Param("id")
	usr := userservice.User{
		LastName:  "Stefan",
		FirstName: "Halapa",
	}

	log.Info("User with user id: ", userId, " found")
	ctx.JSON(http.StatusOK, usr)
}

func (*service) CreateUser(ctx *gin.Context) {

	var user userservice.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Info(user)
	ctx.JSON(http.StatusOK, user)
}

func (*service) DeleteUser(ctx *gin.Context) {

	userId := ctx.Param("id")
	log.Info("Try to delete user with id: ", userId)

	ctx.Status(http.StatusOK)
}

func (*service) UpdateUser(ctx *gin.Context) {

}
