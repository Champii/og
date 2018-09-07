package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
)

type INode interface {
	Eval() string
	Text() string
	GetName() string
}
type Node struct {
	Text_    string
	Children []INode
}

func (this Node) GetName() string {
	return reflect.ValueOf(this).Type().Name()
}
func (this Node) Eval() string {
	return ""
}
func (this Node) Text() string {
	return this.Text_
}
func NewNode(ctx antlr.RuleContext) *Node {
	return &Node{Text_: ctx.GetText()}
}
