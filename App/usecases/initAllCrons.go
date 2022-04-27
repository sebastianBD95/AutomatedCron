package usecases

import (
	"github.com/sebastianBD95/AutomatedCron/App/infrastructure"
	"github.com/sirupsen/logrus"
)

func InitCronUsesCase() {

	crons, err := infrastructure.GetAllCrons()
	if err != nil {
		logrus.Error(err)
	}

	for _, element := range crons {
		ActivateCron(element)
	}

}
