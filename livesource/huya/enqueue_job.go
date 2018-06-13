package huya

import (
	"github.com/W1llyu/ourcrawler/crawler"
	"github.com/W1llyu/ourcrawler/scheduler"
)

const (
	globalSpec = "0 1 * * * ?"
)

type Job struct {
	url  string
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
	return []scheduler.Job{
		&Job{"https://www.huya.com/g/dota2", globalSpec},
		&Job{"https://www.huya.com/g/lol", globalSpec},
		&Job{"https://www.huya.com/g/862", globalSpec},
		&Job{"https://www.huya.com/g/overwatch", globalSpec},
		&Job{"https://www.huya.com/g/starcraft", globalSpec},
		&Job{"https://www.huya.com/g/2793", globalSpec},
		&Job{"https://www.huya.com/g/hearthstone", globalSpec},
		&Job{"https://www.huya.com/g/wzry", globalSpec},
		&Job{"https://www.huya.com/g/qiuqiu", globalSpec},
		&Job{"https://www.huya.com/g/heroes", globalSpec},
		&Job{"https://www.huya.com/g/3379", globalSpec},
		&Job{"https://www.huya.com/g/3683", globalSpec},
		&Job{"https://www.huya.com/g/cf", globalSpec},
		&Job{"https://www.huya.com/g/nz", globalSpec},
		&Job{"https://www.huya.com/g/cfm", globalSpec},
		&Job{"https://www.huya.com/g/802", globalSpec},
		&Job{"https://www.huya.com/g/m3guo", globalSpec},
		&Job{"https://www.huya.com/g/tyj", globalSpec},
		&Job{"https://www.huya.com/g/yhzr", globalSpec},
		&Job{"https://www.huya.com/g/streetball", globalSpec},
		&Job{"https://www.huya.com/g/yys", globalSpec},
	}
}
