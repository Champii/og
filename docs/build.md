# Build

Here is the procedure to regenerate the parser from the grammar if you want to make changes to it.

If you just want to (re)build the binary, you can call `make build` or just `go build` (needs a previously generated parser from grammar. See below)

You will need `Java`, the Antlr4 library is in `./parser/antlr4-4.7.1-SNAPSHOT-complete.jar`

## Get Og
---
```bash
go get -u github.com/champii/og
cd $GOPATH/src/github.com/champii/og
```

## Make
---

Simply call
```bash
make
```
Calling `make` will regenerate the grammar,  
Compile the existing sources from the previous Og (`og lib`)  
And run the tests.  
Needs the last official `og` binary version at global scope.  

To recompile the whole project:
```bash
make re
```

This cleans the `lib` folder,  
Then compiles og from the previous global version (`og lib`)  
Then recomiles it from itself (`./og lib`)  
And run the tests  

![make](https://github.com/Champii/og/raw/master/docs/_media/og_preview.png)

### Simple test
```bash
og exemples/import.og
```

### Build time
---
The current build time of the project is around 5s for all sources files with `./og` alone, and around 20s for full rebootstrap with `make re`.  
That bootstraps from old version then rebootstraps from itself, with `go build` and `go test` each time. 
