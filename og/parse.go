package og

import (
	"Og/og/translator"
	"Og/parser"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Parse(str string) string {
	input := antlr.NewInputStream(str)

	lexer := parser.NewGolangLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewGolangParser(stream)

	res := p.SourceFile()

	t := new(translator.Translator)

	antlr.ParseTreeWalkerDefault.Walk(t, res)

	return strings.Join(t.Out, " ")
}
