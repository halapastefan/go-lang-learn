package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/user", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	err := router.Run(":3300")

	if err != nil {
		fmt.Println("Error", err)
	}
}
