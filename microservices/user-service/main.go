package main

import (
	"./controller"
	"fmt"
	"github.com/gin-gonic/gin"
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

	err := main.Router.Run(":3300")

	if err != nil {
		fmt.Println(err)
	}
}
