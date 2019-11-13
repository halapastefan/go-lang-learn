package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type key string

const (
	usernameKey = key("reactive")
	passwordKey = key("reactive")
)

func GetClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions.SetAuth(options.Credential{
		Username: string(usernameKey),
		Password: string(passwordKey),
	})

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx.Value(passwordKey)
	ctx.Value(usernameKey)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
