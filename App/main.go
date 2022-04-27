package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/sebastianBD95/AutomatedCron/App/configuration"
	cont "github.com/sebastianBD95/AutomatedCron/App/presenter"
	uses "github.com/sebastianBD95/AutomatedCron/App/usecases"
	"github.com/sirupsen/logrus"
)

func homeLink(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome")
}

func main() {

	logrus.Info("init endPoints")
	r := mux.NewRouter()

	r.HandleFunc("/cron", cont.CreateCron).
		Methods("POST")
	r.HandleFunc("/cron", cont.GetCron).
		Methods("GET")
	r.HandleFunc("/cron", cont.DeleteCron).
		Methods("DELETE")
	r.HandleFunc("/cron", cont.UpdateCron).
		Methods("PATCH")

	uses.InitCronUsesCase()

	logrus.Fatal(http.ListenAndServe(":8087", r))
}
