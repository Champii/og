package main

import (
	"github.com/alecthomas/participle"
)

type INI struct {
	Pack     string      `"package" @Ident`
	Imports  []string    `"import" "{" { @String } "}"`
	TopLevel []*TopLevel `{ @@ }`
	// Comms    []*string   `{ "//" }`
}

type TopLevel struct {
	Structs []*Struct `@@`
	Funcs   []*Func   `| @@`
}

type Struct struct {
	Name   string         `"struct" @Ident "{"`
	Fields []*StructField `{ @@ } "}"`
}

type StructField struct {
	Name string `@Ident`
	Type string `@Ident`
}

type Func struct {
	Name       string  `@Ident`
	Args       []*Arg  `[ "(" { @@ } ")" ] "-" ">"`
	ReturnType string  `[ @Ident ]`
	Body       []*Stmt `[ "{" { @@ } "}" ]`
}

type Arg struct {
	Name string `@Ident`
	Type string `@Ident [","]`
}

type Stmt struct {
	FuncCallOrVarDecl *FuncCallOrVarDecl `@@`
	If                *If                `| @@`
	For               *For               `| @@`
	// Value             *Value             `| @@`
	Return *Value `| ("return" @@)`
}

type For struct {
	Iterator string  `"for" @Ident`
	Value    string  `[ "," @Ident ]`
	Source   string  `"in" @Ident "{"`
	Body     []*Stmt `{ @@ } "}"`
}

type If struct {
	Predicat *Predicat `"if" @@ "{"`
	Body     []*Stmt   `{ @@ } "}"`
	ElseIf   []*ElseIf `{ @@ }`
	Else     *Else     `[ @@ ]`
}

type ElseIf struct {
	Predicat *Predicat `"else if" @@ "{"`
	Body     []*Stmt   `{ @@ } "}"`
}

type Else struct {
	Body []*Stmt `"else" "{" { @@ } "}"`
}

type Predicat struct {
	First    *Value    `@@`
	Operator *Operator `@@`
	Second   *Value    `@@`
}

type Operator struct {
	Eq  string `@(("=" "=") | "is")`
	Neq string `| @(("!" "=") | "isnt")`
	Gt  string `| @">"`
	Gte string `| @(">" "=")`
	Lt  string `| @"<"`
	Lte string `| @("<" "=")`
}

type FuncCallOrVarDecl struct {
	Ident    *NestedProperty `@@`
	FuncCall *FuncCall       `( @@`
	VarDecl  *VarDecl        `| @@ )`
}

type ArrAccess struct {
	Value *Value `"[" @@ "]"`
}
type FuncCallOrVar struct {
	Ident    *NestedProperty `@@`
	FuncCall *FuncCall       `[ @@ ]`
}

type VarDecl struct {
	Value *Value `"=" @@`
}

type NestedProperty struct {
	Ident     string          `( @Ident`
	ArrAccess []*ArrAccess    `{ @@ } )`
	Nested    *NestedProperty `[ "." @@ ]`
}

type FuncCall struct {
	Args []*Value `"(" { @@ [","] } ")"`
}

type Value struct {
	Bool          *bool          `(@"true" | "false")`
	String        *string        `| @String`
	Int           *int64         `| @Int`
	Float         *float64       `| @Float`
	FuncCallOrVar *FuncCallOrVar `| @@`
	ArrDecl       *ArrDecl       `| @@`
}

type ArrDecl struct {
	Type   string   `"[" "]" @Ident "{"`
	Values []*Value `{ @@ [ "," ] } "}"`
}

func Build(str string) (*INI, error) {
	parser, err := participle.Build(&INI{}, participle.UseLookahead())

	if err != nil {
		return &INI{}, err
	}

	ast := &INI{}

	err = parser.ParseString(str, ast)

	if err != nil {
		return &INI{}, err
	}

	return ast, nil
}
