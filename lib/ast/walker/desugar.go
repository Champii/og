package walker

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/champii/og/lib/ast"
	"strings"
)

type Template struct {
	Name    string
	Types   []string
	UsedFor [][]string
	Node    ast.INode
}
type GobRegister struct {
	AstWalker
}

func (this *GobRegister) Each(n ast.INode) {
	gob.Register(n)
}
func RunGobRegister(tree ast.INode) {
	p := GobRegister{}
	gob.Register(tree)
	p.type_ = &p
	p.Walk(tree)
}

type Templates struct {
	names     []string
	templates []*Template
}

func (this *Templates) Add(name string, template *Template) {
	this.names = append(this.names, name)
	this.templates = append(this.templates, template)
}
func (this *Templates) Get(name string) *Template {
	for i, n := range this.names {
		if n == name {
			return this.templates[i]
		}
	}
	return nil
}

type Desugar struct {
	AstWalker
	Root      ast.INode
	Templates Templates
}

func (this *Desugar) GenerateStruct(template *Template) {
	source := this.Root.(*ast.SourceFile)
	for _, usedFor := range template.UsedFor {
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
		newStruct := RunTemplateGen(other, template.Types, usedFor).(*ast.StructType)
		newStruct.Name += strings.Join(usedFor, "")
		topLevel := &ast.TopLevel{Declaration: &ast.Declaration{TypeDecl: &ast.TypeDecl{StructType: newStruct},
		},
		}
		source.TopLevels = append(source.TopLevels, topLevel)
	}
}
func (this *Desugar) GenerateTopFns(template *Template) {
	source := this.Root.(*ast.SourceFile)
	for _, usedFor := range template.UsedFor {
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
		newFn := RunTemplateGen(other, template.Types, usedFor).(*ast.FunctionDecl)
		newFn.Name += strings.Join(usedFor, "")
		topLevel := &ast.TopLevel{FunctionDecl: newFn}
		source.TopLevels = append(source.TopLevels, topLevel)
	}
}
func (this *Desugar) GenerateGenerics() {
	source := this.Root.(*ast.SourceFile)
	RunGobRegister(source)
	for _, template := range this.Templates.templates {
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
func (this *Desugar) Arguments(n ast.INode) ast.INode {
	args := n.(*ast.Arguments)
	if args.TemplateSpec != nil {
		callee := args.GetParent().(*ast.SecondaryExpr).GetParent().(*ast.PrimaryExpr).PrimaryExpr.Operand
		calleeName := callee.Eval()
		types := []string{}
		for _, t := range args.TemplateSpec.Result.Types {
			types = append(types, t.Eval())
		}
		template := this.Templates.Get(calleeName)
		template.UsedFor = append(template.UsedFor, types)
		callee.OperandName.Name = calleeName + strings.Join(types, "")
	}
	return n
}
func (this *Desugar) CompositeLit(n ast.INode) ast.INode {
	composite := n.(*ast.CompositeLit)
	if composite.TemplateSpec != nil {
		callee := composite.LiteralType
		calleeName := callee.Eval()
		types := []string{}
		for _, t := range composite.TemplateSpec.Result.Types {
			types = append(types, t.Eval())
		}
		template := this.Templates.Get(calleeName)
		template.UsedFor = append(template.UsedFor, types)
		callee.Type = calleeName + strings.Join(types, "")
	}
	return n
}
func (this *Desugar) StructType(n ast.INode) ast.INode {
	structType := n.(*ast.StructType)
	if structType.TemplateSpec != nil {
		types := []string{}
		for _, t := range structType.TemplateSpec.Result.Types {
			types = append(types, t.Eval())
		}
		this.Templates.Add(structType.Name, &Template{
			Name:    structType.Name,
			Types:   types,
			UsedFor: [][]string{},
			Node:    structType,
		})
	}
	return n
}
func (this *Desugar) Signature(n ast.INode) ast.INode {
	sig := n.(*ast.Signature)
	if sig.TemplateSpec != nil {
		if f, ok := sig.GetParent().(*ast.Function); ok {
			fDecl := f.GetParent().(*ast.FunctionDecl)
			types := []string{}
			for _, t := range sig.TemplateSpec.Result.Types {
				types = append(types, t.Eval())
			}
			this.Templates.Add(fDecl.Name, &Template{
				Name:    fDecl.Name,
				Types:   types,
				UsedFor: [][]string{},
				Node:    fDecl,
			})
		}
	}
	return n
}
func (this *Desugar) VarDecl(n ast.INode) ast.INode {
	varDecl := n.(*ast.VarDecl)
	for _, varSpec := range varDecl.VarSpecs {
		statement := varSpec.Statement
		if statement == nil {
			continue
		}
		ifStmt := statement.IfStmt
		if ifStmt == nil {
			if statement.Block != nil && len(statement.Block.Statements) == 1 && statement.Block.Statements[0].IfStmt != nil {
				ifStmt = statement.Block.Statements[0].IfStmt
			} else {
				continue
			}
		}
		varSpec.Statement = ifStmt.MakeReturnClosureStatement(varSpec.Type)
	}
	return varDecl
}
func (this *Desugar) Function(n ast.INode) ast.INode {
	function := n.(*ast.Function)
	sig := function.Signature
	if sig == nil {
		return n
	}
	retType := sig.Result
	if retType == nil || len(retType.Types) != 1 {
		return n
	}
	block := function.Block
	if block != nil && len(block.Statements) > 0 {
		last := block.Statements[len(block.Statements)-1]
		if last.ReturnStmt == nil {
			if last.SimpleStmt != nil {
				block.AddReturn()
			}
			if last.IfStmt != nil {
				last.IfStmt.AddReturn()
			}
		}
	}
	return n
}
func RunDesugar(tree ast.INode) ast.INode {
	desugar := Desugar{Root: tree}
	desugar.type_ = &desugar
	res := desugar.Walk(tree)
	desugar.GenerateGenerics()
	return res
}
func NewDesugar() *Desugar {
	desugar := &Desugar{}
	desugar.type_ = desugar
	return desugar
}
