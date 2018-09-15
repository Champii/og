package og

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	tm "github.com/buger/goterm"
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/ast/walker"
	"github.com/champii/og/lib/common"
	"github.com/champii/og/lib/translator"
	"github.com/champii/og/parser"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

var (
	yellow = color.New(color.FgHiYellow).SprintFunc()
)
var (
	red = color.New(color.FgHiRed).SprintFunc()
)
var (
	cyan = color.New(color.FgCyan).SprintFunc()
)
var (
	magenta = color.New(color.Bold, color.FgHiMagenta).SprintFunc()
)
var (
	blue = color.New(color.Bold, color.FgHiBlue).SprintfFunc()
)
var (
	green = color.New(color.FgHiGreen).SprintfFunc()
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
	for i := 0; i < 8; i++ {
		tm.Println("                                                                          ")
	}
	tm.MoveCursorUp(9)
	tm.Flush()
	fileInfo := fmt.Sprintf("%s (%s:%s)", green(this.filePath), yellow(line), yellow(column))
	badToken := offendingSymbol.(antlr.Token).GetText()
	fmt.Printf("%s: %s '%s'\n", fileInfo, red("Unexpected"), magenta(badToken))
	badLine := this.source[line-1]
	badLine = cyan(badLine[:column]) + magenta(badToken) + cyan(badLine[column+len(badToken):])
	fmt.Println(badLine)
	fmt.Print(blue("%"+strconv.Itoa(column+1)+"s\n\n", "^"))
}
func NewErrorListener(filePath, source string) *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		filePath:             filePath,
		source:               strings.Split(source, "\n"),
	}
}

type OgParser struct {
	Config *OgConfig
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
	tree := t.VisitSourceFile(res.(*parser.SourceFileContext), t).(*ast.SourceFile)
	if this.Config.Ast {
		walker.Print(tree)
	}
	file.Ast = tree
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
func NewOgParser(config *OgConfig) *OgParser {
	return &OgParser{Config: config}
}
