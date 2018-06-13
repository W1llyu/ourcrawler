package panda

import "testing"

func TestProcess(t *testing.T) {
	task := NewTask("https://www.panda.tv/cate/starcraft")
	task.Run()
}
