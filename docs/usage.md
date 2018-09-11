# Usage

## The Binary
---

```
$> og -h
NAME:
  Oglang - Golang on steroids

VERSION:
  v0.6.2

USAGE:
  og [options] [folders...|files...]

  If a Print argument (-p, -b, -d, -a) is given, NO COMPILATION is done.

  If run without files, it will compile and execute '.'

OPTIONS:
  -r, --run                      Run the binary
  -o directory, --out directory  Output destination `directory` (default: "./")
  -w jobs, --workers jobs        Set the number of jobs (default: 8)
  -p, --print                    Print the file
  -d, --dirty                    Print the file before going through 'go fmt'
  -b, --blocks                   Print the file after it goes to preprocessor.
  -a, --ast                      Print the generated AST
  -i, --interpreter              Run a small interpreter (ALPHA)
  -q, --quiet                    Hide the progress output
  -n, --no-build                 Dont run 'go build'
  -h, --help                     Print help
  -v, --version                  Print version
```

## Basic file compilation
---

By default Og recursively compile every `.og` file in the current folder `.` and produce `.go` files that are along their Og source. It will then run `go build` on the folder 
```bash
./og
```

To additionaly run the produces binary, you can add the `-r` flag
```bash
./og -r
```

With just a file name, the compiler will produce a single `.go` file inside the same directory
```bash
./og folder/file.og
```

You can give multiple files and folder that will be walked recursively
```bash
./og file.og folder/ anotherFile.og
```

The output flag `-o` will save the files into another folder. The folder hierarchy is recreated. 
```bash
./og -o lib src/file.og
```

## Debug
---

You can also print the file without affecting the fs with `-p`
```bash
./og -p src/file.og
```

The `-d` (`--dirty`) option shows you the bare generated file from the parser, without formating with `go fmt`.  
This is useful to check if the generated syntax is valid.

The `-b` (`--block`) option prints the output of the preprocessor who's in charge to create the blocks from indentation. No compilation is done.

The `-a` (`--ast`) option prints the generated AST from the parser


## Interpreter (ALPHA)
---

Og embed a small interpreter that in fact compiles the given string into a `/tmp/main.go` skelton and run it. A better implementation will come.

```bash
./og -i
> 1+1
2
```