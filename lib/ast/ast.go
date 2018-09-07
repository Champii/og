package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type INode interface {
	IsNode() bool
}
type Node struct {
	start    int
	stop     int
	Text     string
	out      string
	parent   *Node
	children []*Node
	ctx      antlr.RuleContext
}

func (this Node) IsNode() bool {
	return true
}
func (this Node) Eval() string {
	return ""
}
func NewNode(ctx antlr.RuleContext) *Node {
	return &Node{Text: ctx.GetText()}
}

type SourceFile struct {
	*Node
	Package   *Package
	Import    *Import
	TopLevels []*TopLevel
}

func (this SourceFile) Eval() string {
	res := ""
	if this.Package != nil {
		res += this.Package.Eval() + "\n"
	}
	if this.Import != nil {
		res += this.Import.Eval() + "\n"
	}
	for _, t := range this.TopLevels {
		res += t.Eval() + "\n"
	}
	return res
}

type Package struct {
	*Node
	Name string
}

func (this Package) Eval() string {
	return "package " + this.Name
}

type Import struct {
	*Node
	Items []*ImportSpec
}

func (this Import) Eval() string {
	res := "import (\n"
	for _, i := range this.Items {
		res += i.Eval()
	}
	return res + "\n)\n"
}

type ImportSpec struct {
	*Node
	Path  string
	Alias string
}

func (this ImportSpec) Eval() string {
	res := ""
	if len(this.Alias) > 0 {
		res += this.Alias + " "
	}
	res += this.Path
	return res
}

type TopLevel struct {
	*Node
	Declaration  *Declaration
	FunctionDecl *FunctionDecl
	MethodDecl   *MethodDecl
}

func (this TopLevel) Eval() string {
	if this.Declaration != nil {
		return this.Declaration.Eval()
	}
	if this.FunctionDecl != nil {
		return this.FunctionDecl.Eval()
	}
	if this.MethodDecl != nil {
		return this.MethodDecl.Eval()
	}
	return ""
}

type Declaration struct {
	*Node
	ConstDecl *ConstDecl
	TypeDecl  *TypeDecl
	VarDecl   *VarDecl
}

func (this Declaration) Eval() string {
	if this.ConstDecl != nil {
		return this.ConstDecl.Eval()
	}
	if this.TypeDecl != nil {
		return this.TypeDecl.Eval()
	}
	if this.VarDecl != nil {
		return this.VarDecl.Eval()
	}
	return ""
}

type ConstDecl struct {
	*Node
	ConstSpecs []*ConstSpec
}

func (this ConstDecl) Eval() string {
	res := "const ("
	for _, spec := range this.ConstSpecs {
		res += spec.Eval() + "\n"
	}
	return res + ")"
}

type ConstSpec struct {
	*Node
	IdentifierList []string
	Type           *Type
	ExpressionList *ExpressionList
}

func (this ConstSpec) Eval() string {
	res := strings.Join(this.IdentifierList, ",")
	if this.Type != nil {
		res += " " + this.Type.Eval()
	}
	if this.ExpressionList != nil {
		res += " = " + this.ExpressionList.Eval()
	}
	return res
}

type ExpressionList struct {
	*Node
	Expressions []*Expression
}

func (this ExpressionList) Eval() string {
	res := ""
	for _, expr := range this.Expressions {
		res += expr.Eval() + ","
	}
	if len(this.Expressions) > 0 {
		res = res[:len(res)-1]
	}
	return res
}

type Parameters struct {
	*Node
	List []*Parameter
}

func (this Parameters) Eval() string {
	res := "("
	for _, expr := range this.List {
		res += expr.Eval() + ","
	}
	if len(this.List) > 0 {
		res = res[:len(res)-1]
	}
	res += ")"
	return res
}

type TypeDecl struct {
	*Node
	TypeSpecs     []*TypeSpec
	StructType    *StructType
	InterfaceType *InterfaceType
}

