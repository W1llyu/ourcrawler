package collector

import "testing"

func TestFileCollectorProcess(t *testing.T) {
	c := NewFileCollector("../data")
	m := make(map[string]interface{})
	m["count"] = 1
	m["name"] = "jacky"
	c.Process(ResultItems{m})
}
