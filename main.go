package main

import (
	og "Og/lib"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"

	"golang.org/x/exp/ebnf"
)

func main() {
	readFile, err := ioutil.ReadFile("./parser/grammar.ebnf")
	if err != nil {
		panic(err)
	}

	reader, err := os.Open("./parser/grammar.ebnf")

	if err != nil {
		panic(err)
	}

	gram, err := ebnf.Parse("grammar.ebnf", reader)

	if err != nil {
		panic(err)
	}

	fmt.Println(gram)

	err = ebnf.Verify(gram, "SourceFile")

	if err != nil {
		panic(err)
	}

	lexer, err := lexer.EBNF(string(readFile))

	if err != nil {
		panic(err)
	}

	parser, err := participle.Build(&og.OgProg{}, participle.Lexer(lexer))

	if err != nil {
		panic(err)
	}

	source, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	ast := &og.OgProg{}

	err = parser.ParseBytes(source, ast)

	if err != nil {
		panic(err)
	}

	fmt.Println("AST", ast)

}

// func main() {
// 	args := os.Args[1:]

// 	for _, arg := range args {
// 		fmt.Println(og.Compile(arg))
// 	}
// }
