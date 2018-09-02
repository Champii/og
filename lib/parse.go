package og

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/champii/og/lib/translator"
	"github.com/champii/og/parser"
)

func Parse(str string) string {
	input := antlr.NewInputStream(str)
	lexer := parser.NewOgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewOgParser(stream)
	res := p.SourceFile()
	t := new(translator.OgVisitor)
	final := t.VisitSourceFile(res.(*parser.SourceFileContext), t)
	return final.(string)
}
