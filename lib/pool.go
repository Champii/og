package og

import (
	tm "github.com/buger/goterm"
	"strconv"
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
	this.Print()
}
func (this Pool) Print() {
	if !this.Verbose {
		return
	}
	tm.Print("                                          \r")
	tm.Println(" ", tm.Color("[", tm.RED), tm.Color(strconv.Itoa(this.Finished), tm.YELLOW), tm.Color("/", tm.RED), tm.Color(strconv.Itoa(this.Total), tm.GREEN), tm.Color("]", tm.RED), "Building sources")
	working := 0
	for i, worker := range this.Workers {
		tm.Print("                                          \r")
		if worker.Processing != "." {
			working++
			tm.Println(tm.Color(strconv.Itoa(i+1), tm.CYAN), tm.Color(":", tm.RED), tm.Color(worker.Processing, tm.MAGENTA))
		}
	}
	for working < len(this.Workers) {
		tm.Print("                                          \n")
		working++
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
