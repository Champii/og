package og

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/ast/walker"
	"github.com/champii/og/lib/common"
	"github.com/champii/og/lib/translator"
	"github.com/champii/og/parser"
	"os"
	"strings"
)

type ErrorHandler struct {
	*antlr.DefaultErrorStrategy
}

func (this ErrorHandler) Recover(p antlr.Parser, r antlr.RecognitionException) {
	os.Exit(1)
}
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{DefaultErrorStrategy: antlr.NewDefaultErrorStrategy()}
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	filePath string
	source   []string
}

func (this *ErrorListener) SyntaxError(rec antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := common.NewError(this.filePath, this.source, line, column, "Unexpected", offendingSymbol.(antlr.Token).GetText())
	common.Print.Error(err)
}
func NewErrorListener(filePath, source string) *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		filePath:             filePath,
		source:               strings.Split(source, "\n"),
	}
}

type OgParser struct {
	Config *common.OgConfig
}

func (this *OgParser) parserInit(file *common.File) *parser.OgParser {
	input := antlr.NewInputStream(string(file.Output))
	lexer := parser.NewOgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewOgParser(stream)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	p.RemoveErrorListeners()
	p.SetErrorHandler(NewErrorHandler())
	p.AddErrorListener(NewErrorListener(file.Path, string(file.Source)))
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	return p
}
func (this *OgParser) Parse(file *common.File) error {
	p := this.parserInit(file)
	res := p.SourceFile()
	if res == nil {
		return errors.New("Cannot parse file: " + file.Path)
	}
	t := new(translator.OgVisitor)
	t.File = file
	tree := t.VisitSourceFile(res.(*parser.SourceFileContext), t).(*ast.SourceFile)
	if this.Config.Ast || this.Config.SimpleAst {
		walker.Print(tree, this.Config.SimpleAst)
	}
	file.Ast = tree
	file.Imports = tree.GetImports()
	return nil
}
func (this *OgParser) ParseStmt(file *common.File) error {
	p := this.parserInit(file)
	res := p.Statement()
	t := new(translator.OgVisitor)
	file.Ast = t.VisitStatement(res.(*parser.StatementContext), t).(*ast.Statement)
	return nil
}
func (this *OgParser) ParseInterpret(file *common.File) error {
	p := this.parserInit(file)
	res := p.Interp()
	t := new(translator.OgVisitor)
	file.Ast = t.VisitInterp(res.(*parser.InterpContext), t).(*ast.Interpret)
	return nil
}
func NewOgParser(config *common.OgConfig) *OgParser {
	return &OgParser{Config: config}
}
