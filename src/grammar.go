package og

import (
	"github.com/alecthomas/participle"
)

type OgProg struct {
	Pack     string    `"package" @Ident`
	ProgBody *ProgBody `[ @@ ]`
}

type ProgBody struct {
	Imports  []string    `[ "import"  "{" { @String } "}" ]`
	TopLevel []*TopLevel `{ @@ }`
}

type TopLevel struct {
	Struct *Struct `( @@`
	Func   *Func   `| @@ )`
}

type Struct struct {
	Name   string         `"struct" @Ident "{"`
	Fields []*StructField `{ @@ } "}"`
}

type StructField struct {
	Name string  `@Ident`
	Type string  `@Ident`
	Tag  *string `[ @String ]`
}

type Func struct {
	Name       string  `[ @Ident ]`
	Args       []*Arg  `[ "(" { @@ } ")" ]`
	ReturnType []*Type `[ ":" { @@ [","] } ]`
	Body       []*Stmt `"-" ">" [ ( "{" { @@ } "}" ) | ( @@ ) ]`
}

type Type struct {
	Array []string `{ @("[" "]") | @"*" }`
	Type  string   `@Ident`
}

type Arg struct {
	Name string `@Ident`
	Type *Type  `@@ [","]`
}

type Stmt struct {
	If             *If             `@@`
	For            *For            `| @@`
	Return         []*OuterValue   `| ("return" { @@ [ "," ] } )`
	GoRoutine      *GoRoutine      `| ("go" @@)`
	IdentOrVarDecl *IdentOrVarDecl `| @@`
}

type GoRoutine struct {
	Value *OuterValue `@@`
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
}

type Else struct {
	Body []*Stmt `"{" { @@ } "}"`
}

type Predicat struct {
	First    *OuterValue       `@@`
	Operator *PredicatOperator `@@`
	Second   *OuterValue       `@@`
}

type PredicatOperator struct {
	Eq  string `@(("=" "=") | "is")`
	Neq string `| @(("!" "=") | "isnt")`
	Or  string `@(("|" "|") | "or")`
	And string `@(("&" "&") | "and")`
	Gt  string `| @">"`
	Gte string `| @(">" "=")`
	Lt  string `| @"<"`
	Lte string `| @("<" "=")`
}

type IdentOrVarDecl struct {
	Ident   *NestedProperty `@@`
	VarDecl *VarDecl        `[ @@ ]`
}

type ArrAccess struct {
	Value *OuterValue `"[" @@ "]"`
}

type VarDecl struct {
	Ident []*NestedProperty `{ "," @@ }`
	Value *OuterValue       `"=" @@`
}

type NestedProperty struct {
	Ref                 []string               `{ @"*" | @"&" }`
	Ident               string                 `@Ident`
	ArrAccessOrFuncCall []*ArrAccessOrFuncCall `{ @@ }`
	Nested              *NestedProperty        `[ "." @@ ]`
	StructInst          *StructInst            `[ @@ ]`
	Increment           *Increment             `[ @@ ]`
}

type ArrAccessOrFuncCall struct {
	ArrAccess *ArrAccess `( @@`
	FuncCall  *FuncCall  `| @@ )`
}

type FuncCall struct {
	Args []*OuterValue `"(" { @@ [","] } ")"`
}

type Number struct {
	Float *float64 `( @Float`
	Int   *int64   `| @Int )`
}

type ParenthesisValue struct {
	Open  string      `@"("`
	Value *OuterValue `@@`
	Close string      `@")"`
}

type OuterValue struct {
	NestedProperty   *NestedProperty   `( @@`
	ParenthesisValue *ParenthesisValue `| @@`
	Operation        *Operation        `| @@`
	Func             *Func             `| @@`
	Value            *Value            `| @@ )`
}

type Value struct {
	Bool    *bool    `( @"true" | "false")`
	Nil     *string  `| @"nil"`
	Number  *Number  `| @@`
	String  *string  `| @String`
	ArrDecl *ArrDecl `| @@`
}

type StructInst struct {
	Open  string        `@"{"`
	Ident []string      `{ @Ident ":"`
	Value []*OuterValue `@@ }`
	Close string        `@"}"`
}

type Operation struct {
	First  *Value      `@@`
	Op     *Operator   `[ @@`
	Second *OuterValue `@@ ]`
}

type Operator struct {
	Plus  *string `( @"+"`
	Less  *string `| @"-"`
	Times *string `| @"*"`
	Div   *string `| @"/"`
	Mod   *string `| @"%" )`
}

type Increment struct {
	Inc *string `( @("+" "+")`
	Dec *string `| @("-" "-") )`
}

type ArrDecl struct {
	Type   string        `"[" "]" @Ident "{"`
	Values []*OuterValue `{ @@ [ "," ] } "}"`
}

func Build(str string) (*OgProg, error) {
	parser, err := participle.Build(&OgProg{})

	if err != nil {
		return &OgProg{}, err
	}

	ast := &OgProg{}

	err = parser.ParseString(str, ast)

	if err != nil {
		return &OgProg{}, err
	}

	return ast, nil
}
