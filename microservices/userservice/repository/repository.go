package repository

import (
	"context"
	"github.com/halapastefan/microservice/userservice"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const DB_NAME string = "golearn"

type repository struct {
	db mongo.Client
}

func New(client *mongo.Client) userservice.Repository {
	return &repository{db: *client}
}

func (r repository) GetAllUsers() []*userservice.User {

	var users []*userservice.User

	collection := r.db.Database(DB_NAME).Collection("data")
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for cur.Next(context.TODO()) {
		var user userservice.User
		err = cur.Decode(&user)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		users = append(users, &user)
	}

	return users
}

func (r repository) GetUser(id int) userservice.User {
	panic("implement me")
}

func (r repository) CreateUser(user userservice.User) userservice.User {
	panic("implement me")
}

func (r repository) DeleteUser(id int) userservice.User {
	panic("implement me")
}

func (r repository) UpdateUser(user userservice.User, id int) userservice.User {
	panic("implement me")
}
