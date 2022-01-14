package usecases

import (
	"github.com/robfig/cron/v3"
	hc "github.com/sebastianBD95/AutomatedCron/App/configuration"
	"github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"github.com/sirupsen/logrus"
)

func ActivateCron(cr model.CronAutomated) {

	logrus.Info("Activate Cron: " + cr.Name)

	cr.Cron = cron.New()
	cr.Cron.AddFunc(cr.Date, func() { callRequest(cr.Url, "", "", cr.Verb) })
	cr.Cron.Start()
}

func callRequest(url, body, header, verb string) {

	hc.ResolveRequest(url, body, header, verb)
}
