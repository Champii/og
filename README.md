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

## Objective

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

  if foo.bar is "bar"
    fmt.Println("Yeah")

  t := Generic
    test: "str"

  fmt.Println(t.fn())
```
