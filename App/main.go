package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/sebastianBD95/AutomatedCron/App/configuration"
	cont "github.com/sebastianBD95/AutomatedCron/App/presenter"
)

func homeLink(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome")
}

func main() {

	fmt.Println("init db connection")

	dbmongo := db.New()
	dbmongo.InitConnection()

	fmt.Println("init endPoints")
	r := mux.NewRouter()
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/cron", cont.CreateCron).
		Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
