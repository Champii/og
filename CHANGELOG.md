# Changelog

## v0.1.10: Current version
  - ## *Warning*: This version cannot be rebuilt from 0.1.9 as this is introducing a grammar breaking change.
  - Better `Makefile` that cares about whether a file has changed or not before applying rules
  - Proper Anonymous field in structs
    ```go
      struct Foo
        *SomeClass
        a int
    ```
  - Channel Send and Receive
    ```go
      c := make(chan int)
      c <- 1
      x := <-c
    ```
  - Select statement
    ```go
      select
        <-c1      => something()
        x := <-c1 => something()
        c1 <- 1   =>
          something()
          somethingElse()
        _         => defaultCase()
    ```
  - Pointer type for receiver in methods
    ```go
      // inline
      struct Foo
        nonPointer       : int -> return 1
        *pointerReceiver : int -> return 1

      // external
      Foo::nonPointer2       : int -> return 1
      Foo::*pointerReceiver2 : int -> return 1
    ```
  - External CHANGELOG
  - Function literal
    ```go
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
    ```go
      fn(a interface): interface{} -> a
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
