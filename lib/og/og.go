package og

import (
	"fmt"
	"github.com/champii/og/lib/common"
	"os"
	"os/exec"
	"path"
)

type Og struct {
	Config   *common.OgConfig
	Compiler *OgCompiler
}

func (this Og) Run() error {
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
		common.Print.NothingToDo()
		if !this.Config.Run {
			return nil
		}
	}
	if this.Config.Print || this.Config.Ast || this.Config.SimpleAst || this.Config.Blocks || this.Config.Dirty {
		return nil
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
	common.Print.Compiling(len(this.Compiler.Files))
	cmd := exec.Command("go", "build")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	if len(this.Compiler.Files) > 0 {
		common.Print.Compiled(len(this.Compiler.Files))
	}
	return nil
}
func (this Og) RunBinary() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	current := path.Base(dir)
	common.Print.Running()
	cmd := exec.Command("./" + current)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()
	return nil
}
func NewOg(config *common.OgConfig) *Og {
	common.Print = common.NewPrinter(config)
	return &Og{
		Config:   config,
		Compiler: NewOgCompiler(config),
	}
}
