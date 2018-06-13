package scheduler

import (
	"errors"
	"fmt"
	"github.com/W1llyu/ourcrawler/collector"
	"github.com/W1llyu/ourcrawler/downloader"
	"github.com/W1llyu/ourcrawler/processor"
	"log"
	"net/http"
)

type Task struct {
	Request    *http.Request
	Downloader downloader.Downloader
	Collector  collector.Collector
	Processor  processor.Processor
}

func (t *Task) Run() {
	if t.Request == nil {
		failOnError(errors.New("nil Request"), "")
		return
	}
	if t.Processor == nil {
		failOnError(errors.New("nil Processor"), "")
		return
	}
	if t.Downloader == nil {
		t.Downloader = downloader.DefaultDownloader
	}
	if t.Collector == nil {
		t.Collector = collector.DefaultCollector
	}
	response := t.Downloader.Download(t.Request)
	resultItems, err := t.Processor.Process(response)
	if err != nil {
		failOnError(err, "")
		return
	}
	err = t.Collector.Process(resultItems)
	if err != nil {
		failOnError(err, "")
	}
}

func infof(format string, a ...interface{}) {
	log.Printf("[DOWNLOADER][INFO] %s", fmt.Sprintf(format, a...))
}

func failOnError(err error, msg string) {
	log.Printf("[DOWNLOADER][ERROR] %s %s", msg, err)
}
