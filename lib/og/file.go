package og

import (
	"errors"
	"fmt"
	"github.com/champii/og/lib/ast"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type File struct {
	Path    string
	OutPath string
	Name    string
	Ast     ast.INode
	Output  string
	Source  []byte
}

func (this File) Write() {
	os.MkdirAll(filepath.Dir(this.OutPath), os.ModePerm)
	ioutil.WriteFile(this.OutPath, []byte(this.Output), os.ModePerm)
}
func (this *File) Format() error {
	cmd := exec.Command("gofmt")
	stdin, _ := cmd.StdinPipe()
	stdin.Write([]byte(this.Output))
	stdin.Close()
	final, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("!!! THIS IS A BUG !!!\n")
		fmt.Println("If you see this text, Oglang have generated an invalid Go code")
		fmt.Println("and cannot go through the go formater.")
		fmt.Println("You should report this as an issue along with the file that produced that error")
		fmt.Println("https://github.com/Champii/og/issues")
		return errors.New(this.Path + ": " + string(final))
	}
	this.Output = string(final)
	return nil
}
func NewFile(filePath, outPath string) *File {
	name := path.Base(filePath)
	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	return &File{
		Path:    filePath,
		OutPath: outPath,
		Name:    name,
		Source:  source,
		Output:  "",
	}
}
