package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"os"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	csvFile, err := os.Open("callingCodes.csv")

	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	t := CreateTrie(csvFile)
	elapsed := start.Sub(time.Now())
	fmt.Printf("Trie inicialization %s\n", elapsed)

	router.GET("/call/:calling/:date", func(context *gin.Context) {
		calling := context.Param("calling")
		date := context.Param("date")

		findPrice(t, calling, date)

		context.Status(http.StatusOK)
	})

	err2 := router.Run(":3300")

	if err2 != nil {
		fmt.Println(err)
	}
}