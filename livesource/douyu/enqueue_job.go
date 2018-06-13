package douyu

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
		&Job{"https://www.douyu.com/directory/game/DOTA2", globalSpec},
		&Job{"https://www.douyu.com/directory/game/LOL", globalSpec},
		&Job{"https://www.douyu.com/directory/game/CSGO", globalSpec},
		&Job{"https://www.douyu.com/directory/game/Overwatch", globalSpec},
		&Job{"https://www.douyu.com/directory/game/SC", globalSpec},
		&Job{"https://www.douyu.com/directory/game/jdqs", globalSpec},
		&Job{"https://www.douyu.com/directory/game/How", globalSpec},
		&Job{"https://www.douyu.com/directory/game/mszb", globalSpec},
		&Job{"https://www.douyu.com/directory/game/wzry", globalSpec},
		&Job{"https://www.douyu.com/directory/game/qqdzz", globalSpec},
		&Job{"https://www.douyu.com/directory/game/HOTS", globalSpec},
		&Job{"https://www.douyu.com/directory/game/qhmy", globalSpec},
		&Job{"https://www.douyu.com/directory/game/FIFAOL4", globalSpec},
		&Job{"https://www.douyu.com/directory/game/CF", globalSpec},
		&Job{"https://www.douyu.com/directory/game/NZ", globalSpec},
		&Job{"https://www.douyu.com/directory/game/CFSY", globalSpec},
		&Job{"https://www.douyu.com/directory/game/SMITE", globalSpec},
		&Job{"https://www.douyu.com/directory/game/COS", globalSpec},
		&Job{"https://www.douyu.com/directory/game/jtlq", globalSpec},
		&Job{"https://www.douyu.com/directory/game/yys", globalSpec},
	}
}
