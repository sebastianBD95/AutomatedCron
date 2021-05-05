package configuration

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBmanager struct {
	clientMongo *mongo.Client
}

var dbmongo DBmanager

func New() DBmanager {

	return dbmongo
}

func GetCollection(coll string) *mongo.Collection {

	return dbmongo.clientMongo.Database("cron_automated").Collection(coll)
}

func (db DBmanager) InitConnection() {

	clientOptions := options.Client().ApplyURI("mongodb://cronUser:cronPass@localhost:27017/cron_automated")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("Mongo: " + err.Error())

	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("Mongo: " + err.Error())
	}

	db.clientMongo = client
	fmt.Println("Connection succesfully !!")

}

func (db DBmanager) disconnect() {

	err := db.clientMongo.Disconnect(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection succesfully close !!")
}
