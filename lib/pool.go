package og

import (
	"fmt"
	tm "github.com/buger/goterm"
	curs "github.com/k0kubun/go-ansi"
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
	spinner  int
}

func (this Pool) Queue(job string) {
	this.In <- job
}
func (this *Pool) Run() {
	if this.Verbose {
		curs.CursorHide()
		defer curs.CursorShow()
	}
	for _, worker := range this.Workers {
		go worker.Run()
	}
	close(this.In)
	ticker := time.NewTicker(200 * time.Millisecond)
	go func() {
		for true {
			<-ticker.C
			this.spinner += 1
			this.spinner %= 4
			this.Print()
		}
	}()
	for this.Finished < this.Total {
		if err := <-this.Out; err != nil {
			fmt.Println(err)
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
	spinners := []string{
		"|",
		"/",
		"-",
		"\\",
	}
	tm.Print("                                          \r")
	tm.Println(spinners[this.spinner], tm.Color("[", tm.RED), tm.Color(strconv.Itoa(this.Finished), tm.YELLOW), tm.Color("/", tm.RED), tm.Color(strconv.Itoa(this.Total), tm.GREEN), tm.Color("]", tm.RED), tm.Color("Building sources", tm.YELLOW))
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
