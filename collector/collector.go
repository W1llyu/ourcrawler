package collector

import (
	"net/http"
)

var (
	DefaultCollector = NewConsoleCollector()
)

type Collector interface {
	Process(resultItems ResultItems) error
}

type ResultItems struct {
	Results map[string]interface{}
	Request *http.Request
}

func (r *ResultItems) Put(key string, value interface{}) {
	if r.Results == nil {
		r.Results = make(map[string]interface{})
	}
	r.Results[key] = value
}
