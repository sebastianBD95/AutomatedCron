package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	fmt.Fprintf(w, "%+v \n", cr)
}