func (this TypeDecl) Eval() string {
	res := "type "
	if len(this.TypeSpecs) > 1 {
		res += "("
	}
	for _, spec := range this.TypeSpecs {
		res += spec.Eval() + "\n"
	}
	if len(this.TypeSpecs) > 0 {
		res = res[:len(res)-1]
	}
	if len(this.TypeSpecs) > 1 {
		res += ")"
	}
	if this.StructType != nil {
		res += this.StructType.Eval()
	} else if this.InterfaceType != nil {
		res += this.InterfaceType.Eval()
	}
	return res
}

type TypeSpec struct {
	*Node
	Name string
	Type *Type
}

func (this TypeSpec) Eval() string {
	return this.Name + " " + this.Type.Eval()
}

type FunctionDecl struct {
	*Node
	Name      string
	Function  *Function
	Signature *Signature
}

func (this FunctionDecl) Eval() string {
	res := "func " + this.Name
	if this.Function != nil {
		res += this.Function.Eval()
	}
	if this.Signature != nil {
		res += this.Signature.Eval()
	}
	return res
}

type Function struct {
	*Node
	Signature *Signature
	Block     *Block
}

func (this Function) Eval() string {
	return this.Signature.Eval() + this.Block.Eval()
}

type MethodDecl struct {
	*Node
	Receiver  *Receiver
	Function  *Function
	Signature *Signature
}

func (this MethodDecl) Eval() string {
	res := "func " + this.Receiver.Eval()
	if this.Function != nil {
		res += this.Function.Eval()
	}
	if this.Signature != nil {
		res += this.Signature.Eval()
	}
	return res
}

type Receiver struct {
	*Node
	Package           string
	IsPointerReceiver bool
	Method            string
}

func (this *Receiver) Eval() string {
	res := "(this "
	if this.IsPointerReceiver {
		res += "*"
	}
	res += this.Package + ")"
	res += this.Method
	return res
}

type VarDecl struct {
	*Node
	VarSpecs []*VarSpec
}

func (this VarDecl) Eval() string {
	res := "var ("
	for _, spec := range this.VarSpecs {
		res += spec.Eval() + "\n"
	}
	return res + ")"
}

type VarSpec struct {
	*Node
	IdentifierList []string
	Type           *Type
	ExpressionList *ExpressionList
}

func (this VarSpec) Eval() string {
	res := strings.Join(this.IdentifierList, ",")
	if this.Type != nil {
		res += " " + this.Type.Eval()
	}
	if this.ExpressionList != nil {
		res += "=" + this.ExpressionList.Eval()
	}
	return res
}

type Block struct {
	*Node
	Statements []*Statement
}

func (this Block) Eval() string {
	res := "{\n"
	for _, spec := range this.Statements {
		res += spec.Eval() + "\n"
	}
	return res + "}"
}

type Statement struct {
	*Node
	SimpleStmt      *SimpleStmt
	LabeledStmt     *LabeledStmt
	GoStmt          *GoStmt
	ReturnStmt      *ReturnStmt
	BreakStmt       *BreakStmt
	ContinueStmt    *ContinueStmt
	GotoStmt        *GotoStmt
	FallthroughStmt *FallthroughStmt
	IfStmt          *IfStmt
	SwitchStmt      *SwitchStmt
	SelectStmt      *SelectStmt
	ForStmt         *ForStmt
	Block           *Block
	DeferStmt       *DeferStmt
}

func (this Statement) Eval() string {
	if this.SimpleStmt != nil {
		return this.SimpleStmt.Eval()
	}
	if this.LabeledStmt != nil {
		return this.LabeledStmt.Eval()
	}
	if this.GoStmt != nil {
		return this.GoStmt.Eval()
	}
	if this.ReturnStmt != nil {
		return this.ReturnStmt.Eval()
	}
	if this.BreakStmt != nil {
		return this.BreakStmt.Eval()
	}
	if this.ContinueStmt != nil {
		return this.ContinueStmt.Eval()
	}
	if this.GotoStmt != nil {
		return this.GotoStmt.Eval()
	}
	if this.FallthroughStmt != nil {
		return this.FallthroughStmt.Eval()
	}
	if this.IfStmt != nil {
		return this.IfStmt.Eval()
	}
	if this.SwitchStmt != nil {
		return this.SwitchStmt.Eval()
	}
	if this.SelectStmt != nil {
		return this.SelectStmt.Eval()
	}
	if this.ForStmt != nil {
		return this.ForStmt.Eval()
	}
	if this.Block != nil {
		return this.Block.Eval()
	}
	if this.DeferStmt != nil {
		return this.DeferStmt.Eval()
	}
	return ""
}

