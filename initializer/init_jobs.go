package initializer

import (
	"github.com/W1llyu/ourcrawler/livesource/douyu"
	"github.com/W1llyu/ourcrawler/livesource/huomao"
	"github.com/W1llyu/ourcrawler/livesource/huya"
	"github.com/W1llyu/ourcrawler/livesource/panda"
	"github.com/W1llyu/ourcrawler/scheduler"
)

func InitSchedulerJobs(s scheduler.Scheduler) {
	s.AddJobs(douyu.NewJobs())
	s.AddJobs(panda.NewJobs())
	s.AddJobs(huya.NewJobs())
	s.AddJobs(huomao.NewJobs())
}
