package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type INode interface {
	Eval() string
	Text() string
	SetParent(n INode)
}
type Node struct {
	Text_    string
	Children []INode
	Parent   INode
}

func (this Node) Eval() string {
	return ""
}
func (this Node) Text() string {
	return this.Text_
}
func (this *Node) SetParent(n INode) {
	if this.Parent == nil {
		this.Parent = n
	}
}
func NewNode(ctx antlr.RuleContext) *Node {
	return &Node{Text_: ctx.GetText()}
}
func NewNodeNoCtx() *Node {
	return &Node{}
}
