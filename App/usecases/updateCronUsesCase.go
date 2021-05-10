package usecases

import (
	"context"
	"time"

	"github.com/reactivex/rxgo/v2"
	"github.com/sebastianBD95/AutomatedCron/App/infrastructure"
	"github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"github.com/sirupsen/logrus"
)

func UpdateCronUsesCase(c interface{}) (model.CronAutomated, error) {

	logrus.Info("Update  channel")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ch := make(chan rxgo.Item, 10)

	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy()).
		Map(infrastructure.UpdateCron)

	observable.Connect(ctx)
	ch <- rxgo.Of(c)
	close(ch)

	var cr model.CronAutomated
	var err error
	for items := range observable.Observe() {

		if items.Error() {
			err = items.E
		} else {
			cr = items.V.(model.CronAutomated)
		}

	}
	logrus.Info("Cron Updated with Id: " + cr.ID)
	return cr, err

}
