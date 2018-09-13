# Introduction
---


'Og' is to be pronounced `Oh-Jee` and stands for ~~`Orgasmic Granny`~~ `Optimistic Golang`  
It is an indentation based language mainly inspired from [Livescript](http://livescript.net) that compiles to GoLang.

```og
!main

class Oglang<T>
  Foo T
  GetFoo: T -> @Foo

main ->
  foo := Oglang<int>
    Foo: 42

  foo.GetFoo()
```

Go to the [Features](/features.md) section for more exemples of code.

### Bootstraped Language
---

Oglang is written in itself. It is said to be a 'Bootstraped' language. In fact, Oglang needs the previous release of itself to build itself.  
See the [Build](/build.md) section for more details

Go to the [./lib](https://github.com/champii/og/tree/master/lib) folder on github for both Oglang and compiled Golang sources.  

Built with [Antlr4](https://github.com/antlr/antlr4) from their Golang grammar.

### Goal
---

The main goal is to simplify the syntax, to borrow some concepts from Livescript and other functional languages, to implement Generics and macro processing, as well as some syntaxic sugar to avoid all the boilerplate code Golang forces us into.
