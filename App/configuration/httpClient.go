package configuration

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func ProccessRequest(id, name, verb, url, body, headers, err string) {

}

func ResolveRequest(url, body, headers, verb string) {

	switch {
	case verb == "GET":
		doGet(url, headers)
	}
}

//do a Get Request
func doGet(url, headers string) {
	fmt.Println("Doing GET")
	response, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(response)
}
