package main

import (
	og "Og/src"
	"fmt"

	"testing"
)

func TestMain(*testing.T) {
	expected := []string{
		// package.og
		`package main
`,
		// import.og
		`package main

import (
	"fmt"
)
`,
		// struct.og
		`package main

type Foo struct {
	bar    string
	foobar int
}
`,
		// top_fn.og
		`package main

func testBody() {
	fmt.Println("a")
}
func testArg(str string) {
	fmt.Println(str)
}
func testRet(str string) string {
	return str
}
`,
	}

	paths := []string{
		`package`,
		`import`,
		`struct`,
		`top_fn`,
	}

	for i, p := range paths {
		res := og.Compile(fmt.Sprint("./exemples/", p, ".og"))

		if res != expected[i] {
			panic(fmt.Sprintln("Error", p, res, expected[i]))
		}
	}
}
