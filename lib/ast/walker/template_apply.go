package walker

import (
	"github.com/champii/og/lib/ast"
)

type TemplateApply struct {
	AstWalker
	typeSrc  []string
	typeDest []string
}

func (this *TemplateApply) Type(n ast.INode) ast.INode {
	t := n.(*ast.Type)
	tName := t.Eval()
	for i, ty := range this.typeSrc {
		if tName == ty {
			t.TypeName = this.typeDest[i]
			t.TypeLit = nil
			t.Type = nil
		}
	}
	return n
}
func RunTemplateApply(tree ast.INode, typeSrc []string, typeDest []string) ast.INode {
	templateApply := TemplateApply{
		typeSrc:  typeSrc,
		typeDest: typeDest,
	}
	templateApply.type_ = &templateApply
	return templateApply.Walk(tree)
}
