package presenter

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCron(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println("done")
	_ = vars
}
