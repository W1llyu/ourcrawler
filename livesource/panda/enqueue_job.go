package panda

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
		&Job{"https://www.panda.tv/cate/dota2", globalSpec},
		&Job{"https://www.panda.tv/cate/lol", globalSpec},
		&Job{"https://www.panda.tv/cate/csgo", globalSpec},
		&Job{"https://www.panda.tv/cate/overwatch", globalSpec},
		&Job{"https://www.panda.tv/cate/starcraft", globalSpec},
		&Job{"https://www.panda.tv/cate/pubg", globalSpec},
		&Job{"https://www.panda.tv/cate/hearthstone", globalSpec},
		&Job{"https://www.panda.tv/cate/war3", globalSpec},
		&Job{"https://www.panda.tv/cate/kingglory", globalSpec},
		&Job{"https://www.panda.tv/cate/heroes", globalSpec},
		&Job{"https://www.panda.tv/cate/fifa", globalSpec},
		&Job{"https://www.panda.tv/cate/cf", globalSpec},
		&Job{"https://www.panda.tv/cate/cfm", globalSpec},
		&Job{"https://www.panda.tv/cate/sycj", globalSpec},
		&Job{"https://www.panda.tv/cate/zjz", globalSpec},
	}
}
