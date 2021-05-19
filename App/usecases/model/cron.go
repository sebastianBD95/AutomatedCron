package model

import (
	"github.com/mileusna/crontab"
)

type CronAutomated struct {
	ID     string            `schema:"id"`
	Name   string            `schema:"name"`
	Url    string            `schema:"url"`
	Header map[string]string `schema:"header"`
	Date   string            `schema:"result"`
	Cron   crontab.Crontab   `schema:"-"`
}
