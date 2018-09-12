package og

import (
	"fmt"
	"github.com/champii/og/lib/ast/walker"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type OgCompiler struct {
	Config  *OgConfig
	Parser  *OgParser
	Preproc *OgPreproc
	Printer *Printer
	Files   []*File
}

func (this *OgCompiler) Compile() error {
	for _, p := range this.Config.Paths {
		if err := filepath.Walk(p, this.walker); err != nil {
			fmt.Println("Error", err)
			return err
		}
	}
	if len(this.Files) == 0 {
		if this.Config.Run {
			return nil
		}
	}
	poolSize := this.Config.Workers
	if len(this.Files) < poolSize {
		poolSize = len(this.Files)
	}
	pool := NewPool(poolSize, len(this.Files), this.Printer, this.ParseFile)
	for _, file := range this.Files {
		pool.Queue(file)
	}
	if err := pool.Run(); err != nil {
		return err
	}
	if this.Config.Blocks {
		return nil
	}
	this.desugar()
	return this.outputFiles()
}
func (this *OgCompiler) outputFiles() error {
	for _, file := range this.Files {
		file.Output = file.Ast.Eval()
		if this.Config.Print || this.Config.Dirty || this.Config.Blocks {
			fmt.Println(file.Output)
		} else {
			if !this.Config.Dirty {
				if err := file.Format(); err != nil {
					return err
				}
			}
			file.Write()
		}
	}
	return nil
}
func (this *OgCompiler) desugar() {
	desugar := walker.NewDesugar()
	for _, file := range this.Files {
		desugar.Root = file.Ast
		file.Ast = desugar.Walk(file.Ast)
	}
	desugar.GenerateGenerics()
}
func (this *OgCompiler) walker(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	if path.Ext(filePath) != ".og" {
		return nil
	}
	if !this.Config.Force && !this.mustCompile(filePath, info) {
		return nil
	}
	this.Files = append(this.Files, NewFile(filePath, this.getNewPath(filePath)))
	return nil
}
func (this OgCompiler) mustCompile(filePath string, info os.FileInfo) bool {
	newPath := this.getNewPath(filePath)
	stat, err := os.Stat(newPath)
	if err != nil {
		return true
	}
	if this.Config.Print || this.Config.Ast || this.Config.Dirty || this.Config.Blocks {
		return true
	}
	return info.ModTime().After(stat.ModTime())
}
func (this *OgCompiler) ParseFile(file *File) error {
	this.Preproc.Run(file)
	if this.Config.Blocks {
		fmt.Println(file.Output)
		return nil
	}
	var (
		err error
	)
	if !this.Config.Interpreter {
		err = this.Parser.Parse(file)
	} else {
		err = this.Parser.ParseInterpret(file)
	}
	if err != nil {
		return err
	}
	return nil
}
func (this OgCompiler) getNewPath(filePath string) string {
	if this.Config.OutPath != "./" {
		splited := strings.SplitN(filePath, "/", 2)
		filePath = splited[1]
	}
	return strings.Replace(path.Join(this.Config.OutPath, filePath), ".og", ".go", 1)
}
func NewOgCompiler(config *OgConfig, printer *Printer) *OgCompiler {
	return &OgCompiler{
		Config:  config,
		Printer: printer,
		Parser:  NewOgParser(config),
		Preproc: NewOgPreproc(),
	}
}
