package main

import (
	"github.com/W1llyu/ourcrawler/crawler"
	"github.com/W1llyu/ourcrawler/scheduler"
	"github.com/W1llyu/ourcrawler/initializer"
)

var (
	forever = make(chan bool)
)

func main() {
	crawler.GetCrawler().Start()
	newScheduler().Start()
	<-forever
}

func newScheduler() scheduler.Scheduler {
	s := scheduler.NewCronScheduler()
	initializer.InitSchedulerJobs(s)
	return s
}
