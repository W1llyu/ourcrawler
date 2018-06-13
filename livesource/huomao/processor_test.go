package huomao

import (
	"testing"
)

func TestProcess(t *testing.T) {
	task := NewTask("https://www.huomao.com/channels/channel.json?page=1&game_url_rule=dota2")
	task.Run()
}
