package crawler

import (
	"github.com/W1llyu/ourcrawler/scheduler"
	"log"
	"time"
)

type Crawler struct {
	ch chan *scheduler.Task
}

var (
	crw = &Crawler{}
)

func GetCrawler() *Crawler {
	return crw
}

func EnqueueTask(task *scheduler.Task) {
	select {
	case <-time.After(time.Second * 10):
		log.Printf("[CRAWLER][ERROR] task enqueue timeout")
	case crw.ch <- task:
		log.Println("[CRAWLER][SUCCESS] task enqueue success")
	}
}

func (c *Crawler) Start() {
	c.ch = make(chan *scheduler.Task, 256)
	go c.consume()
}

func (c *Crawler) Stop() {
	close(c.ch)
}

func (c *Crawler) consume() {
	for task := range c.ch {
		go task.Run()
	}
}
