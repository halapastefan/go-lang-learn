package implementaion

import (
	"../model"
	usrsrvc "../userservice"
)

type service struct {
	repo usrsrvc.Repository
}

func NewService(repository usrsrvc.Repository) usrsrvc.Service {
	return &service{repo: repository}
}

func (*service) GetAllUsers() []model.User {
	return nil
}

func (i *service) GetUser(id int) model.User {
	panic("implement me")
}

func (i *service) CreateUser(user model.User) model.User {
	panic("implement me")
}

func (i *service) DeleteUser(id int) model.User {
	panic("implement me")
}

func (i *service) UpdateUser(user model.User, id int) model.User {
	panic("implement me")
}
