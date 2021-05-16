package model

import "errors"

type httpVerb string

const (
	GET    httpVerb = "GET"
	POST            = "POST"
	PUT             = "PUT"
	DELETE          = "DELETE"
)

func (h httpVerb) IsValid() error {

	switch h {
	case GET, POST, PUT, DELETE:
		return nil
	}
	return errors.New("Invalid Value")
}
