package infrastructure

import (
	"context"
	"fmt"

	db "github.com/sebastianBD95/AutomatedCron/App/configuration"
	m "github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveCron(_ context.Context, i interface{}) (interface{}, error) {

	fmt.Println("infrastructure")
	c := i.(m.CronAutomated)

	collection := db.GetCollection("crons")
	result, err := collection.InsertOne(context.TODO(), c)

	if err != nil {
		fmt.Println(err)
	}

	c.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return c, nil
}
