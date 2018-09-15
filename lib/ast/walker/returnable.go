package walker

import (
	"github.com/champii/og/lib/ast"
)

type Returnable struct {
	AstWalker
	Root ast.INode
}

func (this *Returnable) VarDecl(n ast.INode) ast.INode {
	varDecl := n.(*ast.VarDecl)
	for _, varSpec := range varDecl.VarSpecs {
		statement := varSpec.Statement
		if statement == nil {
			continue
		}
		ifStmt := statement.IfStmt
		if ifStmt == nil {
			if statement.Block != nil && len(statement.Block.Statements) == 1 && statement.Block.Statements[0].IfStmt != nil {
				ifStmt = statement.Block.Statements[0].IfStmt
			} else {
				continue
			}
		}
		varSpec.Statement = ifStmt.MakeReturnClosureStatement(varSpec.Type)
	}
	return varDecl
}
func (this *Returnable) Function(n ast.INode) ast.INode {
	function := n.(*ast.Function)
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
				block.AddReturn()
			}
			if last.IfStmt != nil {
				last.IfStmt.AddReturn()
			}
		}
	}
	return n
}
func RunReturnable(tree ast.INode) ast.INode {
	returnable := Returnable{Root: tree}
	returnable.type_ = &returnable
	return returnable.Walk(tree)
}
