package presenter

import (
	"encoding/json"
	"fmt"
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

	if err := r.ParseForm(); err != nil {
		logrus.Error(err)
	}

	data, err := json.Marshal(r.Form)
	logrus.Info(data)
	if err != nil {
		logrus.Error(err)
	}

	cronA := new(model.CronAutomated)

	if err := schema.NewDecoder().Decode(cronA, r.Form); err != nil {
		logrus.Error(err)
	}

	fmt.Println(cronA)
}
