package scheduler

import (
	"github.com/robfig/cron"
)

type CronScheduler struct {
	c *cron.Cron
}

func NewCronScheduler() *CronScheduler {
	return &CronScheduler{
		c: cron.New(),
	}
}

func (c *CronScheduler) AddJobs(jobs []Job) {
	for _, job := range jobs {
		c.c.AddJob(job.Spec(), job)
	}
}

func (c *CronScheduler) Start() {
	c.c.Start()
}

func (c *CronScheduler) Stop() {
	c.c.Stop()
}
