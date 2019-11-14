package userservice

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Repository interface {
	GetAllUsers() ([]*User, error)
	GetUser(id string) (User, error)
	CreateUser(user *User) error
	DeleteUser(id string) error
	UpdateUser(user User, id int) User
}
