package custom_crons

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func StartCron() {
	fmt.Println("Cron started to run every day 9am")
	scheduler := gocron.NewScheduler(time.Now().Local().Location())
	scheduler.Every(1).Day().At("17:14").Do(RefreshDocuments)
}
