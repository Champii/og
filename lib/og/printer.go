package og

import (
	tm "github.com/buger/goterm"
	curs "github.com/k0kubun/go-ansi"
	"strconv"
)

type Printer struct {
	Config  *OgConfig
	spinner int
}

func (this Printer) Compiling(nbFiles int) {
	if this.Config.Quiet {
		return
	}
	tm.Print("                                          \r")
	tm.Println(tm.Color("[", tm.RED), tm.Color(strconv.Itoa(nbFiles), tm.GREEN), tm.Color("/", tm.RED), tm.Color(strconv.Itoa(nbFiles), tm.GREEN), tm.Color("]", tm.RED), tm.Color("Compiling go", tm.YELLOW))
	tm.MoveCursorUp(1)
	tm.Flush()
}
func (this Printer) Compiled(nbFiles int) {
	if this.Config.Quiet {
		return
	}
	tm.MoveCursorUp(1)
	tm.Print("                                          \r")
	tm.Print(tm.Color("~> ", tm.RED), tm.Color("Oglang: ", tm.MAGENTA), tm.Color("Compiled ", tm.GREEN), tm.Color(strconv.Itoa(nbFiles), tm.YELLOW), tm.Color(" files.", tm.GREEN))
	tm.Flush()
}
func (this Printer) Running() {
	if this.Config.Quiet {
		return
	}
	tm.MoveCursorUp(1)
	tm.Print("                                          \r")
	tm.Print(tm.Color("~> ", tm.RED), tm.Color("Oglang: ", tm.MAGENTA), tm.Color("Running... \n", tm.GREEN))
	tm.Flush()
}
func (this Printer) NothingToDo() {
	tm.Print("                                          \r")
	tm.Print(tm.Color("~> ", tm.RED), tm.Color("Oglang: ", tm.MAGENTA), tm.Color("Nothing to do.", tm.GREEN))
	tm.Flush()
}
func (this *Printer) CompileList(files []string, workerIds []int, nbWorkers, finished, total int) {
	if this.Config.Quiet {
		return
	}
	spinners := []string{
		"|",
		"/",
		"-",
		"\\",
	}
	tm.Print("                                          \r")
	tm.Println(spinners[this.spinner], tm.Color("[", tm.RED), tm.Color(strconv.Itoa(finished), tm.YELLOW), tm.Color("/", tm.RED), tm.Color(strconv.Itoa(total), tm.GREEN), tm.Color("]", tm.RED), tm.Color("Building sources", tm.YELLOW))
	for i, file := range files {
		tm.Print("                                          \r")
		tm.Println(tm.Color(strconv.Itoa(workerIds[i]), tm.CYAN), tm.Color(":", tm.RED), tm.Color(file, tm.MAGENTA))
	}
	workers := nbWorkers
	for len(files) < workers {
		tm.Print("                                          \n")
		workers--
	}
	tm.MoveCursorUp(nbWorkers + 2)
	tm.Flush()
	this.spinner += 1
	this.spinner %= 4
}
func (this Printer) CursorHide() {
	if !this.Config.Quiet {
		curs.CursorHide()
	}
}
func (this Printer) CursorShow() {
	if !this.Config.Quiet {
		curs.CursorShow()
	}
}
func NewPrinter(config *OgConfig) *Printer {
	return &Printer{Config: config}
}
