package og

import (
	"Og/og/translator"
	"Og/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Parse(str string) string {
	input := antlr.NewInputStream(str)

	lexer := parser.NewGolangLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewGolangParser(stream)

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	res := p.SourceFile()

	// res.ToStringTree()

	// res.ToStringTree([]string{"sourceFile"})

	t := new(translator.GolangVisitor)

	// res.Accept(t)
	final := t.VisitSourceFile(res.(*parser.SourceFileContext), t)

	// antlr.ParseTreeWalkerDefault.Walk(t, res)

	return final.(string)
}
