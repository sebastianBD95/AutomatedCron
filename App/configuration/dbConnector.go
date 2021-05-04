package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBmanager struct {
	clientMongo *mongo.Client
}

func (db DBmanager) initConnection() {

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
