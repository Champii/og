package common

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/fatih/color"
	curs "github.com/k0kubun/go-ansi"
	"strconv"
)

var (
	yellow = color.New(color.FgHiYellow).SprintFunc()
)
var (
	red = color.New(color.FgHiRed).SprintFunc()
)
var (
	cyan = color.New(color.FgCyan).SprintFunc()
)
var (
	magenta = color.New(color.Bold, color.FgHiMagenta).SprintFunc()
)
var (
	blue = color.New(color.Bold, color.FgHiBlue).SprintfFunc()
)
var (
	green = color.New(color.FgHiGreen).SprintfFunc()
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
	if this.Config.Quiet {
		return
	}
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
func (this *Printer) Error(err *Error) {
	for i := 0; i < 8; i++ {
		tm.Println("                                                                          ")
	}
	tm.MoveCursorUp(9)
	tm.Flush()
	fileInfo := fmt.Sprintf("%s (%s:%s)", green(err.Path), yellow(err.Line), yellow(err.Column))
	fmt.Printf("\n%s: %s '%s'\n", fileInfo, red(err.Msg), magenta(err.Msg2))
	badLine := err.Source[err.Line-1]
	badLine = cyan(badLine[:err.Column]) + magenta(err.Msg2) + cyan(badLine[err.Column+len(err.Msg2):])
	fmt.Println(badLine)
	fmt.Print(blue("%"+strconv.Itoa(err.Column+1)+"s\n", "^"))
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

var (
	Print *Printer
)
