package og

import (
	"github.com/champii/og/lib/common"
	"time"
)

type WorkerCallback (func(*common.File) error)
type Worker struct {
	In         chan *common.File
	Out        chan error
	Processing *common.File
	Job        WorkerCallback
}

func (this *Worker) Run() {
	for todo := range this.In {
		this.Processing = todo
		res := this.Job(todo)
		this.Processing = nil
		this.Out <- res
	}
}
func NewWorker(i chan *common.File, out chan error, job WorkerCallback) *Worker {
	return &Worker{
		Processing: nil,
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
	In       chan *common.File
	Out      chan error
}

func (this *Pool) Queue(job *common.File) {
	this.In <- job
}
func (this *Pool) Run() error {
	for _, worker := range this.Workers {
		go worker.Run()
	}
	close(this.In)
	common.Print.CursorHide()
	defer common.Print.CursorShow()
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
		if worker.Processing != nil {
			files = append(files, worker.Processing.Path)
			workerIds = append(workerIds, i+1)
		}
	}
	common.Print.CompileList(files, workerIds, len(this.Workers), this.Finished, this.Total)
}
func NewPool(size int, nbJobs int, cb WorkerCallback) *Pool {
	pool := &Pool{
		Size:     size,
		Total:    nbJobs,
		Finished: 0,
		In:       make(chan *common.File, nbJobs),
		Out:      make(chan error, nbJobs),
	}
	for i := 0; i < pool.Size; i++ {
		pool.Workers = append(pool.Workers, NewWorker(pool.In, pool.Out, cb))
	}
	return pool
}
