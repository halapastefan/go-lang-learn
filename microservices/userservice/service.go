package userservice

import (
	"github.com/halapastefan/microservice/userservice/model"
)

type Service interface {
	GetAllUsers() []model.User
	GetUser(id int) model.User
	CreateUser(user model.User) model.User
	DeleteUser(id int) model.User
	UpdateUser(user model.User, id int) model.User
}