type SimpleStmt struct {
	*Node
	SendStmt     *SendStmt
	Expression   *Expression
	IncDecStmt   *IncDecStmt
	ShortVarDecl *ShortVarDecl
	Assignment   *Assignment
	EmptyStmt    bool
}

func (this SimpleStmt) Eval() string {
	if this.SendStmt != nil {
		return this.SendStmt.Eval()
	}
	if this.Expression != nil {
		return this.Expression.Eval()
	}
	if this.IncDecStmt != nil {
		return this.IncDecStmt.Eval()
	}
	if this.ShortVarDecl != nil {
		return this.ShortVarDecl.Eval()
	}
	if this.Assignment != nil {
		return this.Assignment.Eval()
	}
	if this.EmptyStmt {
		return "\n"
	}
	return ""
}

type SendStmt struct {
	*Node
	Left  *Expression
	Right *Expression
}

func (this SendStmt) Eval() string {
	return this.Left.Eval() + "<-" + this.Right.Eval()
}

type IncDecStmt struct {
	*Node
	Expression *Expression
	IsInc      bool
}

func (this IncDecStmt) Eval() string {
	res := this.Expression.Eval()
	if this.IsInc {
		res += "++"
	} else {
		res += "--"
	}
	return res
}

type Assignment struct {
	*Node
	Left  *ExpressionList
	Op    string
	Right *ExpressionList
}

func (this Assignment) Eval() string {
	return this.Left.Eval() + this.Op + this.Right.Eval()
}

type ShortVarDecl struct {
	*Node
	Identifiers []string
	Expressions *ExpressionList
	Statement   *Statement
}

func (this ShortVarDecl) Eval() string {
	res := ""
	res += strings.Join(this.Identifiers, ",") + ":="
	if this.Expressions != nil {
		res += this.Expressions.Eval()
	}
	if this.Statement != nil {
		res += this.Statement.Eval()
	}
	return res
}

type LabeledStmt struct {
	*Node
	Name      string
	Statement *Statement
}

func (this LabeledStmt) Eval() string {
	return this.Name + ": " + this.Statement.Eval()
}

type ReturnStmt struct {
	*Node
	Expressions *ExpressionList
}

func (this ReturnStmt) Eval() string {
	res := "return "
	if this.Expressions != nil {
		res += this.Expressions.Eval()
	}
	return res
}

type BreakStmt struct {
	*Node
	Name string
}

func (this BreakStmt) Eval() string {
	return "break " + this.Name
}

type ContinueStmt struct {
	*Node
	Name string
}

func (this ContinueStmt) Eval() string {
	return "continue " + this.Name
}

type GotoStmt struct {
	*Node
	Name string
}

func (this GotoStmt) Eval() string {
	return "goto " + this.Name
}

type FallthroughStmt struct {
	*Node
}

func (this FallthroughStmt) Eval() string {
	return "fallthrough"
}

type DeferStmt struct {
	*Node
	Expression *Expression
}

func (this DeferStmt) Eval() string {
	return "defer " + this.Expression.Eval()
}

type IfStmt struct {
	*Node
	SimpleStmt *SimpleStmt
	Expression *Expression
	Block      *Block
	IfStmt     *IfStmt
	BlockElse  *Block
}

