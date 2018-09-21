package walker

import (
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/common"
)

type TemplateApply struct {
	AstWalker
	typeSrc  []string
	typeDest []string
}

func (this *TemplateApply) Type(n common.INode) common.INode {
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
func RunTemplateApply(tree common.INode, typeSrc []string, typeDest []string) common.INode {
	templateApply := TemplateApply{
		typeSrc:  typeSrc,
		typeDest: typeDest,
	}
	templateApply.type_ = &templateApply
	return templateApply.Walk(tree)
}
