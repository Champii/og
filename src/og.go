package og

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Compile(path string) string {
	res, err := ioutil.ReadFile(path)

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

	return final
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