func (this IfStmt) Eval() string {
	res := "if "
	if this.SimpleStmt != nil {
		res += this.SimpleStmt.Eval() + ";"
	}
	res += this.Expression.Eval()
	if this.Block != nil {
		res += this.Block.Eval()
	}
	if this.BlockElse != nil {
		res += "else " + this.BlockElse.Eval()
	}
	if this.IfStmt != nil {
		res += "else " + this.IfStmt.Eval()
	}
	return res
}

type SwitchStmt struct {
	*Node
	ExprSwitchStmt *ExprSwitchStmt
	TypeSwitchStmt *TypeSwitchStmt
}

func (this SwitchStmt) Eval() string {
	if this.ExprSwitchStmt != nil {
		return this.ExprSwitchStmt.Eval()
	}
	if this.TypeSwitchStmt != nil {
		return this.TypeSwitchStmt.Eval()
	}
	return ""
}

type ExprSwitchStmt struct {
	*Node
	SimpleStmt      *SimpleStmt
	Expression      *Expression
	ExprCaseClauses []*ExprCaseClause
}

func (this ExprSwitchStmt) Eval() string {
	res := "switch "
	if this.SimpleStmt != nil {
		res += this.SimpleStmt.Eval() + ";"
	}
	if this.Expression != nil {
		res += this.Expression.Eval()
	}
	res += "{\n"
	for _, spec := range this.ExprCaseClauses {
		res += spec.Eval() + "\n"
	}
	return res + "}"
}

type ExprCaseClause struct {
	*Node
	ExprSwitchCase *ExprSwitchCase
	Statements     []*Statement
}

func (this ExprCaseClause) Eval() string {
	res := ""
	res += this.ExprSwitchCase.Eval() + ":"
	for _, spec := range this.Statements {
		res += spec.Eval() + "\n"
	}
	if len(this.Statements) > 0 {
		res = res[:len(res)-1]
	}
	return res
}

type ExprSwitchCase struct {
	*Node
	Expressions *ExpressionList
	IsDefault   bool
}

func (this ExprSwitchCase) Eval() string {
	if this.IsDefault {
		return "default"
	}
	return "case " + this.Expressions.Eval()
}

type TypeSwitchStmt struct {
	*Node
	SimpleStmt      *SimpleStmt
	TypeSwitchGuard *TypeSwitchGuard
	TypeCaseClauses []*TypeCaseClause
}

func (this TypeSwitchStmt) Eval() string {
	res := "switch "
	if this.SimpleStmt != nil {
		res += this.SimpleStmt.Eval() + ";"
	}
	res += this.TypeSwitchGuard.Eval()
	res += "{\n"
	for _, spec := range this.TypeCaseClauses {
		res += spec.Eval() + "\n"
	}
	return res + "}"
}

type TypeSwitchGuard struct {
	*Node
	Name        string
	PrimaryExpr *PrimaryExpr
}

func (this TypeSwitchGuard) Eval() string {
	res := ""
	if len(this.Name) > 0 {
		res += this.Name + "="
	}
	return res + this.PrimaryExpr.Eval() + ".(type)"
}

type TypeCaseClause struct {
	*Node
	TypeSwitchCase *TypeSwitchCase
	Statements     []*Statement
}

func (this TypeCaseClause) Eval() string {
	res := this.TypeSwitchCase.Eval() + ":"
	for _, spec := range this.Statements {
		res += spec.Eval() + "\n"
	}
	if len(this.Statements) > 0 {
		res = res[:len(res)-1]
	}
	return res
}

type TypeSwitchCase struct {
	*Node
	Types []*Type
}

func (this TypeSwitchCase) Eval() string {
	res := ""
	if len(this.Types) == 0 {
		return "default"
	}
	res += "case "
	for _, spec := range this.Types {
		res += spec.Eval() + ","
	}
	res = res[:len(res)-1]
	return res
}

type SelectStmt struct {
	*Node
	CommClauses []*CommClause
}

func (this SelectStmt) Eval() string {
	res := "select {\n"
	for _, spec := range this.CommClauses {
		res += spec.Eval()
	}
	return res + "}"
}

