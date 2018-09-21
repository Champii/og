package walker

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/common"
	"strings"
)

type TemplateGenerator struct {
	Root common.INode
}

func (this *TemplateGenerator) GenerateStruct(template *Template) {
	source := this.Root.(*ast.SourceFile)
	for _, usedFor := range template.UsedFor {
		if template.IsGeneratedFor(usedFor, source.Package.Eval()) {
			continue
		}
		other := &ast.StructType{}
		var (
			buf bytes.Buffer
		)
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(template.Node.(*ast.StructType)); err != nil {
			fmt.Println("ERROR ENCODE", err)
		}
		dec := gob.NewDecoder(&buf)
		if err := dec.Decode(&other); err != nil {
			fmt.Println("ERROR DECODE", err)
		}
		newStruct := RunTemplateApply(other, template.Types, usedFor).(*ast.StructType)
		newStruct.Name += strings.Join(usedFor, "")
		topLevel := &ast.TopLevel{Declaration: &ast.Declaration{TypeDecl: &ast.TypeDecl{StructType: newStruct},
		},
		}
		source.TopLevels = append(source.TopLevels, topLevel)
		template.AddGeneratedFor(usedFor, source.Package.Eval())
	}
}
func (this *TemplateGenerator) GenerateTopFns(template *Template) {
	source := this.Root.(*ast.SourceFile)
	for _, usedFor := range template.UsedFor {
		if template.IsGeneratedFor(usedFor, source.Package.Eval()) {
			continue
		}
		other := &ast.FunctionDecl{}
		var (
			buf bytes.Buffer
		)
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(template.Node.(*ast.FunctionDecl)); err != nil {
			fmt.Println("ERROR ENCODE", err)
		}
		dec := gob.NewDecoder(&buf)
		if err := dec.Decode(&other); err != nil {
			fmt.Println("ERROR DECODE", err)
		}
		newFn := RunTemplateApply(other, template.Types, usedFor).(*ast.FunctionDecl)
		newFn.Name += strings.Join(usedFor, "")
		topLevel := &ast.TopLevel{FunctionDecl: newFn}
		source.TopLevels = append(source.TopLevels, topLevel)
		template.AddGeneratedFor(usedFor, source.Package.Eval())
	}
}
func (this *TemplateGenerator) GenerateGenerics(templates *Templates) {
	if this.Root == nil {
		return
	}
	source := this.Root.(*ast.SourceFile)
	RunGobRegister(source)
	for _, template := range templates.templates {
		topArr := source.TopLevels
		for i, top := range topArr {
			if top.FunctionDecl == template.Node || (top.Declaration != nil && top.Declaration.TypeDecl != nil && top.Declaration.TypeDecl.StructType != nil && top.Declaration.TypeDecl.StructType == template.Node) {
				if len(topArr)-1 == i {
					source.TopLevels = source.TopLevels[:i]
				} else if i == 0 {
					source.TopLevels = source.TopLevels[1:]
				} else {
					source.TopLevels = append(source.TopLevels[0:i], source.TopLevels[i+1:]...)
				}
				break
			}
		}
		switch template.Node.(type) {
		case *ast.FunctionDecl:
			this.GenerateTopFns(template)
		case *ast.StructType:
			this.GenerateStruct(template)
		}
	}
}
func RunTemplateGenerator(tree common.INode, templates *Templates) {
	templateGenerator := TemplateGenerator{Root: tree}
	templateGenerator.GenerateGenerics(templates)
}
