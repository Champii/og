package walker

import (
	"fmt"
	"github.com/champii/og/lib/ast"
	"reflect"
	"strings"
)

type Printer struct {
	AstWalker
	indent int
}

func (this *Printer) Before(n ast.INode) {
	this.indent++
}
func (this *Printer) After(n ast.INode) {
	this.indent--
}
func (this *Printer) Each(n ast.INode) ast.INode {
	name := reflect.TypeOf(n).String()[5:]
	fmt.Printf("%s\n", strings.Repeat(" ", this.indent)+name)
	return n
}
func Print(ast ast.INode) {
	p := Printer{}
	p.type_ = &p
	p.Walk(ast)
}