type CommClause struct {
	*Node
	CommCase *CommCase
	Block    *Block
}

func (this CommClause) Eval() string {
	block := this.Block.Eval()
	return this.CommCase.Eval() + ":" + block[1:len(block)-1]
}

type CommCase struct {
	*Node
	SendStmt  *SendStmt
	RecvStmt  *RecvStmt
	IsDefault bool
}

func (this CommCase) Eval() string {
	if this.IsDefault {
		return "default"
	}
	if this.SendStmt != nil {
		return "case " + this.SendStmt.Eval()
	}
	if this.RecvStmt != nil {
		return "case " + this.RecvStmt.Eval()
	}
	return ""
}

type RecvStmt struct {
	*Node
	Expressions *ExpressionList
	Identifiers []string
	Expression  *Expression
}

func (this RecvStmt) Eval() string {
	res := ""
	if this.Expressions != nil {
		res += this.Expressions.Eval() + "="
	}
	res += strings.Join(this.Identifiers, ",")
	if len(this.Identifiers) > 0 {
		res += ":="
	}
	res += this.Expression.Eval()
	return res
}

type ForStmt struct {
	*Node
	Expression  *Expression
	ForClause   *ForClause
	RangeClause *RangeClause
	Block       *Block
}

func (this ForStmt) Eval() string {
	res := "for "
	if this.Expression != nil {
		res += this.Expression.Eval()
	}
	if this.ForClause != nil {
		res += this.ForClause.Eval()
	}
	if this.RangeClause != nil {
		res += this.RangeClause.Eval()
	}
	return res + this.Block.Eval()
}

type ForClause struct {
	*Node
	LeftSimpleStmt  *SimpleStmt
	Expression      *Expression
	RightSimpleStmt *SimpleStmt
}

func (this ForClause) Eval() string {
	res := ""
	if this.LeftSimpleStmt != nil {
		res += this.LeftSimpleStmt.Eval()
	}
	res += ";"
	if this.Expression != nil {
		res += this.Expression.Eval()
	}
	res += ";"
	if this.RightSimpleStmt != nil {
		res += this.RightSimpleStmt.Eval()
	}
	return res
}

type RangeClause struct {
	*Node
	Expressions *ExpressionList
	Identifiers []string
	Expression  *Expression
}

func (this RangeClause) Eval() string {
	res := ""
	if this.Expressions != nil {
		res += this.Expressions.Eval() + "="
	}
	res += strings.Join(this.Identifiers, ",")
	if len(this.Identifiers) > 0 {
		res += ":= "
	}
	return res + "range " + this.Expression.Eval()
}

type GoStmt struct {
	*Node
	Function   *Function
	Expression *Expression
}

func (this GoStmt) Eval() string {
	if this.Function != nil {
		return "go func " + this.Function.Eval() + "()"
	}
	if this.Expression != nil {
		return "go " + this.Expression.Eval()
	}
	return ""
}

type Type struct {
	*Node
	TypeName string
	TypeLit  *TypeLit
	Type     *Type
}

func (this Type) Eval() string {
	if this.TypeLit != nil {
		return this.TypeLit.Eval()
	}
	if this.Type != nil {
		return "(" + this.Type.Eval() + ")"
	}
	return this.TypeName
}

type TypeLit struct {
	*Node
	ArrayType     *ArrayType
	StructType    *StructType
	PointerType   *PointerType
	FunctionType  *FunctionType
	InterfaceType *InterfaceType
	SliceType     *SliceType
	MapType       *MapType
	ChannelType   *ChannelType
}

func (this TypeLit) Eval() string {
	if this.ArrayType != nil {
		return this.ArrayType.Eval()
	}
	if this.StructType != nil {
		return this.StructType.Eval()
	}
	if this.PointerType != nil {
		return this.PointerType.Eval()
	}
	if this.FunctionType != nil {
		return this.FunctionType.Eval()
	}
	if this.InterfaceType != nil {
		return this.InterfaceType.Eval()
	}
	if this.SliceType != nil {
		return this.SliceType.Eval()
	}
	if this.MapType != nil {
		return this.MapType.Eval()
	}
	if this.ChannelType != nil {
		return this.ChannelType.Eval()
	}
	return ""
}

