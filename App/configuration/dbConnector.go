package configuration

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientMongo *mongo.Client

func GetCollection(coll string) *mongo.Collection {

	return ClientMongo.Database("cron_automated").Collection(coll)
}

func InitConnection() {

	clientOptions := options.Client().ApplyURI("mongodb://cronUser:cronPass@localhost:27017/cron_automated")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("Mongo: " + err.Error())

	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("Mongo: " + err.Error())
	}

	ClientMongo = client
	fmt.Println("Connection succesfully !!")

}

func Disconnect() {

	err := ClientMongo.Disconnect(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection succesfully close !!")
}
