package huomao

import (
	"github.com/W1llyu/ourcrawler/collector"
	"github.com/W1llyu/ourcrawler/scheduler"
	"net/http"
)

func NewTask(url string) *scheduler.Task {
	request, _ := http.NewRequest("GET", url, nil)
	return &scheduler.Task{
		Request:   request,
		Processor: &Processor{},
		Collector: collector.NewFileCollector("/tmp/live_data/huomao"),
	}
}
