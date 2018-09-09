package ast

type TemplateGen struct {
	AstWalker
	typeSrc  []string
	typeDest []string
}

func (this *TemplateGen) Type(n INode) INode {
	t := n.(*Type)
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
func RunTemplateGen(ast INode, typeSrc []string, typeDest []string) INode {
	templateGen := TemplateGen{
		typeSrc:  typeSrc,
		typeDest: typeDest,
	}
	templateGen.type_ = &templateGen
	res := templateGen.Walk(ast)
	return res
}