type ArrayType struct {
	*Node
	Length      *Expression
	ElementType *Type
}

func (this ArrayType) Eval() string {
	return "[" + this.Length.Eval() + "]" + this.ElementType.Eval()
}

type PointerType struct {
	*Node
	Type *Type
}

func (this *PointerType) Eval() string {
	return "*" + this.Type.Eval()
}

type InterfaceType struct {
	*Node
	Name        string
	MethodSpecs []*MethodSpec
}

func (this InterfaceType) Eval() string {
	res := ""
	if len(this.MethodSpecs) > 0 {
		res += "\n"
	}
	for _, spec := range this.MethodSpecs {
		res += spec.Eval() + "\n"
	}
	return this.Name + " interface {" + res + "}"
}

type SliceType struct {
	*Node
	Type *Type
}

func (this SliceType) Eval() string {
	return "[]" + this.Type.Eval()
}

type MapType struct {
	*Node
	InnerType *Type
	OuterType *Type
}

func (this MapType) Eval() string {
	return "map[" + this.InnerType.Eval() + "]" + this.OuterType.Eval()
}

type ChannelType struct {
	*Node
	ChannelDecl string
	Type        *Type
}

func (this ChannelType) Eval() string {
	return this.ChannelDecl + " " + this.Type.Eval()
}

type MethodSpec struct {
	*Node
	Name       string
	Parameters *Parameters
	Result     *Result
	Type       string
}

func (this MethodSpec) Eval() string {
	res := this.Name
	if this.Parameters != nil {
		res += this.Parameters.Eval()
	}
	if this.Result != nil {
		res += this.Result.Eval()
	}
	res += this.Type
	return res
}

type FunctionType struct {
	*Node
	Signature *Signature
}

func (this FunctionType) Eval() string {
	return "func " + this.Signature.Eval()
}

type Signature struct {
	*Node
	Parameters *Parameters
	Result     *Result
}

func (this Signature) Eval() string {
	res := ""
	if this.Parameters != nil {
		res += this.Parameters.Eval()
	}
	if this.Result != nil {
		res += this.Result.Eval()
	}
	return res
}

type Result struct {
	*Node
	Types []*Type
}

func (this Result) Eval() string {
	res := "("
	for _, spec := range this.Types {
		res += spec.Eval() + ","
	}
	res += ")"
	return res
}

type Parameter struct {
	*Node
	Names      []string
	Type       *Type
	IsVariadic bool
}

func (this Parameter) Eval() string {
	res := ""
	for _, spec := range this.Names {
		res += spec + ","
	}
	if len(this.Names) > 0 {
		res = res[:len(res)-1]
	}
	res += " "
	if this.IsVariadic {
		res += "..."
	}
	return res + this.Type.Eval()
}

type Operand struct {
	*Node
	Literal     *Literal
	OperandName *OperandName
	MethodExpr  *MethodExpr
	Expression  *Expression
}

func (this Operand) Eval() string {
	if this.Literal != nil {
		return this.Literal.Eval()
	}
	if this.OperandName != nil {
		return this.OperandName.Eval()
	}
	if this.MethodExpr != nil {
		return this.MethodExpr.Eval()
	}
	if this.Expression != nil {
		return "(" + this.Expression.Eval() + ")"
	}
	return ""
}

type Literal struct {
	*Node
	Basic     string
	Composite *CompositeLit
}

func (this Literal) Eval() string {
	if this.Composite != nil {
		return this.Composite.Eval()
	}
	return this.Basic
}

type OperandName struct {
	*Node
	Name string
}

func (this OperandName) Eval() string {
	return this.Name
}

type CompositeLit struct {
	*Node
	LiteralType  *LiteralType
	LiteralValue *LiteralValue
}

