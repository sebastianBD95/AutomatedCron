package presenter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sebastianBD95/AutomatedCron/App/usecases"
	"github.com/sebastianBD95/AutomatedCron/App/usecases/model"
)

func CreateCron(w http.ResponseWriter, r *http.Request) {

	var cr model.CronAutomated
	var err error

	fmt.Println("Controler")
	err = json.NewDecoder(r.Body).Decode(&cr)
	cr, err = usecases.CreateCronUsesCase(cr)

	if nil != err {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "%+v \n", cr)
}
