package huomao

import (
	"github.com/W1llyu/ourcrawler/crawler"
	"github.com/W1llyu/ourcrawler/scheduler"
)

const (
	globalSpec = "0 1 * * * ?"
)

type Job struct {
	url string
	spec string
}

func (j *Job) Run() {
	task := NewTask(j.url)
	crawler.EnqueueTask(task)
}

func (j *Job) Spec() string {
	return j.spec
}

func NewJobs() []scheduler.Job {
	return []scheduler.Job {
		&Job{"https://www.huomao.com/channels/channel.json?page=1&game_url_rule=dota2&cate=dota2", globalSpec},
		&Job{"https://www.huomao.com/channels/channel.json?page=1&game_url_rule=lol&cate=英雄联盟", globalSpec},
		&Job{"https://www.huomao.com/channels/channel.json?page=1&game_url_rule=CSGO&cate=csgo", globalSpec},
		&Job{"http://www.huomao.com/channels/channel.json?page=1&game_url_rule=battlegrounds&cate=绝地求生", globalSpec},
		&Job{"http://www.huomao.com/channels/channel.json?page=1&game_url_rule=all&labelid=93&cate=王者荣耀", globalSpec},
		&Job{"https://www.huomao.com/channels/channel.json?page=1&game_url_rule=FIFAOnline3&cate=FIFAOnline", globalSpec},
	}
}
