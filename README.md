# ourcrawler

## Usage

### implement the interface Processor

```go
type Processor interface {
	Process(response *downloader.Response) (collector.ResultItems, error)
}
```

### define you own timer jobs to push task into the task queue

```go
package douyu

func NewTask(url string) *scheduler.Task {
	request, _ := http.NewRequest("GET", url, nil)
	return &scheduler.Task{
		Request:   request,
		Processor: &Processor{},
		Collector: collector.NewFileCollector("/tmp/live_data/douyu"),
	}
}

// timer job
type Job struct {
	url string
	spec string
}

func (j *Job) Run() {
	task := NewTask(j.url)
	crawler.EnqueueTask(task)
}

// crontab spec like "0 1 * * * ?"
func (j *Job) Spec() string {
	return j.spec
}

func NewJobs() []scheduler.Job {
	return []scheduler.Job {
		&Job{"https://www.douyu.com/directory/game/DOTA2", globalSpec},
	}
}
```

### add timer jobs into scheduler

```go
// initializer/init_jobs,go
func InitSchedulerJobs(s scheduler.Scheduler) {
	s.AddJobs(douyu.NewJobs())
}
```

### run crawler server

```
// build
go build launcher.go

// run
./launcher
```
