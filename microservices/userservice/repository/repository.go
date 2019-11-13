package repository

import (
	"github.com/halapastefan/microservice/userservice"
	"github.com/halapastefan/microservice/userservice/model"
)

type Repository struct {
}

func NewRepository() userservice.Repository {
	return Repository{}
}

func (r Repository) GetAllUsers() []model.User {
	panic("implement me")
}

func (r Repository) GetUser(id int) model.User {
	panic("implement me")
}

func (r Repository) CreateUser(user model.User) model.User {
	panic("implement me")
}

func (r Repository) DeleteUser(id int) model.User {
	panic("implement me")
}

func (r Repository) UpdateUser(user model.User, id int) model.User {
	panic("implement me")
}
