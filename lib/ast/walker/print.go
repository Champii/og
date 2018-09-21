package walker

import (
	"fmt"
	"github.com/champii/og/lib/common"
	"github.com/fatih/color"
	"reflect"
	"strconv"
	"strings"
)

var (
	magenta = color.New(color.Bold, color.FgHiMagenta).SprintFunc()
)
var (
	yellow = color.New(color.FgHiYellow).SprintfFunc()
)
var (
	green = color.New(color.FgHiGreen).SprintfFunc()
)
var (
	cyan = color.New(color.FgCyan).SprintFunc()
)

type Printer struct {
	AstWalker
	Simple bool
	indent int
}

func (this *Printer) Before(n common.INode) {
	this.indent++
}
func (this *Printer) After(n common.INode) {
	this.indent--
}
func (this Printer) getTerminalText(n common.INode) string {
	txt := ""
	struc := reflect.ValueOf(n).Elem()
	for i := 0; i < struc.NumField(); i++ {
		if struc.Field(i).Kind() == reflect.String {
			txt = struc.Field(i).String()
		}
		if struc.Field(i).Kind() == reflect.Slice {
			if struc.Field(i).Type().Elem().Kind() == reflect.String {
				txt = strings.Join(struc.Field(i).Interface().([]string), ",")
			}
		}
	}
	return txt
}
func (this *Printer) Each(n common.INode) common.INode {
	name := reflect.TypeOf(n).String()[5:]
	txt := this.getTerminalText(n)
	if this.Simple && len(txt) == 0 {
		return n
	}
	line := fmt.Sprintf("(%s:%s):", yellow("%d", n.Line()), yellow("%d", n.Col()))
	line += strings.Repeat(" ", 8-len(strconv.Itoa(n.Line())+strconv.Itoa(n.Col())))
	if len(txt) > 0 {
		name = green(name)
	} else {
		name = cyan(name)
	}
	fmt.Printf("%s %s %s\n", line, strings.Repeat(" ", this.indent)+name, magenta(txt))
	return n
}
func Print(ast common.INode, simple bool) {
	p := Printer{Simple: simple}
	p.type_ = &p
	p.Walk(ast)
}
