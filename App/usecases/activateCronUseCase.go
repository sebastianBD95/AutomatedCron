package usecases

import (
	"github.com/mileusna/crontab"
	"github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"github.com/sirupsen/logrus"
)

func ActivateCron(cr model.CronAutomated) {

	logrus.Info("Creating Cron: " + cr.Name)

	URL := cr.Url
	c := crontab.New()
	c.MustAddJob("*/1 * * * *", callRequest, URL)

}

func callRequest(url string) {

	logrus.Error(url)
}
