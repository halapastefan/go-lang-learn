package userservice

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Repository interface {
	GetAllUsers() []*User
	GetUser(id int) User
	CreateUser(user User) User
	DeleteUser(id int) User
	UpdateUser(user User, id int) User
}
