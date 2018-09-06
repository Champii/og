package og

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/champii/og/lib/translator"
	"github.com/champii/og/parser"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

var yellow = color.New(color.FgHiYellow).SprintFunc()
var red = color.New(color.FgHiRed).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()
var magenta = color.New(color.Bold, color.FgHiMagenta).SprintFunc()
var blue = color.New(color.Bold, color.FgHiBlue).SprintfFunc()
var green = color.New(color.FgHiGreen).SprintfFunc()

type ErrorHandler struct {
	*antlr.DefaultErrorStrategy
}

func (this ErrorHandler) Recover(p antlr.Parser, r antlr.RecognitionException) {
	os.Exit(1)
}
func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		DefaultErrorStrategy: antlr.NewDefaultErrorStrategy(),
	}
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	source []string
}

func (this *ErrorListener) SyntaxError(rec antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fileInfo := fmt.Sprintf("%s (%s:%s)", green("path/file.og"), yellow(line), yellow(column))
	badToken := offendingSymbol.(antlr.Token).GetText()
	fmt.Printf("%s: %s '%s'\n", fileInfo, red("Unexpected"), magenta(badToken))
	badLine := this.source[line-1]
	badLine = cyan(badLine[:column]) + magenta(badToken) + cyan(badLine[column+len(badToken):])
	fmt.Println(badLine)
	fmt.Print(blue("%"+strconv.Itoa(column+1)+"s\n", "^"))
}
func NewErrorListener(source string) *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
		source:               strings.Split(source, "\n"),
	}
}
func Parse(str string) string {
	input := antlr.NewInputStream(str)
	lexer := parser.NewOgLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewOgParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(NewErrorListener(str))
	p.SetErrorHandler(NewErrorHandler())
	res := p.SourceFile()
	t := new(translator.OgVisitor)
	final := t.VisitSourceFile(res.(*parser.SourceFileContext), t)
	return final.(string)
}
