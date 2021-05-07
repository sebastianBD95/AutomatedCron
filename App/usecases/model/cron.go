package model

type CronAutomated struct {
	ID     string            `schema: "id"`
	Name   string            `schema: "name"`
	Url    string            `schema: "url"`
	Header map[string]string `schema: "header"`
	Result string            `schema: "result"`
}
