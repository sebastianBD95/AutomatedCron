package infrastructure

import (
	"context"
	"fmt"

	db "github.com/sebastianBD95/AutomatedCron/App/configuration"
	m "github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveCron(_ context.Context, i interface{}) (interface{}, error) {

	logrus.Info("Saving cron data")
	c := i.(m.CronAutomated)
	db.InitConnection()
	collection := db.GetCollection("crons")
	result, err := collection.InsertOne(context.TODO(), c)

	if err != nil {
		fmt.Println(err)
	}

	c.ID = result.InsertedID.(primitive.ObjectID).Hex()
	logrus.Info("saving succesfully")
	db.Disconnect()
	return c, nil
}
