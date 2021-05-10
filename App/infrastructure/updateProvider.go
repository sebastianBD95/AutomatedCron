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

func UpdateCron(_ context.Context, i interface{}) (interface{}, error) {

	logrus.Info("Updating cron data")
	c := i.(m.CronAutomated)
	logrus.Error(c)
	objId, _ := primitive.ObjectIDFromHex(c.ID)

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"name": c.Name, "url": c.Url}}
	db.InitConnection()
	collection := db.GetCollection("crons")
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Warn("Mongo : ", err)
	}
	logrus.Warn(updateResult.UpsertedID)
	logrus.Info("saving succesfully")
	db.Disconnect()
	_, _ = c, updateResult
	return c, nil
}
