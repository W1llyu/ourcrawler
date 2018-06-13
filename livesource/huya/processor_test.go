package huya

import (
	"testing"
)

func TestProcess(t *testing.T) {
	task := NewTask("https://www.huya.com/g/dota2")
	task.Run()
}
