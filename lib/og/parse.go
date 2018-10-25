package og

import (
	"errors"
	"fmt"
	"github.com/champii/antlr4/runtime/Go/antlr"
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/ast/walker"
	"github.com/champii/og/lib/common"
	"github.com/champii/og/lib/translator"
	"github.com/champii/og/parser"
	"os"
)

type ErrorHandler struct {
	*antlr.DefaultErrorStrategy
}

func (this ErrorHandler) Recover(p antlr.Parser, r antlr.RecognitionException) {

}
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{DefaultErrorStrategy: antlr.NewDefaultErrorStrategy()}
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	file     *common.File
	IsFaulty bool
	NbErrors int
}

func (this *ErrorListener) SyntaxError(rec antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	this.IsFaulty = true
	common.Print.Error(this.file.Error(line, column, "Unexpected", offendingSymbol.(antlr.Token).GetText()))
	this.NbErrors++
	if this.NbErrors == 5 {
		fmt.Println("Too many errors, exiting")
		os.Exit(1)
	}
}
func NewErrorListener(file *common.File) *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		file:                 file,
	}
}

type OgParser struct {
	Config      *common.OgConfig
	ErrListener *ErrorListener
}

func (this *OgParser) parserInit(file *common.File) *parser.OgParser {
	input := antlr.NewInputStream(string(file.Output))
	lexer := parser.NewOgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewOgParser(stream)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	p.RemoveErrorListeners()
	this.ErrListener = NewErrorListener(file)
	p.AddErrorListener(this.ErrListener)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	return p
}
func (this *OgParser) Parse(file *common.File) error {
	p := this.parserInit(file)
	res := p.SourceFile()
	if res == nil {
		return errors.New("Cannot parse file: " + file.Path)
	}
	if this.ErrListener.IsFaulty {
		return errors.New("")
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
