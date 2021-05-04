package infrastructure

import (
	"context"
	"fmt"
	"strings"

	"github.com/sebastianBD95/AutomatedCron/App/usecases/model"
)

func SaveCron(_ context.Context, i interface{}) (interface{}, error) {

	fmt.Println("infrastructure")
	c := i.(model.CronAutomated)
	c.Name = strings.ToUpper(c.Name)
	return c, nil
}