func (this CompositeLit) Eval() string {
	return this.LiteralType.Eval() + this.LiteralValue.Eval()
}

type LiteralType struct {
	*Node
	Struct  *StructType
	Array   *ArrayType
	Element *Type
	Slice   *SliceType
	Map     *MapType
	Type    string
}

func (this LiteralType) Eval() string {
	if this.Struct != nil {
		return this.Struct.Eval()
	}
	if this.Array != nil {
		return this.Array.Eval()
	}
	if this.Element != nil {
		return "[...]" + this.Element.Eval()
	}
	if this.Slice != nil {
		return this.Slice.Eval()
	}
	if this.Map != nil {
		return this.Map.Eval()
	}
	return this.Type
}

type LiteralValue struct {
	*Node
	Elements []*KeyedElement
}

func (this LiteralValue) Eval() string {
	res := "{"
	if len(this.Elements) > 1 {
		res += "\n"
	}
	for _, spec := range this.Elements {
		res += spec.Eval() + ",\n"
	}
	return res + "}"
}

type KeyedElement struct {
	*Node
	Key     *Key
	Element *Element
}

func (this KeyedElement) Eval() string {
	res := ""
	if this.Key != nil {
		res += this.Key.Eval() + ":"
	}
	return res + this.Element.Eval()
}

type Key struct {
	*Node
	Name         string
	Expression   *Expression
	LiteralValue *LiteralValue
}

func (this Key) Eval() string {
	if this.Expression != nil {
		return this.Expression.Eval()
	}
	if this.LiteralValue != nil {
		return this.LiteralValue.Eval()
	}
	return this.Name
}

type Element struct {
	*Node
	Expression   *Expression
	LiteralValue *LiteralValue
}

func (this Element) Eval() string {
	if this.Expression != nil {
		return this.Expression.Eval()
	}
	if this.LiteralValue != nil {
		return this.LiteralValue.Eval()
	}
	return ""
}

type StructType struct {
	*Node
	Name   string
	Fields []*FieldDecl
}

func (this *StructType) Eval() string {
	res := this.Name + " struct {\n"
	methods := ""
	for _, spec := range this.Fields {
		if spec.InlineStructMethod != nil {
			receiver := this.Name
			if spec.InlineStructMethod.IsPointerReceiver {
				receiver = "*" + receiver
			}
			methods += "\nfunc (this " + receiver + ")" + spec.InlineStructMethod.Eval()
		} else {
			res += spec.Eval() + "\n"
		}
	}
	return res + "}" + methods
}

type FieldDecl struct {
	*Node
	Identifiers        []string
	Type               *Type
	Anonymous          *AnonymousField
	Tag                string
	InlineStructMethod *InlineStructMethod
}

func (this FieldDecl) Eval() string {
	if this.InlineStructMethod != nil {
		return this.InlineStructMethod.Eval()
	}
	res := strings.Join(this.Identifiers, ",")
	if this.Type != nil {
		res += " " + this.Type.Eval()
	}
	if this.Anonymous != nil {
		res += this.Anonymous.Eval()
	}
	return res + " " + this.Tag
}

type InlineStructMethod struct {
	*Node
	IsPointerReceiver bool
	FunctionDecl      *FunctionDecl
}

func (this InlineStructMethod) Eval() string {
	return this.FunctionDecl.Eval()[4:]
}

type AnonymousField struct {
	*Node
	IsPointerReceiver bool
	Type              string
}

func (this *AnonymousField) Eval() string {
	res := this.Type
	if this.IsPointerReceiver {
		res = "*" + res
	}
	return res
}

type FunctionLit struct {
	*Node
	Function *Function
}

func (this FunctionLit) Eval() string {
	return "func " + this.Function.Eval()
}

type PrimaryExpr struct {
	*Node
	Operand       *Operand
	Conversion    *Conversion
	PrimaryExpr   *PrimaryExpr
	Selector      string
	Index         *Index
	Slice         *Slice
	TypeAssertion *TypeAssertion
	Arguments     *Arguments
}

