package infrastructure

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"
	db "github.com/sebastianBD95/AutomatedCron/App/configuration"
	m "github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCron(_ context.Context, i interface{}) (interface{}, error) {

	logrus.Info("Saving cron data")
	c := i.(m.CronAutomated)
	var cronA m.CronAutomated

	filter := bson.D{}

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
