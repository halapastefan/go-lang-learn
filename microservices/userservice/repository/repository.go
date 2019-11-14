package repository

import (
	"context"
	"github.com/halapastefan/microservice/userservice"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DB_NAME string = "golearn"

type repository struct {
	db mongo.Client
}

func New(client *mongo.Client) userservice.Repository {
	return &repository{db: *client}
}

func (r repository) GetAllUsers() ([]*userservice.User, error) {

	var users []*userservice.User

	collection := r.db.Database(DB_NAME).Collection("data")
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Error("Error on Finding all the documents", err)
		return users, err
	}

	for cur.Next(context.TODO()) {
		var user userservice.User
		err = cur.Decode(&user)
		if err != nil {
			log.Error("Error on Decoding the document", err)
			return users, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r repository) GetUser(id string) (userservice.User, error) {

	collection := r.db.Database(DB_NAME).Collection("data")

	objID, _ := primitive.ObjectIDFromHex(id)
	reuslt := collection.FindOne(context.TODO(), bson.M{"_id": objID})

	var user userservice.User
	err := reuslt.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r repository) CreateUser(user *userservice.User) error {

	collection := r.db.Database(DB_NAME).Collection("data")

	_, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteUser(id string) error {

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.GetCollection().DeleteOne(context.TODO(), bson.M{"_id": objID})

	if err != nil {
		return err
	}
	return nil
}

func (r repository) UpdateUser(user userservice.User, id int) userservice.User {
	panic("implement me")
}

func (r repository) GetCollection() *mongo.Collection {
	return r.db.Database(DB_NAME).Collection("data")
}
