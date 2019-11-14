package http

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/halapastefan/go-kit-micro/awesomeservice/transport"
	"net/http"
)

func NewService(endpoints Endpoints) {

	router := gin.Default()

	uppercaseHandler := func(ctx *gin.Context) {

	}

	api := router.Group("/api")
	{
		api.GET("/")
	}

}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
