package ast

import (
	"fmt"
	"reflect"
	"strings"
)

type Printer struct {
	AstWalker
	indent int
}

func (this *Printer) Before(n INode) {
	this.indent++
}
func (this *Printer) After(n INode) {
	this.indent--
}
func (this *Printer) Each(n INode) {
	name := reflect.TypeOf(n).String()[5:]
	fmt.Printf("%s\n", strings.Repeat(" ", this.indent)+name)
}
func Print(ast INode) {
	p := Printer{}
	p.type_ = &p
	p.Walk(ast)
}
