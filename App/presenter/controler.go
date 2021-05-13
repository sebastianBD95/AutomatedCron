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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something badhappened !"))
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
	_ = data
	if err != nil {
		logrus.Error(err)
	}

	cronA := new(model.CronAutomated)

	if err := schema.NewDecoder().Decode(cronA, r.Form); err != nil {
		logrus.Error(err)
	}

	cr, err = usecases.GetCronUsesCase(cronA)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Something badhappened !"))
	}
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

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Something badhappened !"))
	}
	_ = data
	json.NewEncoder(w).Encode(cr)
}

func UpdateCron(w http.ResponseWriter, r *http.Request) {

	var cr model.CronAutomated
	var err error

	q := r.URL.Query()

	err = json.NewDecoder(r.Body).Decode(&cr)
	cr.ID = q.Get("id")
	cr, err = usecases.UpdateCronUsesCase(cr)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Something badhappened !"))
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("204 - Update succesfully"))
}
