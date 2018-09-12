package og

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

type Og struct {
	Config   *OgConfig
	Compiler *OgCompiler
	Printer  *Printer
}

func (this Og) Run() error {
	if this.Config.Print || this.Config.Ast || this.Config.Dirty || this.Config.Blocks {
		this.Config.Quiet = true
	}
	if len(this.Config.Paths) == 0 {
		this.Config.Paths = []string{"."}
	}
	if this.Config.Interpreter {
		RunInterpreter(this.Compiler)
		return nil
	}
	if err := this.Compiler.Compile(); err != nil {
		return err
	}
	if len(this.Compiler.Files) == 0 {
		this.Printer.NothingToDo()
		if !this.Config.Run {
			return nil
		}
	}
	if !this.Config.NoBuild {
		if err := this.Build(); err != nil {
			return err
		}
	}
	if this.Config.Run {
		if err := this.RunBinary(); err != nil {
			return err
		}
	}
	return nil
}
func (this Og) Build() error {
	this.Printer.Compiling(len(this.Compiler.Files))
	cmd := exec.Command("go", "build")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	if len(this.Compiler.Files) > 0 {
		this.Printer.Compiled(len(this.Compiler.Files))
	}
	return nil
}
func (this Og) RunBinary() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	current := path.Base(dir)
	this.Printer.Running()
	cmd := exec.Command("./" + current)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()
	return nil
}
func NewOg(config *OgConfig) *Og {
	printer := NewPrinter(config)
	return &Og{
		Config:   config,
		Compiler: NewOgCompiler(config, printer),
		Printer:  printer,
	}
}
