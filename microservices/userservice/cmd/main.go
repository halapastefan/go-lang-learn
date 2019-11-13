package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/halapastefan/microservice/userservice"
	"github.com/halapastefan/microservice/userservice/controller"
	srvc "github.com/halapastefan/microservice/userservice/implementaion"
	"github.com/halapastefan/microservice/userservice/repository"
)

type Main struct {
	Router *gin.Engine
}

func main() {

	main := Main{Router: gin.Default()}

	c := controller.User{}

	api := main.Router.Group("/api")
	{
		api.GET("/users", c.GetAllUsers)
		api.GET("/user/:id", c.GetUser)
		api.POST("/user", c.CreateUser)
		api.DELETE("/user/:id", c.DeleteUser)
		api.PUT("/user/:id", c.UpdateUser)
	}

	repo := repository.NewRepository()
	// Constructing service
	var service userservice.Service
	{
		service = srvc.NewService(repo)
	}

	fmt.Println("Inicialized service", service)

	err := main.Router.Run(":3300")

	if err != nil {
		fmt.Println(err)
	}
}
