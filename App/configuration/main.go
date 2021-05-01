package main

import (
	"fmt"

	cont "github.com/sebastianBD95/AutomatedCron/presenter"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("init")

	r := mux.NewRouter()
	r.HandleFunc("/cron", cont.CreateCron).
		Methods("POST")
}
