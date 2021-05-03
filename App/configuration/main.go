package main

import (
	"fmt"
	"log"
	"net/http"

	cont "github.com/sebastianBD95/AutomatedCron/App/presenter"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome")
}

func main() {

	fmt.Println("init")

	r := mux.NewRouter()
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/cron", cont.CreateCron).
		Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
