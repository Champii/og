package walker

import (
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/common"
	"os"
	"path"
	"strings"
)

type TemplateUsage struct {
	AstWalker
	File      *common.File
	Root      common.INode
	Package   string
	Templates *Templates
}

func (this *TemplateUsage) computeTypes(callee common.INode, templateSpec *ast.TemplateSpec) string {
	calleeName := callee.Eval()
	types := []string{}
	for _, t := range templateSpec.Result.Types {
		types = append(types, t.Eval())
	}
	splited := strings.Split(calleeName, ".")
	pack := this.Package
	if len(splited) > 1 {
		pack = this.File.Imports[splited[0]]
		calleeName = splited[1]
	}
	template := this.Templates.Get(calleeName, pack)
	if template == nil {
		err := this.File.Error(callee.Line(), callee.Col(), "Unknown template name", callee.Eval())
		common.Print.Error(err)
		os.Exit(1)
		return calleeName
	}
	template.AddUsedFor(types)
	prefix := path.Base(template.Pack) + "_"
	return prefix + calleeName + "_" + strings.Join(types, "_")
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
		Package:   path.Dir(file.FullPath),
	}
	templateUsage.type_ = &templateUsage
	templateUsage.Walk(file.Ast)
}
