package main

import (
	"fmt"
	"io/ioutil"
	og "og/lib"

	"testing"
)

func TestMain(*testing.T) {
	expected := []string{
		// package.og
		`package main
`,
		// package_short.og
		`package main
`,
		// import.og
		`package main

import (
	"fmt"
	"strings"
)
`,
		// struct.og
		`package main

type Foo struct {
}
type Foo struct {
}
type Foo struct {
}
type Foo struct {
}
type Foo struct {
	bar string
}
type Foo struct {
	foo int
	bar string
}
type FooTag struct {
	bar    string ` + "`json: \"test\"`" + `
	foobar int    ` + "`sql: \"-\"`" + `
}
`,
		// top_fn.og
		`package main

func oneLiner() { fmt.Println("a") }
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
func testRet2(str string, str2 string) (string, string) {
	return str, str2
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
	if a != a {
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
		// slices.og
		`package main

func main() {
	a := []string{
		"a",
		"b",
		"c",
	}
	b := a[1:2]
	b := a[1:]
	b := a[:]
	b := a[:3]
	b := a[1:2:3]
	b := a[:2:3]
}
`,
		// interface.og
		`package main

type Foo interface{}
type Foo2 interface{}
type Foo4 interface {
	Foo(a int) string
}
type Foo5 interface {
	SomeType
	Foo()
	Bar(a int) string
	Other() interface{}
}
`,
		// switch.og
		`package main

func main() {
	switch a {
	case 2:
		1
	}
	switch b {
	case "x":
		b
	default:
		other
	}
	switch t.(type) {
	case string:
		"string"
	default:
		"other"
	}
}
`,
		// this.og
		`package main

type Test struct {
	i int
}

func (this Test) f() Test { return this }
func (this *Test) g() int { return this.i }
`,
		// internal_method.og
		`package main

type Foo struct {
	a int
}

func (this Foo) f()  { fmt.Println(this.a) }
func (this *Foo) g() { fmt.Println(this.a) }
func main() {
	foo := Foo{}
	foo.f()
}
`,
		// const.og
		`package main

const foo = "bar"
const foo string = "bar"
`,
		// var.og
		`package main

var foo string
var foo = "bar"
var foo string = "bar"
`,
		// map.og
		`package main

func fn(m map[string]string) map[string]string { return m }
func main() {
	m := map[string]string{}
}
`,
		// chan.og
		`package main

func fn(c1 chan<- int, c2 <-chan int) chan int { 0 }
func main() {
	c := make(chan string)
	c := make(chan<- string)
	c := make(<-chan string)
}
`,
		// loop_flow.og
		`package main

func main() {
here:
	toto()

	for i := range foo {
		continue
		continue label
		break
		break label
		fallthrough
		goto here
	}
}
`,
		// anonymous.og
		`package main

type Foo struct {
	*SomeClass
}
type Foo struct {
	SomeClass
}
type Foo struct {
	SomeClass
	*SomeClass2
}
`,
		// select.og
		`package main

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	select {
	case <-c1:
		fmt.Println("Yeah")
	case x := <-c2:
		fmt.Println("Yeah", x)
	case x := <-c2:
		fmt.Println("Yeah", x)
		fmt.Println("Yeah2", x)
	case c2 <- 1:
		fmt.Println("Yeah", x)
	default:
		fmt.Println("Default")
	}
}
`,
		// func_literal.og
		`package main

func main() {
	a := func() { return 1 }
	b := func(a int) { fmt.Println(a) }
	c := func(a int) int { return a }
}
`,
	}

	paths := []string{
		`package`,
		`package_short`,
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
		`slices`,
		`interface`,
		`switch`,
		`this`,
		`internal_method`,
		`const`,
		`var`,
		`map`,
		`chan`,
		`loop_flow`,
		`anonymous`,
		`select`,
		`func_literal`,
	}

	for i, p := range paths {
		data, err := ioutil.ReadFile(fmt.Sprint("./exemples/", p, ".og"))

		if err != nil {
			panic(err)
		}

		res := og.ProcessFile(fmt.Sprint("./exemples/", p, ".og"), string(data))

		if res != expected[i] {
			panic(fmt.Sprint("Error: ", p, "\nGot: \n---\n", res, "\n---\nExpected: \n---\n", expected[i], "\n---\n"))
		}
	}
}
