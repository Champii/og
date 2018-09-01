Og-Lang
===

### v0.1.0

Language inspired from [Livescript](http://livescript.net) that compiles to a subset of GoLang.

To be pronounced `Oh-Jee`.

Built with [Antlr4](https://github.com/antlr/antlr4) from their `Golang` grammar.

`og` is written in `og`. See the [src](https://github.com/champii/og/tree/master/src) folder for `og` source, and the [lib](https://github.com/champii/og/tree/master/lib) folder for the compiled one.

# Index

- [Goal](#goal)
- [Exemple](#exemple)
- [Build](#build)
- [Long term goal](#long-term-goal)
- [Todo](#todo)
- [Changelog](#changelog)

# Goal

To provide a usable language that compiles to Golang

The main goal is to simplify the syntax, to borrow some concepts from Livescript and other functional languages, to implement Generics and macro processing, as well as some syntaxic sugar to avoid all the boilerplate Golang force us into.

# Exemple

This is an exemple of how `og` looks like actualy. See the [Exemples](https://github.com/champii/og/tree/master/tests/exemples) folder.

```go
!main

import
  fmt
  strings
  "some/repo"

struct Foo
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

Here is the procedure to regenerate the parser from the grammar if you want to make changes to it.

If you just want to (re)build the binary, you can call `make build` or just `go build`

## Build Antlr

This implementation needs the `TreeVisitor` pattern from `Antlr`. You have to get and build the [https://github.com/wxio/antlr4/tree/go-visitor](https://github.com/wxio/antlr4/tree/go-visitor) into jar and Go runtime, as the official antlr4 repo don't have fully implemented them yet.

You will need `Maven`.

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
```

## Og

```bash
# Get Og
go get -u github.com/champii/og
cd $GO_ROOT/src/github.com/champii/og

# This will generate the grammar,
# Compile the existing sources,
# Regenerate the go sources from og,
# Recompile the new go sources to be sure
# And run the tests.
make

# This will just build the sources (if needed)
make build

# And install
sudo make install

# Simple exemple
og exemples/import.og
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

- [ ] Beautyful and meaningful compile error with source context
- [ ] Efficient multi-path recursive compile
- [ ] External type declaration like Haskell: `myFunc :: string -> Foo -> Bar`
- [ ] OneLiner if/for: `if a => 1`, `for b => b--`
- [ ] Predicat recapture: `if a => that`
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
- [ ] Method receiver pointer type
- [ ] Class-like method declaration (nested into the struct)
- [ ] Pattern matching
- [ ] Import renaming and pattern matching

# Changelog

## Current working tree
  - Add support for simple `struct ID {}` declaration. Still support `type ID struct {}`.
  - Add an alias `class` => `struct`.
  - Add shorthand for `package main` => `!main`.

## 0.1.0
  - Initial release
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
