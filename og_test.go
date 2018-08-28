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

type FooTag struct {
	bar    string ` + "`json: \"test\"`" + `
	foobar int    ` + "`sql: \"-\"`" + `
}
`,
		// top_fn.og
		`package main

func oneLiner() {
	fmt.Println("a")
}
func testBody() {
	fmt.Println("a")
}
func testArg(str string) {
	fmt.Println(str)
}
func arrArg(str []string) {
	fmt.Println(str)
}
func doubleArrArg(str [][]string) {
	fmt.Println(str)
}
func doubleArrPointerArg(str **[]*[]*string) {
	fmt.Println(str)
}
func testRet(str string) string {
	return str
}
func testRet2(str string) (string, string) {
	return str, str
}
func testAssign() {
	a := foo()

	a, b := foo()

	a, b, c := foo()

}
`,
		// if.og
		`package main

func main() {
	a := 2

	if a == 2 {
		fmt.Println(a)
	}

	if a == 2 {
		fmt.Println(a)
	} else {
		fmt.Println(a)
	}

	if a == 2 {
		fmt.Println(a)
	} else if a == 3 {
		fmt.Println(a)
	}

	if a == 2 {
		fmt.Println(a)
	} else if a == 3 {
		fmt.Println(a)
	} else if a == 4 {
		fmt.Println(a)
	} else {
		fmt.Println(a)
	}

}
`,
		// for.og
		`package main

import (
	"fmt"
)

func main() {
	a := []string{
		"foo",
		"bar",
	}

	for _, i := range a {
		fmt.Println(i)
	}
}
`,
		// nested_property.og
		`package main

func main() {
	a.b[1][2].c()()(a)[1].d
}
`,
		// goroutine.og
		`package main

func main() {
	go func() {
		fmt.Println("test")
	}()
	go some.fn()
}
`,
		// operation.og
		`package main

func main() {
	a := 1 + 1

	b := 1 + 2/4 - 21%42*84

}
`,
		// ref.og
		`package main

func main() {
	a := &b

	b := **a.b

}
`,
		// inc.og
		`package main

func main() {
	a := 1

	a++
	a--
	a.b.c().d++
}
`,
		// struct_inst.og
		`package main

type Foo struct {
	foo int
	bar string
}

func main() {
	a := Foo{
		foo: 1,
		bar: "lol",
	}

}
`,
	}

	paths := []string{
		`package`,
		`import`,
		`struct`,
		`top_fn`,
		`if`,
		`for`,
		`nested_property`,
		`goroutine`,
		`operation`,
		`ref`,
		`inc`,
		`struct_inst`,
	}

	for i, p := range paths {
		res := og.Compile(fmt.Sprint("./exemples/", p, ".og"))

		if res != expected[i] {
			panic(fmt.Sprint("Error: ", p, "\nGot: \n---\n", res, "\n---\nExpected: \n---\n", expected[i], "\n---\n"))
		}
	}
}
