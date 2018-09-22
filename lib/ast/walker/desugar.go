package walker

import (
	"github.com/champii/og/lib/common"
)

type Desugar struct {
	Templates *Templates
}

func (this *Desugar) Run(files []*common.File) {
	RunGobRegister()
	for _, file := range files {
		file.Ast = RunReturnable(file.Ast)
		RunTemplateLoader(file.Ast, this.Templates)
		RunTemplateParse(file, this.Templates)
	}
	this.Templates.Store()
	for _, file := range files {
		RunTemplateUsage(file, this.Templates)
		RunTemplateGenerator(file.Ast, this.Templates)
		this.Templates.ResetUsedFor()
	}
}
func NewDesugar() *Desugar {
	return &Desugar{Templates: NewTemplates()}
}
