package usecases

import (
	"github.com/mileusna/crontab"
	hc "github.com/sebastianBD95/AutomatedCron/App/configuration"
	"github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"github.com/sirupsen/logrus"
)

func ActivateCron(cr model.CronAutomated) {

	logrus.Info("Activate Cron: " + cr.Name)

	cr.Cron = crontab.New()
	cr.Cron.MustAddJob(cr.Date, callRequest, cr.Url, "", "", cr.Verb)

}

func callRequest(url, body, header, verb string) {

	logrus.Error(url)
	hc.ResolveRequest(url, body, header, verb)
}
