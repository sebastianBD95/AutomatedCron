package configuration

import "fmt"

func ProccessRequest(id, name, verb, url, body, headers, err string) {

}

func ResolveRequest(url, body, headers, verb string) {

	switch {
	case verb == "GET":
		doGet(url, headers)
	}
}

//do a Get Request
func doGet(url, headers) {
	fmt.Println("Doing GET")
}
