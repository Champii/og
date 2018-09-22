package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type File struct {
	Path        string
	Imports     map[string]string
	FullPath    string
	OutPath     string
	Name        string
	Ast         INode
	Output      string
	LineMapping []int
	Source      []byte
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
func (this *File) Error(line, column int, msg, msg2 string) *Error {
	source := strings.Split(string(this.Source), "\n")
	return NewError(this.Path, source, line, column, msg, msg2)
}
func NewFile(filePath, outPath string) *File {
	name := path.Base(filePath)
	fullPath := filePath
	dir, _ := os.Getwd()
	fullPath = path.Join(dir, filePath)
	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	return &File{
		Path:     filePath,
		FullPath: fullPath,
		OutPath:  outPath,
		Name:     name,
		Source:   source,
		Output:   "",
	}
}
