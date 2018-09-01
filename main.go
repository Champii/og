package main

import (
	"Og/toto"
	"fmt"
	"os"
)

//go:generate java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar" org.antlr.v4.Tool -Dlanguage=Go Golang.g4 -visitor -o parser/
//go:generate ./scripts/fix_parser_imports.sh
func main() {
	args := os.Args[1:]

	for _, arg := range args {
		fmt.Println(og.Compile(arg))
	}
}
