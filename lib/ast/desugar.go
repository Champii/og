package ast

type Desugar struct {
	AstWalker
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
		if last.SimpleStmt != nil && last.ReturnStmt == nil {
			block.Statements[len(block.Statements)-1] = &Statement{ReturnStmt: &ReturnStmt{Expressions: &ExpressionList{Expressions: []*Expression{last.SimpleStmt.Expression},
			},
			},
			}
		}
	}
	return n
}
func RunDesugar(ast INode) INode {
	desugar := Desugar{}
	desugar.type_ = &desugar
	return desugar.Walk(ast)
}
