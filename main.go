package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	res, err := ioutil.ReadFile("./exemples/test.og")

	if err != nil {
		fmt.Println("error", err)
	}

	preprocessed := Preproc(string(res))

	// fmt.Println(preprocessed)

	ast, err := Build(string(preprocessed))

	if err != nil {
		fmt.Println("error", err)
	}

	// spew.Dump(ast)

	goSrc := parseAst(ast)

	// fmt.Println(goSrc)

	final := format(goSrc)

	fmt.Println(final)
}

func format(str string) string {
	cmd := exec.Command("gofmt")

	in, _ := cmd.StdinPipe()

	go func() {
		defer in.Close()

		in.Write([]byte(str))
	}()

	final, _ := cmd.CombinedOutput()

	return string(final)
}
