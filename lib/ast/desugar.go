package ast

var (
	mangle = 0
)

type Desugar struct {
	AstWalker
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
			continue
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
	desugar := Desugar{}
	desugar.type_ = &desugar
	return desugar.Walk(ast)
}
