package model

type CronAutomated struct {
	ID     string            `json: "id"`
	Name   string            `Json: "name"`
	Url    string            `Json: "url"`
	Header map[string]string `json: "header"`
	Result string            `json: result`
}
