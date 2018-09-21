package walker

import (
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/common"
	"os"
	"strings"
)

type TemplateUsage struct {
	AstWalker
	File      *common.File
	Root      common.INode
	Templates *Templates
}

func (this *TemplateUsage) computeTypes(callee common.INode, templateSpec *ast.TemplateSpec) string {
	calleeName := callee.Eval()
	types := []string{}
	for _, t := range templateSpec.Result.Types {
		types = append(types, t.Eval())
	}
	template := this.Templates.Get(calleeName)
	if template == nil {
		err := this.File.Error(callee.Line(), callee.Col(), "Unknown template name", calleeName)
		common.Print.Error(err)
		os.Exit(1)
		return calleeName
	}
	template.AddUsedFor(types)
	return calleeName + strings.Join(types, "")
}
func (this *TemplateUsage) Arguments(n common.INode) common.INode {
	args := n.(*ast.Arguments)
	if args.TemplateSpec != nil {
		callee := args.GetParent().(*ast.SecondaryExpr).GetParent().(*ast.PrimaryExpr).PrimaryExpr.Operand
		callee.OperandName.Name = this.computeTypes(callee, args.TemplateSpec)
	}
	return n
}
func (this *TemplateUsage) CompositeLit(n common.INode) common.INode {
	composite := n.(*ast.CompositeLit)
	if composite.TemplateSpec != nil {
		callee := composite.LiteralType
		callee.Type = this.computeTypes(callee, composite.TemplateSpec)
	}
	return n
}
func RunTemplateUsage(file *common.File, templates *Templates) {
	templateUsage := TemplateUsage{
		Root:      file.Ast,
		File:      file,
		Templates: templates,
	}
	templateUsage.type_ = &templateUsage
	templateUsage.Walk(file.Ast)
}
