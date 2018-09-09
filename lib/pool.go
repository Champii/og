package og

import (
	tm "github.com/buger/goterm"
	"time"
)

type Worker struct {
	In         chan string
	Out        chan error
	Processing string
	Job        func(string) error
}

func (this *Worker) Run() {
	for todo := range this.In {
		this.Processing = todo
		res := this.Job(todo)
		this.Processing = "."
		this.Out <- res
	}
	this.Processing = "."
}
func NewWorker(i chan string, out chan error, job func(string) error) *Worker {
	return &Worker{
		Processing: ".",
		In:         i,
		Out:        out,
		Job:        job,
	}
}

type Pool struct {
	Size     int
	Workers  []*Worker
	Total    int
	Finished int
	In       chan string
	Out      chan error
	Verbose  bool
	lineSize []int
}

func (this Pool) Queue(job string) {
	this.In <- job
}
func (this *Pool) Run() {
	for _, worker := range this.Workers {
		go worker.Run()
	}
	close(this.In)
	ticker := time.NewTicker(200 * time.Millisecond)
	go func() {
		for true {
			<-ticker.C
			this.Print()
		}
	}()
	for this.Finished < this.Total {
		if err := <-this.Out; err != nil {
			tm.Println(err)
		}
		this.Finished++
	}
	ticker.Stop()
}
func (this Pool) Print() {
	if !this.Verbose {
		return
	}
	tm.Print("                                          \r")
	tm.Println("[", this.Finished, "/", this.Total, "]")
	for i, worker := range this.Workers {
		tm.Print("                                          \r")
		tm.Println(i+1, ":", worker.Processing)
	}
	tm.MoveCursorUp(len(this.Workers) + 2)
	tm.Flush()
}
func NewPool(size int, nbJobs int, verbose bool, cb func(string) error) *Pool {
	pool := &Pool{
		Size:     size,
		Total:    nbJobs,
		Finished: 0,
		Verbose:  verbose,
		In:       make(chan string, nbJobs),
		Out:      make(chan error, nbJobs),
	}
	for i := 0; i < pool.Size; i++ {
		pool.Workers = append(pool.Workers, NewWorker(pool.In, pool.Out, cb))
	}
	return pool
}
