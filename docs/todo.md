# Todo
---

## Golang basics
---

- [ ] Named return

## Syntaxic Sugar
---

- [ ] Slices declaration without type `[1, 2, 3]`, `[]string` (need type inference)
- [ ] `*` Operator in slices to reference own lenght `arr = arr[*-1]`
- [ ] Suffix keyword `return foo if bar`, `foo += i for i in array`
- [ ] Returnable and assignable statements (for, switch, ...)
- [ ] Predicat recapture: `if a => that`
- [ ] External type declaration like Haskell: `myFunc :: string -> Foo -> Bar`
- [ ] Existance test (if foo? => bar) for non-nil value test
- [ ] `pub` visibility instead of capitalizing
- [ ] For with a range (for i in [0..10])
- [ ] Pattern matching
- [ ] Auto setup package name with folder name if not specified
- [ ] Error bubbling
- [ ] Function currying
- [ ] Function shorthand `(+ 10)`, `map(structArr, (.SomeField))` 
- [ ] Import pattern matching
- [ ] Remove that `fn` keywork that diminish lisibility
- [ ] Conditionnal expression like `res := getFirst() || getSecond()` that make `if` statements
- [ ] Assignation and return for `for`, `switch`
- [ ] `super` keyword
- [ ] Default arguments `f := (i int = 2, s string = "foo") ->`
- [ ] Extended arguments to set to this `Foo::bar(@someVar) -> `
- [ ] Global type inference and auto-generic generation to access full parametric polymorphism

## Technical
---

- [ ] Cannot call template funcs from another package
- [ ] Make a distinction between Generics(run time) and Templates(compile time)
- [ ] Early checks like import validity or undefined references based on first pass analysis of the project
- [ ] Recursive Generics
- [ ] Fix: Cannot have a function literal other than last argument in function (actual workaround is to define a `type` and to pass it instead)
- [ ] Perfs !! (More specific rules, reduce size and workload of Walkers, remove ambiguity in grammar)
- [ ] Do a single pass on AST instead of multiple walkers (for perfs)
- [ ] Fix bad perfs for nested struct instantiation and if/else block
- [ ] Simple type checker to catch errors before Golang formater/compiler does
- [ ] Language extensions ?
- [ ] Make tests truly executable
- [ ] VSCode extension
- [ ] Adapt Golang tooling like `gofmt` or `golint`
- [ ] Allow to give arguments to finaly executed binary with `og -r -- arg1 arg2`
- [ ] Fix typeswitch that cannot allow a pointer as second case without `;`
    ```go
      switch template.Node.(type)
        *FunctionDecl => @GenerateTopFns(template); // <- that `;` is needed here
        *StructType => @GenerateStruct(template)
    ```
    