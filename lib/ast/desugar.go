package ast

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

type Template struct {
	Name    string
	Types   []string
	UsedFor [][]string
	Node    INode
}
type GobRegister struct {
	AstWalker
}

func (this *GobRegister) Each(n INode) {
	gob.Register(n)
}
func RunGobRegister(tree INode) {
	p := GobRegister{}
	gob.Register(tree)
	p.type_ = &p
	p.Walk(tree)
}

type Desugar struct {
	AstWalker
	root      INode
	Templates map[string]*Template
}

func (this *Desugar) GenerateTopFns() {
	source := this.root.(*SourceFile)
	RunGobRegister(source)
	for _, template := range this.Templates {
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
		for _, usedFor := range template.UsedFor {
			other := &FunctionDecl{}
			var (
				buf bytes.Buffer
			)
			enc := gob.NewEncoder(&buf)
			if err := enc.Encode(template.Node.(*FunctionDecl)); err != nil {
				fmt.Println("ERROR ENCODE", err)
			}
			dec := gob.NewDecoder(&buf)
			if err := dec.Decode(&other); err != nil {
				fmt.Println("ERROR DECODE", err)
			}
			newFn := RunTemplateGen(other, template.Types, usedFor).(*FunctionDecl)
			newFn.Name += strings.Join(usedFor, "")
			topLevel := &TopLevel{FunctionDecl: newFn}
			source.TopLevels = append(source.TopLevels, topLevel)
		}
	}
}
func (this *Desugar) Arguments(n INode) INode {
	args := n.(*Arguments)
	if args.TemplateSpec != nil {
		callee := args.parent.(*SecondaryExpr).parent.(*PrimaryExpr).PrimaryExpr.Operand
		calleeName := callee.Eval()
		types := []string{}
		for _, t := range args.TemplateSpec.Result.Types {
			types = append(types, t.Eval())
		}
		this.Templates[calleeName].UsedFor = append(this.Templates[calleeName].UsedFor, types)
		callee.OperandName.Name = calleeName + strings.Join(types, "")
	}
	return n
}
func (this *Desugar) StructType(n INode) INode {
	structType := n.(*StructType)
	if structType.TemplateSpec != nil {
		types := []string{}
		for _, t := range structType.TemplateSpec.Result.Types {
			types = append(types, t.Eval())
		}
		this.Templates[structType.Name] = &Template{
			Name:    structType.Name,
			Types:   types,
			UsedFor: [][]string{},
			Node:    structType,
		}
	}
	return n
}
func (this *Desugar) Signature(n INode) INode {
	sig := n.(*Signature)
	if sig.TemplateSpec != nil {
		if f, ok := sig.parent.(*Function); ok {
			fDecl := f.parent.(*FunctionDecl)
			types := []string{}
			for _, t := range sig.TemplateSpec.Result.Types {
				types = append(types, t.Eval())
			}
			this.Templates[fDecl.Name] = &Template{
				Name:    fDecl.Name,
				Types:   types,
				UsedFor: [][]string{},
				Node:    fDecl,
			}
		}
	}
	return n
}
func (this *Desugar) VarDecl(n INode) INode {
	varDecl := n.(*VarDecl)
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
func (this *Desugar) Function(n INode) INode {
	function := n.(*Function)
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
func (this *IfStmt) MakeReturnClosureStatement(t *Type) *Statement {
	this.AddReturn()
	funcLit := &FunctionLit{
		Node: NewNodeNoCtx(),
		Function: &Function{
			Node: NewNodeNoCtx(),
			Signature: &Signature{
				Node: NewNodeNoCtx(),
				Parameters: &Parameters{
					Node: NewNodeNoCtx(),
					List: []*Parameter{},
				},
				Result: &Result{
					Node:  NewNodeNoCtx(),
					Types: []*Type{t},
				},
			},
			Block: &Block{
				Node: NewNodeNoCtx(),
				Statements: []*Statement{&Statement{
					Node:   NewNodeNoCtx(),
					IfStmt: this,
				},
				},
			},
		},
	}
	primary := &PrimaryExpr{
		Node: NewNodeNoCtx(),
		Operand: &Operand{
			Node: NewNodeNoCtx(),
			Literal: &Literal{
				Node:        NewNodeNoCtx(),
				FunctionLit: funcLit,
			},
		},
	}
	unary := &UnaryExpr{
		Node: NewNodeNoCtx(),
		PrimaryExpr: &PrimaryExpr{
			Node:        NewNodeNoCtx(),
			PrimaryExpr: primary,
			SecondaryExpr: &SecondaryExpr{
				Node:      NewNodeNoCtx(),
				Arguments: &Arguments{Node: NewNodeNoCtx()},
			},
		},
	}
	expr := &Expression{
		Node:      NewNodeNoCtx(),
		UnaryExpr: unary,
	}
	stmt := &Statement{
		Node: NewNodeNoCtx(),
		SimpleStmt: &SimpleStmt{
			Node:       NewNodeNoCtx(),
			Expression: expr,
		},
	}
	return stmt
}
func (this *IfStmt) AddReturn() {
	if this.Block != nil {
		this.Block.AddReturn()
	}
	if this.IfStmt != nil {
		this.IfStmt.AddReturn()
	}
	if this.BlockElse != nil {
		this.BlockElse.AddReturn()
	}
}
func (this *Block) AddReturn() {
	last := this.Statements[len(this.Statements)-1]
	if last.ReturnStmt == nil {
		if last.IfStmt != nil {
			last.IfStmt.AddReturn()
		}
		if last.SimpleStmt != nil {
			this.Statements[len(this.Statements)-1] = &Statement{
				Node: NewNodeNoCtx(),
				ReturnStmt: &ReturnStmt{
					Node: NewNodeNoCtx(),
					Expressions: &ExpressionList{
						Node:        NewNodeNoCtx(),
						Expressions: []*Expression{last.SimpleStmt.Expression},
					},
				},
			}
		}
	}
}
func RunDesugar(ast INode) INode {
	desugar := Desugar{
		root:      ast,
		Templates: make(map[string]*Template),
	}
	desugar.type_ = &desugar
	res := desugar.Walk(ast)
	desugar.GenerateTopFns()
	return res
}
