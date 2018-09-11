Og-Lang *v0.6.2* (Optimistic Golang)
===

### [Documentation]()

# Disclamer

<h3> This software is in its early stage.<br/>
New features come fast, at the cost of breaking things often.<br />
Testers and Contributors are most welcome
</h3>


# Quick overview

![Overview](https://github.com/Champii/og/raw/master/docs/_media/overview_color.png)

```go
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

## Demo

<p float="left">
  <img height="220" src="https://github.com/Champii/og/blob/master/docs/hello_preview.gif" />
  <img height="220" src="https://github.com/Champii/og/blob/master/docs/og_preview.gif" />
</p>


# Build

Here is the procedure to regenerate the parser from the grammar if you want to make changes to it.  
If you just want to (re)build the binary, you can call `make build` or just `go build` (needs a previously generated parser from grammar. See below)  
You will need `Java`, the Antlr4 library is in `./parser/antlr4-4.7.1-SNAPSHOT-complete.jar`

```bash
# Get Og
go get -u github.com/champii/og
cd $GOPATH/src/github.com/champii/og

# This will regenerate the grammar,
# Compile the existing sources from the previous Og (`og lib`)
# And run the tests.
# Needs the last official `og` binary version at global scope.
make

# It cleans the `lib` folder,
# Then compiles og from the previous global version (`og lib`)
# Then recomiles it from itself (`./og lib`)
# And run the tests
make re

# Simple exemple
og exemples/import.og
```

The current build time of the project is around 5s for all sources files with `./og` alone, and around 20s for full rebootstrap with `make re` (That bootstraps from old version then rebootstraps from itself, with `go build` and `go test` each time). 

# Quotes

<table>
  <tr><td><b>"Golang On Steroids"</b></td>         <td>- <em>Socrates</em></td></tr>
  <tr><td><b>"The Code like it should be. 5/7"</b></td><td>- <em>Mahatma Gandhi</em></td></tr>
  <tr><td><b>"(..Recursive Facepalm..)"</b></td> <td>- <em>Google</em></td></tr>
</table>
