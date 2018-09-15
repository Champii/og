package walker

import (
	"github.com/champii/og/lib/common"
)

type Desugar struct {
	Templates *Templates
}

func (this *Desugar) Run(files []*common.File) {
	for _, file := range files {
		file.Ast = RunReturnable(file.Ast)
		RunTemplateParse(file.Ast, this.Templates)
	}
	for _, file := range files {
		RunTemplateUsage(file.Ast, this.Templates)
		RunTemplateGenerator(file.Ast, this.Templates)
		this.Templates.ResetUsedFor()
	}
}
func NewDesugar() *Desugar {
	return &Desugar{Templates: NewTemplates()}
}
