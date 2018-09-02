Og-Lang
===
### *v0.1.7*

<table>
  <tr><td><b>"Golang On Steroids"</b></td>         <td>- <em>Socrates</em></td></tr>
  <tr><td><b>"The Code like it should be. 5/7"</b></td><td>- <em>Mahatma Gandhi</em></td></tr>
  <tr><td><b>"...[Recursive Facepalm]..."</b></td> <td>- <em>Google</em></td></tr>
</table>

# Index

1. [Intro](#intro)
1. [Install](#install)
1. [Basics](#basics)
1. [Usage](#usage)
1. [Build](#build)
1. [Changelog](#changelog)
1. [Todo](#todo)
1. [Long term utopia](#long-term-utopia)

# Intro

To be pronounced `Oh-Jee`.

`Oglang` is an indentation based languagem mainly inspired from [Livescript](http://livescript.net) that compiles to a subset of `GoLang` but aim to be a superset and to be totally backward compatible with it. The goal is for every `Go` file to be a valid `Og` file

### Bootstraped Language

`Oglang` is written in itself. It is said to be a 'Bootstraped' language. In fact, `Oglang` needs the previous release of itself to build itself.

See the [Src](https://github.com/champii/og/tree/master/src) folder for `Oglang` source.

And the [Lib](https://github.com/champii/og/tree/master/lib) folder for the compiled one (in Golang).

Built with [Antlr4](https://github.com/antlr/antlr4) from their `Golang` grammar.

### Goal

The main goal is to simplify the syntax, to borrow some concepts from Livescript and other functional languages, to implement Generics and macro processing, as well as some syntaxic sugar to avoid all the boilerplate code Golang forces us into.

# Install

```bash
# You just have to `go get` the repo
go get -u github.com/champii/og

# If your `$PATH` includes `$GOPATH/bin` (and it should)
og --version
```

# Basics

This is an exemple of how `Oglang` looks like actualy. See the [Exemples](https://github.com/champii/og/tree/master/tests/exemples) folder or the [Src](https://github.com/champii/og/tree/master/src) folder.

```go
!main // package shorthand

// No parenthesis and single word can omit quotes
import
  fmt
  strings
  "some/repo"

// Alias for `type Foo struct`
struct Foo
  // Classical property declaration
  bar int

  // Inline method declaration
  // with no arguments that returns `this.bar`
  getBar : int -> return @bar

// External Foo method declaration
Foo::inc(foo int) -> @bar = @bar + foo

// Alias for `type Bar interface`
interface Bar
  Foo()
  Bar(i int): SomeType

// Alias for struct
class Test

// Classical top-level function declaration
otherFunc(a string): string -> return a

main ->
  test := Foo{}

  test.inc(42)

  if test.getBar() == 42
    answerTheUltimateQuestion()
  else
    enterTheVoid()

  someArr := []string
    "value1"
    "value2"

  for i, v in someArray
    fmt.Println(i, v)

  switch test.getBar()
    42 => doSomething()
    _  => doDefault()

  go someFunc()

  // Auto executed closure when in goroutines
  // No need to add the extra `()`
  go -> doSomething()
```

# Usage

```
NAME:
  Oglang - Golang on steroids

USAGE:
  og [options] Folders|Files

VERSION:
  v0.1.7

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

If you just want to (re)build the binary, you can call `make build` or just `go build` (needs a previously generated parser from grammar. See below)

You will need `Java`, the Antlr4 library is in `./parser/antlr4-4.7.1-SNAPSHOT-complete.jar`

```bash
# Get Og
go get -u github.com/champii/og
cd $GOPATH/src/github.com/champii/og

# This needs java (see above)
# As it will generate the grammar,
# Compile the existing sources,
# Regenerate the go sources from og,
# Recompile the new go sources to be sure
# And run the tests.
make

# This will just build the sources (if needed)
make build

# This needs a previously built version of `og` binary at local scope.
# Recompiles Og to Go + build + test
make bootstrap

# Needs the last official `og` binary version at global scope.
# It recompiles og from the previous global version
# Handy in case of mistake in the source or for release
make rebootstrap

# Simple exemple
og exemples/import.og
```

# Changelog

## v0.1.7: Current version
  - Add custom antlr4 directly into sources. No need to build it yourself

## v0.1.6
  - Release system for develop

## v0.1.5
  - Forced `Go` syntax highlight on `Og` files for `Github`
  - Rework translator to adapt to new method syntax and `@` alias
  - No need to specify the `{}` in `interface{}` types
    ```go
      fn(a interface): interface -> a
    ```

## v0.1.4
  - Class-like method declaration (nested into the struct)
    ```go
      struct Foo
        bar int
        f : int -> return @bar
    ```

## v0.1.3
  - Slice manipulation
      ```go
      someArr[1:x]
      ```
  - Interfaces
      ```go
      interface Foo
        Fn(a ArgType): ResType
      ```
  - Alias `@` => `this`
    ```go
      Foo::bar : SomeType -> return @someProp
    ```
## v0.1.2
  - Support for simple `struct ID {}` declaration. Still support `type ID struct {}`.
    ```go
      // Equivalent
      struct Foo {}
      type Foo struct {}
    ```
  - Alias `class` => `struct`.
    ```go
      // Equivalent
      struct Foo {}
      class Foo {}
    ```
  - Allow empty `class` declaration
    ```go
      struct Foo
    ```
  - Shorthand for `package main` => `!main`.
    ```go
      // Equivalent
      !main
      package main
    ```
  - Cli support with options and recursive directory walking
    ```bash
      og -o lib src
    ```
  - External method declaration
    ```go
      struct Foo
      Foo::bar -> doSomething()
    ```

## v0.1.0
  - Initial release


# TODO

## Golang superset goal
- [ ] `const` declaration
- [ ] `map` Type
- [ ] `chan` Type
- [ ] Labels
- [ ] `break`
- [ ] `goto`
- [ ] `continue`
- [ ] `fallthrough`
- [ ] `select`
- [ ] For with a custom variable (for i = 0; i < 10; i++)
- [ ] Method receiver pointer type
- [ ] Function literal (assignable)
- [ ] Rest params `...`
- [ ] Recever type in methodExpr
- [ ] Send statement `a <- b`
- [ ] Binary operator (`<<`, `>>`, `.`, `|`)
- [ ] Import renaming and pattern matching
- [ ] Proper Anonymous field in structs
- [ ] Incrementation everywhere
- [ ] Make tests truly executable
- [ ] Beautyful and meaningful compile error with source context
- [ ] VSCode extension
- [ ] Adapt Golang tooling like `gofmt` or `golint`

## Syntaxic Sugar Goal
- [ ] Empty statement (need to avoid skiping empty lines in preproc)
- [ ] Empty Function body
- [ ] OneLiner if/for: `if a => 1`, `for b => b--`
- [ ] Returnable and assignable statements (if, for, ...)
- [ ] Auto return for last statement in a block
- [ ] Predicat recapture: `if a => that`
- [ ] External type declaration like Haskell: `myFunc :: string -> Foo -> Bar`
- [ ] Struct compostion ("Inheritance")
- [ ] Existance test (if toto?) for non-nil value test
- [ ] Auto add `default` to switch ?
- [ ] `pub` visibility instead of capitalizing
- [ ] For with a range (for i in [0..10])
- [ ] Error bubbling
- [ ] Pattern matching
- [ ] Function currying
- [ ] Function shorthand `(+ 10)`
- [ ] Perfs
- [ ] Generics

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