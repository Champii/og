package walker

import (
	"github.com/champii/og/lib/ast"
	"strings"
)

type TemplateUsage struct {
	AstWalker
	Root      ast.INode
	Templates *Templates
}

func (this *TemplateUsage) computeTypes(callee ast.INode, templateSpec *ast.TemplateSpec) string {
	calleeName := callee.Eval()
	types := []string{}
	for _, t := range templateSpec.Result.Types {
		types = append(types, t.Eval())
	}
	template := this.Templates.Get(calleeName)
	template.AddUsedFor(types)
	return calleeName + strings.Join(types, "")
}
func (this *TemplateUsage) Arguments(n ast.INode) ast.INode {
	args := n.(*ast.Arguments)
	if args.TemplateSpec != nil {
		callee := args.GetParent().(*ast.SecondaryExpr).GetParent().(*ast.PrimaryExpr).PrimaryExpr.Operand
		callee.OperandName.Name = this.computeTypes(callee, args.TemplateSpec)
	}
	return n
}
func (this *TemplateUsage) CompositeLit(n ast.INode) ast.INode {
	composite := n.(*ast.CompositeLit)
	if composite.TemplateSpec != nil {
		callee := composite.LiteralType
		callee.Type = this.computeTypes(callee, composite.TemplateSpec)
	}
	return n
}
func RunTemplateUsage(tree ast.INode, templates *Templates) {
	templateUsage := TemplateUsage{
		Root:      tree,
		Templates: templates,
	}
	templateUsage.type_ = &templateUsage
	templateUsage.Walk(tree)
}
