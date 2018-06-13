package scheduler

type Scheduler interface {
	Start()
	Stop()
	AddJobs([]Job)
}

type Job interface {
	Run()
	Spec() string
}