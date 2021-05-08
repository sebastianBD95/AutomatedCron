package infrastructure

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"
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
		log.Warn("Mongo : ", err)
	}

	logrus.Info("saving succesfully")
	db.Disconnect()
	_ = c
	return cronA, nil
}
