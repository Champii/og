# Changelog

## DEV: Current version

## v0.6.4

## v0.6.3
  - Generics for Struct
    ```og
      struct Foo<T>
        bar T

      main ->
        foo := Foo<int>
          bar: 2
    ```

## v0.6.2
  - Be `--quiet` when either `-p`, `-d`, `-b` or `-a`
  - Force recompile if `-p`, `-d`, `-b` or `-a`
  - Hide and show cursor while compiling
  - Show a spinner
  - Fix `switch res := t.(type)`
  - Go format error is no more swallowed
  - Add template syntax for `struct Foo<T>` but not implemented yet (needed that switch fix)

## v0.6.0
  - Generics that are produced at compile time
    ```og
      genericFunction<T>(arg T): T -> arg

      main ->
        genericFunction<int>(a)
        genericFunction<string>(a)
    ```
  - Cleaner Makefile
  - Awesome cli tool to use multi-threaded compilation 
  - Multithreaded tests
  - Overall project compilation time down to 8s, tests down to 2s.
  - Beautiful progress, made verbose default
  - Removed option `-v, --verbose` to be the opposite: `-q, --quiet`
  - Add option `-w, --workers` to specify a number of thread for compilation
  - Changed `-V, --version` for little `-v, --version`
  - Auto build after compile, add option `-n, --no-build` to avoid that behaviour
  - Add option `-r, --run` to run after compile
  - Overview improvements
  - Allow comments starting with `#` 
  - First type-checker implementation
  - Better `AstWalker`
  - Colors everywhere !


## v0.5.0
  - Removed every non necessary terminal `return`
  - Makefile now compiles with the previous released version by default instead of the last compiled one. New rule `make new` that compiles itself with the last present `og`
  - Fix a bug on ParameterDecl that wanted a name for each type
  - Auto return for `if` statement
    ```og
      someFunc: int ->
        if true => 42
        else    => 0
    ```
  - Assignation for `if` statement
    ```og
      var a int = 
        if true => 42
        else    => 0
    ```
  - Overview section in the docs

## v0.4.2
  - Auto return for function with no return keyword (only for simple statement yet)
    ```og
      someFunc: int -> 2
    ```

## v0.4.1
  - Abstract AST walker and a derived AST Printer
  - `-a` CLI option to show the AST
  - Disambiguation of `for`, `if` and `else if` statements with a `;` before the block
  - Empty statement and empty onliner-function body with `;`
    ```og
      foobar -> ;

      if foo == bar => ;
    ```

## v0.4.0
  - New AST system for more flexibility
  - Switched to SLL prediction mode to greatly improve performances
  - Multiline `type`, `const` and `var` declaration
    ```og
      var (
        foo string
        bar = 2
      )
    ```
  - Disambiguation character `~` for `label` statements
    ```og
      ~myLabel: something()
    ```

## v0.3.0
  - Allow single line If statement
    ```og
      if b == 0 => 0
      else      => 1
    ```
  - Pretty error display (but trivial for some cases)
    ```
      path/file.og (8:6): Unexpected 'foo'
        bar foo
            ^
    ```
  - Allow empty Return statement
  - Small and dirty interpreter that allow only a few expressions

## v0.2.0
  - Variadic arguments
    ```og
      someFunc(a ...int) -> fmt.Println(a)
      main -> someFunc([]int{1, 2, 3}...)
    ```
  - Parenthesis allowed for import declaration
    ```og
      import (
        fmt
        strings
      )
    ```
  - For statement with 3 flavors
    ```og
      for _, i in a
      for i < 10
      for i := 0; i < 10; i++
    ```
  - Bitwise operators `&`, `|`, `^`, `&^`, `<<`, `>>`
  - Import renaming
    ```og
      import
        fmt
        strings
        "some/repo": otherName
        "another/repo": .
    ```

## v0.1.11
  - Function literal but with a mandatory desambiguation symbol `fn`
    ```og
      a := fn              -> fmt.Println(1)
      b := fn (a int)      -> fmt.Println(a)
      c := fn (a int) :int -> return a
      a := fn         :int -> return 1
    ```
  - Function Type with the same mandatory `fn`
    ```og
      a(arg fn: int) -> fmt.Println(arg())
      main          -> a(fn: int -> return 2)
    ```

## v0.1.10
  - Better `Makefile` that cares about whether a file has changed or not before applying rules
  - Proper Anonymous field in structs
    ```og
      struct Foo
        *SomeClass
        a int
    ```
  - Channel Send and Receive
    ```og
      c := make(chan int)
      c <- 1
      x := <-c
    ```
  - Select statement
    ```og
      select
        <-c1      => something()
        x := <-c1 => something()
        c1 <- 1   =>
          something()
          somethingElse()
        _         => defaultCase()
    ```
  - Pointer type for receiver in methods
    ```og
      // inline
      struct Foo
        nonPointer       : int -> return 1
        *pointerReceiver : int -> return 1

      // external
      Foo::nonPointer2       : int -> return 1
      Foo::*pointerReceiver2 : int -> return 1
    ```
  - External CHANGELOG
  - Function literal (Big performance issue caused by the grammar that keeps increasing in complexity, disabled until it's sorted out)
    ```og
      a := -> fmt.Println(1)
      a := (a int) -> fmt.Println(a)
    ```

## v0.1.9
  - `break`, `goto`, `continue`, `fallthrough`, `labels`

## v0.1.8
  - Proper `const` declaration
  - Proper `var` declaration
  - Map declaration
  - Chan type (`chan T`, `<-chan T` and `chan<- T`)
  - Fixed a bug with result type that were not being parsed. We can now use `interface` instread of `interface{}` everywhere

## v0.1.7
  - Add custom antlr4 directly into sources. No need to build it yourself

## v0.1.6
  - Release system for develop

## v0.1.5
  - Forced `Go` syntax highlight on `Og` files for `Github`
  - Rework translator to adapt to new method syntax and `@` alias
  - No need to specify the `{}` in `interface{}` for arguments types (not for return values yet)
    ```og
      fn(a interface): interface{} -> a
    ```

## v0.1.4
  - Class-like method declaration (nested into the struct)
    ```og
      struct Foo
        bar int
        f : int -> return @bar
    ```

## v0.1.3
  - Slice manipulation
      ```og
      someArr[1:x]
      ```
  - Interfaces
      ```og
      interface Foo
        Fn(a ArgType): ResType
      ```
  - Alias `@` => `this`
    ```og
      Foo::bar : SomeType -> return @someProp
    ```
## v0.1.2
  - Support for simple `struct ID {}` declaration. Still support `type ID struct {}`.
    ```og
      // Equivalent
      struct Foo {}
      type Foo struct {}
    ```
  - Alias `class` => `struct`.
    ```og
      // Equivalent
      struct Foo {}
      class Foo {}
    ```
  - Allow empty `class` declaration
    ```og
      struct Foo
    ```
  - Shorthand for `package main` => `!main`.
    ```og
      // Equivalent
      !main
      package main
    ```
  - Cli support with options and recursive directory walking
    ```bash
      og -o lib src
    ```
  - External method declaration
    ```og
      struct Foo
      Foo::bar -> doSomething()
    ```

## v0.1.0
  - Initial release
