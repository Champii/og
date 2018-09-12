package walker

import (
	"github.com/champii/og/lib/ast"
)

type TemplateGen struct {
	AstWalker
	typeSrc  []string
	typeDest []string
}

func (this *TemplateGen) Type(n ast.INode) ast.INode {
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
func RunTemplateGen(ast ast.INode, typeSrc []string, typeDest []string) ast.INode {
	templateGen := TemplateGen{
		typeSrc:  typeSrc,
		typeDest: typeDest,
	}
	templateGen.type_ = &templateGen
	res := templateGen.Walk(ast)
	return res
}
