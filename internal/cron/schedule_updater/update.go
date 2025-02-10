package schedule_updater

import (
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func (j *CronJob) Update() {
	err := j.adminService.UpdateClasses(context.Background())
	if err != nil {
		j.logger.Errorw("Unhandled error occurred",
			zap.Error(err),
		)
	}
	j.logger.Info("Job completed successfully")
}

func (j *CronJob) Run() *cron.Cron {
	c := cron.New()

	_, err := c.AddFunc("0 * * * *", j.Update)
	if err != nil {
		panic(err)
	}

	c.Start()
	return c
}
