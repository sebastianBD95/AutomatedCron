package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
	"github.com/sebastianBD95/AutomatedCron/App/infrastructure"
	"github.com/sebastianBD95/AutomatedCron/App/usecases/model"
)

func CreateCronUsesCase(c interface{}) (model.CronAutomated, error) {

	fmt.Println("usesCases")
	fmt.Println("c:", c)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ch := make(chan rxgo.Item, 10)

	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy()).
		Map(infrastructure.SaveCron)

	observable.Connect(ctx)
	ch <- rxgo.Of(c)
	close(ch)
	fmt.Println("procesando")

	var cr model.CronAutomated
	var err error
	for items := range observable.Observe() {

		if items.Error() {
			err = items.E
		} else {
			cr = items.V.(model.CronAutomated)
		}

	}

	return cr, err

}
