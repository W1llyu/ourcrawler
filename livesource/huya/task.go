package huya

import (
	"github.com/W1llyu/ourcrawler/scheduler"
	"net/http"
	"github.com/W1llyu/ourcrawler/collector"
)

func NewTask(url string) *scheduler.Task {
	request, _ := http.NewRequest("GET", url, nil)
	return &scheduler.Task{
		Request:   request,
		Processor: &Processor{},
		Collector: collector.NewFileCollector("/tmp/live_data/huya"),
	}
}

