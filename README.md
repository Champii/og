Og-Lang v0.1.3
===

#### `Golang On Steroids` - _Socrates_.
#### `The world like it should be` - _Mahatma Gandhi_.
#### `...[Facepalm]...` - _Google_.

# Intro
`Oglang` is an indentation based languagem mainly inspired from [Livescript](http://livescript.net) that compiles to a subset of GoLang (for the moment :D).

To be pronounced `Oh-Jee`.

`Oglang` is written in itself. It is said to be a 'Bootstraped' language. In fact, `Oglang` needs the previous release of itself to build itself.

See the [src](https://github.com/champii/og/tree/master/src) folder for `Oglang` source.

And the [lib](https://github.com/champii/og/tree/master/lib) folder for the compiled one (in Golang).

Built with [Antlr4](https://github.com/antlr/antlr4) from their `Golang` grammar.

# Index

- [Goal](#goal)
- [Exemple](#exemple)
- [Usage](#usage)
- [Build](#build)
- [Todo](#todo)
- [Changelog](#changelog)
- [Long term utopia](#long-term-utopia)

# Goal

To provide a superset of `Golang` to be a usable language that compiles to Golang

The main goal is to simplify the syntax, to borrow some concepts from Livescript and other functional languages, to implement Generics and macro processing, as well as some syntaxic sugar to avoid all the boilerplate Golang force us into.

# Exemple

This is an exemple of how `Oglang` looks like actualy. See the [Exemples](https://github.com/champii/og/tree/master/tests/exemples) folder.

```go
!main

import
  fmt
  strings
  "some/repo"

struct Foo
  bar int

Foo::myFunc(foo int) : int -> return @bar + foo

otherFunc(a string): string -> return a

main ->
  test := "foo"

  if test == "foo"
    fmt.Println(test)
  else
    fmt.Println("Not foo")

  for _, v in someArray
    fmt.Println(v)

  switch test
    "foo" => doSomething()
    _     => doDefault()
```
# Usage

```
NAME:
  Oglang - Golang on steroids

USAGE:
  og [options] Folders|Files

VERSION:
  0.1.3

OPTIONS:
  -o value, --out value  Output directory. If input is recursive folder, the tree is recreated (default: "./")
  -p, --print            Print only to stdout. No files created
  -d, --dirty            Don't use 'go fmt'
  -b, --blocks           Get only the generated blocks from indent. No compilation to go.
  -v, --verbose          Show the filenames
  -h, --help             Print help
  -V, --version          Print version
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

# This needs Antlr4 (see above)
# As it will generate the grammar,
# Compile the existing sources,
# Regenerate the go sources from og,
# Recompile the new go sources to be sure
# And run the tests.
make

# This will just build the sources (if needed)
make build

# This needs the previous version of `og` binary at global scope.
# To Recompile Og to Go + build + test
make bootstrap

# Needs the previous version of `og` binary at global scope.
# It recompiles og from the previous global version
# Handy in case of mistake in the source
make rebootstrap

# And install
sudo make install

# Simple exemple
og exemples/import.og
```

# TODO

- [ ] Make `Oglang` a valid super-set of `Golang`.
- [ ] Make tests truly executable
- [ ] Beautyful and meaningful compile error with source context
- [ ] External type declaration like Haskell: `myFunc :: string -> Foo -> Bar`
- [ ] OneLiner if/for: `if a => 1`, `for b => b--`
- [ ] Predicat recapture: `if a => that`
- [ ] Perfs
- [ ] Binary operator (`<<`, `>>`, `.`, `|`)
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
- [ ] `map` Type
- [ ] `chan` Type
- [ ] `const` declaration
- [ ] Send statement `a <- b`
- [ ] Incrementation everywhere
- [ ] Send statement `a <- b`
- [ ] Empty statement (need to avoid skiping empty lines in preproc)
- [ ] Labels
- [ ] `break`
- [ ] `goto`
- [ ] `continue`
- [ ] `fallthrough`
- [ ] `select`
- [ ] Rest params `...`
- [ ] Recever type in methodExpr
- [ ] Function literal (assignable)
- [ ] Proper Anonymous field in structs
- [ ] Method receiver pointer type
- [ ] Auto add `default` to switch ?
- [ ] Class-like method declaration (nested into the struct)
- [ ] Pattern matching
- [ ] Function currying
- [ ] Function shorthand `(+ 10)`
- [ ] Import renaming and pattern matching
- [ ] Adapt Golang tooling like `gofmt` or `golint`
- [ ] VSCode extension

# Changelog

## Current working tree
  - Slice manipulation
  - Interfaces
  - Alias `@` => `this`

## 0.1.2
  - Support for simple `struct ID {}` declaration. Still support `type ID struct {}`.
  - Alias `class` => `struct`.
  - Allow empty `class` declaration
  - Shorthand for `package main` => `!main`.
  - Cli support with options and recursive directory walking

## 0.1.0
  - Initial release

# Long term utopia

```go
!main

import fmt

// Generics
struct Generic<T>
  // attributes
  pub test T

  // Instance method
  fn : T -> @test

  // Class method with external type declaration
  $ T -> Generic<T>
  pub @new(v) -> Generic<T>{ v }


// External grouped method definition
Generic<T>::(
  // With '*', tells the receiver (this or @) is a pointer to Generic<T>
  *method1 : int -> 1

  // External type declaration with no return
  $ int -> SomeComplexType -> (AnotherOne -> interface{}) -> ()
  method2(a, b, c) -> 2
)

// Returnable statements, nil existence test, predicat recapture (that)
genericFunc<T>(g Generic<T>): T -> if g.fn()? => that
                                   else       => somethingElse

// External type declaration
$ Generic<T> -> T
genericFunc<T>(g) -> g.test

// Multiple return values
$ int -> (int, string)
multipleReturn(i) -> i, "foo"

// Automatic "it" argument when not specified
$ string -> string
someFunc -> it + "Hello"

// No arguments, multiple return values, error bubbling
# (error, SomeType)
failableFunc ->
  res1 := funcsThatMayFail()?
  res2 := funcsThatFail()?

  res1 + res2

// Macro definition (like rust)
$macro my_macro ->
  ($number:expr) =>
    myFunc$(number) : int -> $number
    if $number > 0
      $my_macro($number - 1)

// Generate 10 function `myFunc10(), myFunc9(), .. myFunc0()` that all return their number,
$my_macro(10)

// Operator definition with precedence
operator ~~ 9

// Typeclass (could be implemented with macros ? lets see...)
impl SomeTypeClass for MyType
  x ~~ y = x.DoSomething(y)

main ->
  t := Generic<string>::new("str")

  // Range array creation, call chaining,
  // function currying and function shorthand.
  // Here a == [10, 11, 12, 13, 14, 15]
  a := []int{0..10}
    |> map((+ 10))
    |> filter((<= 15))

  // Function composition
  f := map >> filter
```