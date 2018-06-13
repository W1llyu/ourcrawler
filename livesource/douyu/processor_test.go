package douyu

import (
	"testing"
)

func TestProcess(t *testing.T) {
	task := NewTask("https://www.douyu.com/directory/game/DOTA2")
	task.Run()
}