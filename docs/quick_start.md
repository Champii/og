# Quick Start

## Install
---

You just have to `go get` the repo

```bash
go get -u github.com/champii/og
```

If your `$PATH` includes `$GOPATH/bin` (and it should), you can print the version

```bash
og -v
```

## Hello world
---

![Hello](https://github.com/Champii/og/raw/master/docs/_media/hello_preview.gif)

Create a new folder and a new file

```bash
mkdir hello
cd hello
nano hello.og
```

Type in this little hello world

```og
!main

import fmt

main -> fmt.Println("Hello world !")
```

Compile and execute

```bash
og -r
```

You should get an output like this one:

```
~> Oglang: Compiled 1 files
~> Oglang: Running...
Hello world !
```

A new file `hello.go` should apprear in your folder and it should look like this

```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello world !")
}
```

## Full Exemple


```og
!main

import
  fmt
  strings
  "some/repo"

struct Foo
  bar int
  getBar: int    -> @bar
  *setBar(i int) -> @bar = i

Foo::inc    (foo int): int -> @bar + foo
Foo::*setInc(foo int)      -> @bar = @bar + foo

interface Bar
  Foo()
  Bar(i int): SomeType

otherFunc(a string): string -> a

autoIfReturn: int ->
  if true => 1
  else    => 0

genericFunction<T>(arg T): T -> arg

main ->
  test := Foo{}
  test.inc(42)

  genericFunction<int>(42)

  var a int = 
    if true => 1
    else    => 0

  someArr := []string
    "value1"
    "value2"

  for i, v in someArray
    fmt.Println(i, v)

  switch test.getBar()
    42 => doSomething()
    _  => doDefault()

  callback := fn (arg int): int -> arg + 1

  go someFunc()
  go -> doSomething()
```

This example, while producing a valid Go syntax, calls some undefined function and serves only as a showcase of the features of the language.  
If you want to compile it to Golang anyway, you must use the `og -n` option that does not build the go files after compilation.

Go to the [Build](/build.md) section for more informations about the options
