package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sebastianBD95/AutomatedCron/App/usecases"
	"github.com/sebastianBD95/AutomatedCron/App/usecases/model"
	"github.com/sirupsen/logrus"
)

func CreateCron(w http.ResponseWriter, r *http.Request) {

	var cr model.CronAutomated
	var err error

	err = json.NewDecoder(r.Body).Decode(&cr)
	cr, err = usecases.CreateCronUsesCase(cr)

	if nil != err {
		logrus.Error(err)
		return
	}

	json.NewEncoder(w).Encode(cr)
}

func GetCron(w http.ResponseWriter, r *http.Request) {

	logrus.Info("Controler")
	var cr model.CronAutomated
	if err := r.ParseForm(); err != nil {
		logrus.Error(err)
	}

	data, err := json.Marshal(r.Form)
	if err != nil {
		logrus.Error(err)
	}

	cronA := new(model.CronAutomated)

	if err := schema.NewDecoder().Decode(cronA, r.Form); err != nil {
		logrus.Error(err)
	}

	cr, err = usecases.GetCronUsesCase(cronA)

	_ = data
	json.NewEncoder(w).Encode(cr)
}

func DeleteCron(w http.ResponseWriter, r *http.Request) {

	logrus.Info("Controler")
	var cr model.CronAutomated
	if err := r.ParseForm(); err != nil {
		logrus.Error(err)
	}

	data, err := json.Marshal(r.Form)
	if err != nil {
		logrus.Error(err)
	}

	cronA := new(model.CronAutomated)

	if err := schema.NewDecoder().Decode(cronA, r.Form); err != nil {
		logrus.Error(err)
	}

	cr, err = usecases.DeleteCronUsesCase(cronA)

	_ = data
	json.NewEncoder(w).Encode(cr)
}
