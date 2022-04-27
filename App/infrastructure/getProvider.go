package infrastructure

import (
	"context"

	db "github.com/sebastianBD95/AutomatedCron/App/configuration"
	m "github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCron(_ context.Context, i interface{}) (interface{}, error) {

	logrus.Info("Getting cron data")
	c := i.(*m.CronAutomated)
	var cronA m.CronAutomated

	objId, _ := primitive.ObjectIDFromHex(c.ID)

	filter := bson.M{"_id": objId, "name": c.Name}

	db.InitConnection()
	collection := db.GetCollection("crons")
	err := collection.FindOne(context.TODO(), filter).Decode(&cronA)

	if err != nil {
		logrus.Warn("Mongo : ", err)
		return nil, err
	}

	logrus.Info("saving succesfully")
	db.Disconnect()
	_ = c
	return cronA, nil
}

func GetAllCrons() ([]m.CronAutomated, error) {

	logrus.Info("Getting all crons!!")
	db.InitConnection()
	collection := db.GetCollection("crons")

	cursor, err := collection.Find(context.TODO(), bson.M{})

	var results []m.CronAutomated
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		logrus.Error(err)
	}
	for _, result := range results {
		logrus.Info(result)
	}

	return results, nil
}