func (this PrimaryExpr) Eval() string {
	if this.Operand != nil {
		return this.Operand.Eval()
	}
	if this.Conversion != nil {
		return this.Conversion.Eval()
	}
	if this.PrimaryExpr != nil {
		res := this.PrimaryExpr.Eval()
		if len(this.Selector) > 0 {
			return res + this.Selector
		}
		if this.Index != nil {
			return res + this.Index.Eval()
		}
		if this.Slice != nil {
			return res + this.Slice.Eval()
		}
		if this.TypeAssertion != nil {
			return res + this.TypeAssertion.Eval()
		}
		if this.Arguments != nil {
			return res + this.Arguments.Eval()
		}
	}
	return ""
}

type Index struct {
	*Node
	Expression *Expression
}

func (this Index) Eval() string {
	return "[" + this.Expression.Eval() + "]"
}

type Slice struct {
	*Node
	LeftExpr   *Expression
	MiddleExpr *Expression
	RightExpr  *Expression
}

func (this Slice) Eval() string {
	res := "["
	if this.LeftExpr != nil {
		res += this.LeftExpr.Eval()
	}
	res += ":"
	if this.MiddleExpr != nil {
		res += this.MiddleExpr.Eval()
	}
	if this.RightExpr != nil {
		res += ":" + this.RightExpr.Eval()
	}
	return res + "]"
}

type TypeAssertion struct {
	*Node
	Type *Type
}

func (this TypeAssertion) Eval() string {
	return ".(" + this.Type.Eval() + ")"
}

type Arguments struct {
	*Node
	Expressions *ExpressionList
	Type        *Type
	IsVariadic  bool
}

func (this Arguments) Eval() string {
	res := "("
	if this.Type != nil {
		res += this.Type.Eval()
		if this.Expressions != nil {
			res += ","
		}
	}
	if this.Expressions != nil {
		res += this.Expressions.Eval()
	}
	if this.IsVariadic {
		res += "..."
	}
	return res + ")"
}

type MethodExpr struct {
	*Node
	ReceiverType *ReceiverType
	Name         string
}

func (this MethodExpr) Eval() string {
	return this.ReceiverType.Eval() + "." + this.Name
}

type ReceiverType struct {
	*Node
	Type         string
	IsPointer    bool
	ReceiverType *ReceiverType
}

func (this *ReceiverType) Eval() string {
	res := ""
	if this.ReceiverType != nil {
		return "(" + this.ReceiverType.Eval() + ")"
	}
	res += this.Type
	if this.IsPointer {
		res = "(*" + res + ")"
	}
	return res
}

type Expression struct {
	*Node
	UnaryExpr       *UnaryExpr
	FunctionLit     *FunctionLit
	LeftExpression  *Expression
	Op              string
	RightExpression *Expression
}

func (this Expression) Eval() string {
	if this.UnaryExpr != nil {
		return this.UnaryExpr.Eval()
	}
	if this.FunctionLit != nil {
		return this.FunctionLit.Eval()
	}
	return this.LeftExpression.Eval() + this.Op + this.RightExpression.Eval()
}

type UnaryExpr struct {
	*Node
	PrimaryExpr *PrimaryExpr
	Op          string
	UnaryExpr   *UnaryExpr
}

func (this UnaryExpr) Eval() string {
	if this.PrimaryExpr != nil {
		return this.PrimaryExpr.Eval()
	}
	return this.Op + this.UnaryExpr.Eval()
}

type Conversion struct {
	*Node
	Type       *Type
	Expression *Expression
}

func (this Conversion) Eval() string {
	return this.Type.Eval() + "(" + this.Expression.Eval() + ")"
}

type Interpret struct {
	*Node
	Statement *Statement
	TopLevel  *TopLevel
}

func (this Interpret) Eval() string {
	if this.Statement != nil {
		return this.Statement.Eval()
	}
	if this.TopLevel != nil {
		return this.TopLevel.Eval()
	}
	return ""
}
