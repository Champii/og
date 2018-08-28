package og

import (
	"github.com/alecthomas/participle"
)

type INI struct {
	Pack     string    `"package" @Ident`
	ProgBody *ProgBody `[ @@ ]`
	// Comms    []*string   `{ "//" }`
}

type ProgBody struct {
	Imports  []string    `[ "import"  "{" { @String } "}" ]`
	TopLevel []*TopLevel `{ @@ }`
}

type TopLevel struct {
	Struct *Struct `@@`
	Func   *Func   `| @@`
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
	Name       string  `[ @Ident ]`
	Args       []*Arg  `[ "(" { @@ } ")" ] "-" ">"`
	ReturnType string  `[ @Ident ]`
	Body       []*Stmt `[ "{" { @@ } "}" ]`
}

type Arg struct {
	Name string `@Ident`
	Type string `@Ident [","]`
}

type Stmt struct {
	If             *If             `@@`
	For            *For            `| @@`
	Return         *Value          `| ("return" @@)`
	GoRoutine      *GoRoutine      `| @@`
	IdentOrVarDecl *IdentOrVarDecl `| @@`
	// Value             *Value             `| @@`
}

type GoRoutine struct {
	Func  *Func  `( "go" @@ )`
	Value *Value `| ( "go" @@ )`
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
	ElseIf   *ElseIf   `[ "else" @@ ]`
}

type ElseIf struct {
	If   *If   `( @@`
	Else *Else `| @@ )`
	// Body     []*Stmt   `{ @@ } "}"`
	// Nested   *ElseIf   `[ @@ ]`
}

type Else struct {
	Body []*Stmt `"{" { @@ } "}"`
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

type IdentOrVarDecl struct {
	Ident   *NestedProperty `@@`
	VarDecl *VarDecl        `[ @@ ]`
	// FuncCall *FuncCall       `( @@`
}

type ArrAccess struct {
	Value *Value `"[" @@ "]"`
}

// type FuncCallOrVar struct {
// 	Ident    *NestedProperty `@@`
// 	FuncCall *FuncCall       `[ @@ ]`
// }

type VarDecl struct {
	Value *Value `"=" @@`
}

type NestedProperty struct {
	Ident               string                 `@Ident`
	ArrAccessOrFuncCall []*ArrAccessOrFuncCall `[ [ { @@ } ]`
	Nested              *NestedProperty        `[ "." @@ ] ]`
}

type ArrAccessOrFuncCall struct {
	ArrAccess *ArrAccess `( @@`
	FuncCall  *FuncCall  `| @@ )`
}

type FuncCall struct {
	Args []*Value `"(" { @@ [","] } ")"`
}

type Value struct {
	Bool           *bool           `(@"true" | "false")`
	String         *string         `| @String`
	Int            *int64          `| @Int`
	Float          *float64        `| @Float`
	NestedProperty *NestedProperty `| @@`
	ArrDecl        *ArrDecl        `| @@`
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
