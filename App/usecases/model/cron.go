package model

type CronAutomated struct {
	ID     int               `json: "id"`
	Name   string            `Json: "name"`
	Url    string            `Json: "url"`
	Header map[string]string `json: "header"`
	Result string            `json: result`
}
