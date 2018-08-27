package main

import (
	"github.com/alecthomas/participle"
)

type INI struct {
	Pack     string      `"package" @Ident`
	Imports  []string    `"import" "{" { @String } "}"`
	TopLevel []*TopLevel `{ @@ }`
	// Comms    []*string   `{ "//" @String }`
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
	Name string  `@Ident`
	Args []*Arg  `[ "(" { @@ } ")" ] "-" ">"`
	Body []*Stmt `[ "{" { @@ } "}" ]`
}

type Arg struct {
	Name string `@Ident`
	Type string `@Ident [","]`
}

type Stmt struct {
	Ident    *NestedProperty `@@`
	FuncCall *FuncCall       `( @@ |`
	VarDecl  *VarDecl        `@@ )`
}

type VarDecl struct {
	Value *Value `"=" @@`
}

type NestedProperty struct {
	Ident  string          `@Ident`
	Nested *NestedProperty `[ "." @@ ]`
}

type FuncCall struct {
	Args []*Value `"(" { @@ [","] } ")"`
}

type Value struct {
	String   string          `@String |`
	Number   float64         `@Float  |`
	Ident    *NestedProperty `@@ |`
	FuncCall *FuncCall       `@@`
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
