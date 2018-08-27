# Og
Language that compiles to a subset of GoLang

## Goal

To provide a usable language that compiles to Golang

The main goal is to simplify the syntax, to borrow some concepts from Livescript, to implement Generics and macro processing, as well as some syntaxic sugar to avoid all the boilerplate Golang force us into.

## Exemple

This is an exemple of how Og looks like actualy

```go
package main

import
  "fmt"

struct Foo
  bar string

myFunc (a string) -> string
  return a

main ->
  test = "foo"
  toto = []string
    "toto"
    "tata"

  if test is "foo"
    fmt.Println(test)
  else
    fmt.Println("Not foo")

  for _, v in toto
    fmt.Println(v)

  fmt.Println(toto[1])

  fmt.Println(myFunc("a"))

```

## Long term goal

```go
package main

import "fmt"

struct Foo
  num string

struct Generic<T>
  test T
  fn ->
    fmt.Println(this.test)

genericFunc<T>(g: Generic<T>) T ->
  g.test

main ->
  foo = Foo
    bar: "bar"

  t := Generic
    test: "str"

  fmt.Println(t.fn())
```

## Todo

- Multiple return values in assign
- Generics
- Auto return for last statement in a block
- Returnable and assignable statements (if, for, ...)
- For with a range (for i in [0..10])
- For with a custom variable (for i = 0; i < 10; i++) or (for i < 10)
- Increment/Decrement
- Existance test (if toto?) for non-nil value test
- Error bubbling
- External classic Method declaration
- Class-like method declaration (nested into the struct)
- Interfaces
- One liner functions
- Variable shadowing ? Or error ?
- Tests