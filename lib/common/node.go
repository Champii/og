package common

import (
	"github.com/champii/antlr4/runtime/Go/antlr"
)

type INode interface {
	Eval() string
	Text() string
	Line() int
	Col() int
	SetParent(n INode)
	GetParent() INode
	ChildrenCount() int
	T() interface{}
}
type Node struct {
	Line_          int
	Col_           int
	Text_          string
	Children       []INode
	parent         INode
	ChildrenCount_ int
	t              interface{}
}

func (this Node) Eval() string {
	return ""
}
func (this Node) Text() string {
	return this.Text_
}
func (this Node) Line() int {
	return this.Line_
}
func (this Node) Col() int {
	return this.Col_
}
func (this Node) ChildrenCount() int {
	return this.ChildrenCount_
}
func (this *Node) SetParent(n INode) {
	if this.parent == nil {
		this.parent = n
	}
}
func (this *Node) GetParent() INode {
	return this.parent
}
func (this Node) T() interface{} {
	return this.t
}
func NewNode(ctx antlr.ParserRuleContext, file *File, t interface{}) *Node {
	tok := ctx.GetStart()
	line := file.LineMapping[tok.GetLine()]
	col := tok.GetColumn()
	return &Node{
		Text_:          ctx.GetText(),
		Line_:          line,
		Col_:           col,
		ChildrenCount_: len(ctx.GetChildren()),
		t:              t,
	}
}
func NewNodeNoCtx(t interface{}) *Node {
	return &Node{t: t}
}
