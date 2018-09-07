package ast

type Desugar struct {
	AstWalker
}

func (this *Desugar) FunctionDecl(n INode) INode {
	return n
}
func RunDesugar(ast INode) INode {
	desugar := Desugar{}
	desugar.type_ = &desugar
	return desugar.Walk(ast)
}
