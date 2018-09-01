Og-Lang
===
Language that compiles to a subset of GoLang

To be pronounced `Oh-Jee`

# Index

- [Goal](#goal)
- [Exemple](#exemple)
- [Build](#build)
- [Long term goal](#long-term-goal)
- [Todo](#todo)

# Goal

To provide a usable language that compiles to Golang

The main goal is to simplify the syntax, to borrow some concepts from Livescript, to implement Generics and macro processing, as well as some syntaxic sugar to avoid all the boilerplate Golang force us into.

# Exemple

This is an exemple of how `og` looks like actualy

```go
package main

import
  fmt
  strings
  "some/repo"

type Foo struct
  bar int

Foo::myFunc(foo int) : int -> return this.bar + foo

myFunc(a string): string -> return a

main ->
  test := "foo"

  if test == "foo"
    fmt.Println(test)
  else
    fmt.Println("Not foo")

  for _, v in someArray
    fmt.Println(v)
```

# Build

To have the `Visitor` pattern in GO, you have to get and build the [https://github.com/wxio/antlr4/tree/go-visitor](https://github.com/wxio/antlr4/tree/go-visitor) into jar and Go runtime, as the official antlr4 repo don't have fully implemented them yet

You will need `Maven`

```bash
# Install maven
sudo apt install maven

# Get the repo
cd $GO_ROOT/src
go get -u github.com/wxio/antlr4
cd github.com/wxio/antlr4

# Switch to go-visitor branch
git checkout go-visitor

# Build the jar
mvn install -DskipTests=true

# Get Og
go get -u github.com/champii/og
cd $GO_ROOT/src/github.com/champii/og

# This will generate the grammar,
# Compile the existing sources,
# Regenerate the go sources from og,
# Recompile the new go sources to be sure
# And run the tests.
make

# Simple exemple
./og exemples/import.og
```
# Long term goal

```go
!main

import fmt

struct Generic<T>
  // attributes
  pub test T

  // Class method
  pub @new(v T) : Generic<T> -> Generic<T>{ v }

  // Instance method
  fn : T -> @test

genericFunc<T>(g Generic<T>): T -> g.test

main ->
  t := Generic<string>::new("str")

  fmt.Println(t.fn())
```

# TODO

- [ ] Perfs
- [ ] Binary operator (`<<`, `>>`, `.`, `|`)
- [ ] Interfaces
- [ ] Empty Function body
- [ ] Struct compostion ("Inheritance")
- [ ] Auto return for last statement in a block
- [ ] For with a range (for i in [0..10])
- [ ] For with a custom variable (for i = 0; i < 10; i++) or (for i < 10)
- [ ] `pub` visibility instead of capitalizing
- [ ] Existance test (if toto?) for non-nil value test
- [ ] Returnable and assignable statements (if, for, ...)
- [ ] Generics
- [ ] Error bubbling
- [ ] Class-like method declaration (nested into the struct)
- [ ] Pattern matching
- [ ] Import renaming and pattern matching

## Done

- [x] Rewrite Og in Og
- [x] Package declaration
- [x] Import
- [x] Structure
- [x] Top level Function
- [x] Function arguments and type
- [x] Function return type
- [x] Return keyword
- [x] Assignation
- [x] Bool
- [x] Int
- [x] Float
- [x] String
- [x] Array
- [x] Nested property (a.b.c)
- [x] Function call
- [x] Array access (a[0])
- [x] If
- [x] Else If
- [x] Else
- [x] Predicate operator
- [x] Is / Isnt alias of `==` / `!=`
- [x] For In
- [x] Goroutine
- [x] `gofmt` to format the output
- [x] One liner functions
- [x] Math Operators
- [x] Array type
- [x] Pointer type
- [x] Struct tags
- [x] Parenthesis
- [x] Reference and Dereference
- [x] Increment/Decrement
- [x] Struct instantiation
- [x] Nil value
- [x] Logical predicat operator (`||` / `&&`  `or` / `and`)
- [x] Multiple return in func
- [x] Multiple return values in assign
- [x] Type assertion
- Tests
  - [x] Package
  - [x] Import
  - [x] Struct
  - [x] Top Level Function
  - [x] If/ElseIf/Else
  - [x] For
  - [x] NestedProperty
  - [x] GoRoutine
  - [x] Math Operation
  - [x] Reference/Dereference
  - [x] Increment/Decrement
  - [x] Struct instantiation
  - [x] Multiple return in func
  - [x] Multiple return values in assign
