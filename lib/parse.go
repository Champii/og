package og

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"og/lib/translator"
	"og/parser"
)

func Parse(str string) string {
	input := antlr.NewInputStream(str)
	lexer := parser.NewGolangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGolangParser(stream)
	res := p.SourceFile()
	t := new(translator.GolangVisitor)
	final := t.VisitSourceFile(res.(*parser.SourceFileContext), t)
	return final.(string)
}

