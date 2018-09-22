package walker

import (
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/common"
)

type TemplateReplace struct {
	AstWalker
	typeSrc  []string
	typeDest []string
}

func (this *TemplateReplace) Type(n common.INode) common.INode {
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
func RunTemplateReplace(tree common.INode, typeSrc []string, typeDest []string) common.INode {
	templateReplace := TemplateReplace{
		typeSrc:  typeSrc,
		typeDest: typeDest,
	}
	templateReplace.type_ = &templateReplace
	return templateReplace.Walk(tree)
}
