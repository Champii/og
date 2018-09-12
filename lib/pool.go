package og

import (
	curs "github.com/k0kubun/go-ansi"
	"time"
)

type WorkerCallback (func(string) error)
type Worker struct {
	In         chan string
	Out        chan error
	Processing string
	Job        WorkerCallback
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
func NewWorker(i chan string, out chan error, job WorkerCallback) *Worker {
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
	Printer  *Printer
}

func (this Pool) Queue(job string) {
	this.In <- job
}
func (this *Pool) Run() error {
	for _, worker := range this.Workers {
		go worker.Run()
	}
	close(this.In)
	curs.CursorHide()
	defer curs.CursorShow()
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for true {
			<-ticker.C
			this.Print()
		}
	}()
	for this.Finished < this.Total {
		if err := <-this.Out; err != nil {
			ticker.Stop()
			return err
		}
		this.Finished++
	}
	ticker.Stop()
	this.Print()
	return nil
}
func (this Pool) Print() {
	workerIds := []int{}
	files := []string{}
	for i, worker := range this.Workers {
		if worker.Processing != "." {
			files = append(files, worker.Processing)
			workerIds = append(workerIds, i+1)
		}
	}
	this.Printer.CompileList(files, workerIds, len(this.Workers), this.Finished, this.Total)
}
func NewPool(size int, nbJobs int, printer *Printer, cb WorkerCallback) *Pool {
	pool := &Pool{
		Size:     size,
		Total:    nbJobs,
		Finished: 0,
		In:       make(chan string, nbJobs),
		Out:      make(chan error, nbJobs),
		Printer:  printer,
	}
	for i := 0; i < pool.Size; i++ {
		pool.Workers = append(pool.Workers, NewWorker(pool.In, pool.Out, cb))
	}
	return pool
}
