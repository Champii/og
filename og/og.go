package og

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Compile(path string) string {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("error reading file", path, err)
	}

	preprocessed := Preproc(string(file))

	// fmt.Println(preprocessed)

	res := Parse(string(preprocessed))

	fmt.Println("OUAT", res)

	final := format(res)

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
