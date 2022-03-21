package configuration

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientMongo *mongo.Client

func GetCollection(coll string) *mongo.Collection {

	return ClientMongo.Database("cron_automated").Collection(coll)
}

func InitConnection() {

	clientOptions := options.Client().ApplyURI("mongodb://cronUser:cronPass@mongo:27017/cron_automated")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logrus.Error("Mongo: " + err.Error())

	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logrus.Error("Mongo: " + err.Error())
	}

	ClientMongo = client
	logrus.Info("Connection succesfully !!")

}

func Disconnect() {

	err := ClientMongo.Disconnect(context.TODO())

	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("Connection succesfully close !!")
}
