package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/halapastefan/microservice/userservice"
	"github.com/halapastefan/microservice/userservice/db"
	srvc "github.com/halapastefan/microservice/userservice/implementaion"
	"github.com/halapastefan/microservice/userservice/repository"
)

type Main struct {
	Router *gin.Engine
}

func main() {

	main := Main{Router: gin.Default()}

	mongoClient := db.GetClient()
	repo := repository.New(mongoClient)
	// Constructing service
	var service userservice.Service
	{
		service = srvc.New(repo)
	}

	api := main.Router.Group("/api")
	{
		api.GET("/users", service.GetAllUsers)
		api.GET("/user/:id", service.GetUser)
		api.POST("/user", service.CreateUser)
		api.DELETE("/user/:id", service.DeleteUser)
		api.PUT("/user/:id", service.UpdateUser)
	}

	err := main.Router.Run(":3300")

	if err != nil {
		fmt.Println(err)
	}
}
