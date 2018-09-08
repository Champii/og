package ast

import (
	"fmt"
	"reflect"
)

var (
	mangle = 0
)

type Desugar struct {
	AstWalker
}

func (this *Desugar) IfStmt(n INode) INode {
	node := n.(*IfStmt)
	fmt.Println("PARENT", reflect.TypeOf(node.Parent).String()[5:], node.Parent)
	parentName := reflect.TypeOf(node.Parent).String()[5:]
	switch parentName {
	case "IfStmt":
		fmt.Println("HERE")
	case "Statement":
		fmt.Println("HERE2")
	}
	return n
}
func (this *Desugar) Function(n INode) INode {
	function := n.(*Function)
	sig := function.Signature
	if sig == nil {
		return n
	}
	retType := sig.Result
	if retType == nil || len(retType.Types) != 1 {
		return n
	}
	block := function.Block
	if block != nil && len(block.Statements) > 0 {
		last := block.Statements[len(block.Statements)-1]
		if last.ReturnStmt == nil {
			if last.SimpleStmt != nil {
				block.Statements[len(block.Statements)-1] = &Statement{
					Node: NewNodeNoCtx(),
					ReturnStmt: &ReturnStmt{
						Node: NewNodeNoCtx(),
						Expressions: &ExpressionList{
							Node:        NewNodeNoCtx(),
							Expressions: []*Expression{last.SimpleStmt.Expression},
						},
					},
				}
			}
		}
	}
	return n
}
func (this *Block) ReplaceLastStatement(s *Statement) {
	last := this.Statements[len(this.Statements)-1]
	if last.ReturnStmt == nil {
		if last.SimpleStmt != nil {
			this.Statements[len(this.Statements)-1] = s
		}
	}
}
func RunDesugar(ast INode) INode {
	desugar := Desugar{}
	desugar.type_ = &desugar
	return desugar.Walk(ast)
}
