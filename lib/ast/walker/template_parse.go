package walker

import (
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/common"
	"path"
)

type TemplateParse struct {
	AstWalker
	Root       common.INode
	Package    string
	Templates  *Templates
	ImportName string
}

func (this *TemplateParse) StructType(n common.INode) common.INode {
	structType := n.(*ast.StructType)
	if structType.TemplateSpec != nil {
		types := []string{}
		for _, t := range structType.TemplateSpec.Result.Types {
			types = append(types, t.Eval())
		}
		this.Templates.Add(structType.Name, this.ImportName, NewTemplate(structType.Name, this.ImportName, types, structType))
	}
	return n
}
func (this *TemplateParse) Signature(n common.INode) common.INode {
	sig := n.(*ast.Signature)
	if sig.TemplateSpec != nil {
		if f, ok := sig.GetParent().(*ast.Function); ok {
			fDecl := f.GetParent().(*ast.FunctionDecl)
			types := []string{}
			for _, t := range sig.TemplateSpec.Result.Types {
				types = append(types, t.Eval())
			}
			this.Templates.Add(fDecl.Name, this.ImportName, NewTemplate(fDecl.Name, this.ImportName, types, fDecl))
		}
	}
	return n
}
func RunTemplateParse(file *common.File, templates *Templates) {
	templateParse := TemplateParse{
		Root:       file.Ast,
		Templates:  templates,
		Package:    file.Ast.(*ast.SourceFile).Package.Name,
		ImportName: path.Dir(file.FullPath),
	}
	templateParse.type_ = &templateParse
	templateParse.Walk(file.Ast)
}
