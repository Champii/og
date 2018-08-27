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

main ->
  test = "lol"
  fmt.Println(test)
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
