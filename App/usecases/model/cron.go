package model

import (
	"github.com/mileusna/crontab"
)

type CronAutomated struct {
	ID     string            `schema:"id"`
	Name   string            `schema:"name"`
	Url    string            `schema:"url"`
	Verb   string            `schema:"verb"`
	Header map[string]string `schema:"header"`
	Date   string            `schema:"date"`
	Cron   *crontab.Crontab  `json:"-"`
}
