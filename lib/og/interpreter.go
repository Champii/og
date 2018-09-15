package og

import (
	"bufio"
	"fmt"
	"github.com/champii/og/lib/common"
	"io/ioutil"
	"os"
	"os/exec"
)

func RunInterpreter(compiler *OgCompiler) {
	fmt.Println("-> This is an experimental feature, only small expressions are allowed, no declaration")
	running := true
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for running {
		fmt.Print("> ")
		if !scanner.Scan() {
			return
		}
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}
		file := &common.File{
			Path:   "STDIN",
			Name:   "STDIN",
			Source: []byte(ln),
		}
		err := compiler.ParseFile(file)
		if err != nil {
			fmt.Println(err)
		} else {
			execCode(file.Output)
		}
	}
}
func execCode(str string) {
	skelton := `package main
  import "fmt"
  func main() {
  fmt.Print(` + str[:len(str)-1] + `)
  }`
	ioutil.WriteFile("/tmp/main.go", []byte(skelton), os.ModePerm)
	cmd := exec.Command("go", "run", "/tmp/main.go")
	stdin, _ := cmd.StdinPipe()
	stdin.Write([]byte(str))
	stdin.Close()
	final, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(final))
}
