package translator

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/champii/og/parser"
	"strings"
)

type OgVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (this *OgVisitor) Aggregate(resultSoFar, childResult interface {
}) interface{} {
	switch childResult.(type) {
	default:
		return resultSoFar
	case string:
		{
			switch resultSoFar.(type) {
			case string:
				return resultSoFar.(string) + childResult.(string)
			default:
				return childResult
			}
		}
	}
	return nil
}
func (this *OgVisitor) VisitSourceFile(ctx *parser.SourceFileContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitPackageClause(ctx *parser.PackageClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "package " + ctx.IDENTIFIER().GetText() + "\n"
}
func (this *OgVisitor) VisitImportDecl(ctx *parser.ImportDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "import (\n" + this.VisitChildren(ctx, delegate).(string) + ")\n"
}
func (this *OgVisitor) VisitImportSpec(ctx *parser.ImportSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitImportPath(ctx *parser.ImportPathContext, delegate antlr.ParseTreeVisitor) interface{} {
	txt := ctx.GetText()
	if txt[0] == '"' {
		return txt + "\n"
	} else {
		return "\"" + txt + "\"\n"
	}
}
func (this *OgVisitor) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate).(string) + "\n"
}
func (this *OgVisitor) VisitDeclaration(ctx *parser.DeclarationContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitConstDecl(ctx *parser.ConstDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitConstSpec(ctx *parser.ConstSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitExpressionList(ctx *parser.ExpressionListContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
	if ctx.GetChildCount() > 1 {
		r += "," + this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(string)
	}
	return r
}
func (this *OgVisitor) VisitTypeDecl(ctx *parser.TypeDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "type " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.IDENTIFIER().GetText() + " " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "func " + ctx.IDENTIFIER().GetText() + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitFunction(ctx *parser.FunctionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitMethodDecl(ctx *parser.MethodDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "func " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitReceiver(ctx *parser.ReceiverContext, delegate antlr.ParseTreeVisitor) interface{} {
	c := ctx.IDENTIFIER(0).GetText()
	method := ctx.IDENTIFIER(1).GetText()
	return "(this *" + c + ") " + method
}
func (this *OgVisitor) VisitVarDecl(ctx *parser.VarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "var " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitVarSpec(ctx *parser.VarSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitBlock(ctx *parser.BlockContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "{\n" + this.VisitChildren(ctx, delegate).(string) + "}"
}
func (this *OgVisitor) VisitStatementList(ctx *parser.StatementListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitStatement(ctx *parser.StatementContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate).(string) + "\n"
}
func (this *OgVisitor) VisitStatementNoBlock(ctx *parser.StatementNoBlockContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "{" + this.VisitChildren(ctx, delegate).(string) + "}"
}
func (this *OgVisitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitSendStmt(ctx *parser.SendStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitIncDecStmt(ctx *parser.IncDecStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitAssignment(ctx *parser.AssignmentContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitAssign_op(ctx *parser.Assign_opContext, delegate antlr.ParseTreeVisitor) interface{} {
	if len(ctx.GetText()) == 1 {
		return "="
	}
	return ctx.GetText()
}
func (this *OgVisitor) VisitBinary_op(ctx *parser.Binary_opContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitShortVarDecl(ctx *parser.ShortVarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	idList := this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).(string)
	exprList := this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(string)
	return idList + " := " + exprList
}
func (this *OgVisitor) VisitEmptyStmt(ctx *parser.EmptyStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitLabeledStmt(ctx *parser.LabeledStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetChildCount() == 0 {
		{
			return "return"
		}
	}
	return "return " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitGotoStmt(ctx *parser.GotoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitFallthroughStmt(ctx *parser.FallthroughStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitDeferStmt(ctx *parser.DeferStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "defer " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitIfStmt(ctx *parser.IfStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := "if "
	if ctx.SimpleStmt() != nil {
		r += this.VisitSimpleStmt(ctx.SimpleStmt().(*parser.SimpleStmtContext), delegate).(string) + ";"
	}
	r += this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
	r += this.VisitBlock(ctx.Block(0).(*parser.BlockContext), delegate).(string)
	if ctx.Block(1) != nil {
		r += "else " + this.VisitBlock(ctx.Block(1).(*parser.BlockContext), delegate).(string)
	} else if ctx.IfStmt() != nil {
		r += "else " + this.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext), delegate).(string)
	}
	return r
}
func (this *OgVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitExprSwitchStmt(ctx *parser.ExprSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := "switch "
	if ctx.Expression() != nil {
		r += this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
	}
	r += "{\n"
	for _, c := range ctx.AllExprCaseClause() {
		r += this.VisitExprCaseClause(c.(*parser.ExprCaseClauseContext), delegate).(string)
	}
	r += "}"
	return r
}
func (this *OgVisitor) VisitExprCaseClause(ctx *parser.ExprCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	sCase := this.VisitExprSwitchCase(ctx.ExprSwitchCase().(*parser.ExprSwitchCaseContext), delegate).(string)
	stmts := this.VisitStatementList(ctx.StatementList().(*parser.StatementListContext), delegate).(string)
	return sCase + ":" + stmts
}
func (this *OgVisitor) VisitExprSwitchCase(ctx *parser.ExprSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetText() == "_" {
		return "default"
	}
	return "case " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitTypeSwitchStmt(ctx *parser.TypeSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := "switch "
	r += this.VisitTypeSwitchGuard(ctx.TypeSwitchGuard().(*parser.TypeSwitchGuardContext), delegate).(string)
	r += "{"
	for _, c := range ctx.AllTypeCaseClause() {
		r += this.VisitTypeCaseClause(c.(*parser.TypeCaseClauseContext), delegate).(string)
	}
	r += "}"
	return r
}
func (this *OgVisitor) VisitTypeSwitchGuard(ctx *parser.TypeSwitchGuardContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := ""
	if ctx.IDENTIFIER() != nil {
		r += ctx.IDENTIFIER().GetText() + "="
	}
	expr := this.VisitPrimaryExpr(ctx.PrimaryExpr().(*parser.PrimaryExprContext), delegate).(string)
	return r + expr + ".(type)"
}
func (this *OgVisitor) VisitTypeCaseClause(ctx *parser.TypeCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	sCase := this.VisitTypeSwitchCase(ctx.TypeSwitchCase().(*parser.TypeSwitchCaseContext), delegate).(string)
	stmts := this.VisitStatementList(ctx.StatementList().(*parser.StatementListContext), delegate).(string)
	return sCase + ":" + stmts
}
func (this *OgVisitor) VisitTypeSwitchCase(ctx *parser.TypeSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetText() == "_" {
		return "default"
	}
	return "case " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitTypeList(ctx *parser.TypeListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitSelectStmt(ctx *parser.SelectStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitCommClause(ctx *parser.CommClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitCommCase(ctx *parser.CommCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitRecvStmt(ctx *parser.RecvStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitForStmt(ctx *parser.ForStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "for " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitForClause(ctx *parser.ForClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitRangeClause(ctx *parser.RangeClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := ""
	if ctx.IdentifierList() != nil {
		r = this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).(string)
	} else if ctx.ExpressionList() != nil {
		r = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(string)
	}
	return r + " := range " + this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
}
func (this *OgVisitor) VisitGoStmt(ctx *parser.GoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	expr := this.VisitChildren(ctx, delegate).(string)
	if ctx.Function() != nil {
		return "go func " + expr + "()"
	}
	return "go " + expr
}
func (this *OgVisitor) VisitType_(ctx *parser.Type_Context, delegate antlr.ParseTreeVisitor) interface{} {
	return " " + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitTypeName(ctx *parser.TypeNameContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitTypeLit(ctx *parser.TypeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitArrayType(ctx *parser.ArrayTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitArrayLength(ctx *parser.ArrayLengthContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitElementType(ctx *parser.ElementTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitPointerType(ctx *parser.PointerTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "*" + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitInterfaceType(ctx *parser.InterfaceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	idx := ""
	methods := ""
	if ctx.IDENTIFIER() != nil {
		idx = ctx.IDENTIFIER().GetText() + " "
	}
	if len(ctx.AllMethodSpec()) != 0 {
		methods = this.VisitChildren(ctx, delegate).(string)
	}
	return idx + "interface" + "{\n" + methods + "}"
}
func (this *OgVisitor) VisitSliceType(ctx *parser.SliceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "[]" + this.VisitChildren(ctx, delegate).(string)
}
func (this *OgVisitor) VisitMapType(ctx *parser.MapTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitChannelType(ctx *parser.ChannelTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitMethodSpec(ctx *parser.MethodSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	idx := ""
	if ctx.IDENTIFIER() != nil {
		idx = ctx.IDENTIFIER().GetText()
	}
	return idx + this.VisitChildren(ctx, delegate).(string) + "\n"
}
func (this *OgVisitor) VisitFunctionType(ctx *parser.FunctionTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitSignature(ctx *parser.SignatureContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitResult(ctx *parser.ResultContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "(" + ctx.GetText() + ")"
}
func (this *OgVisitor) VisitParameters(ctx *parser.ParametersContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.ParameterList() == nil {
		return "()"
	}
	return "(" + this.VisitChildren(ctx, delegate).(string) + ")"
}
func (this *OgVisitor) VisitParameterList(ctx *parser.ParameterListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitParameterDecl(ctx *parser.ParameterDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate).(string) + ","
}
func (this *OgVisitor) VisitOperand(ctx *parser.OperandContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.Expression() != nil {
		return "(" + this.VisitChildren(ctx, delegate).(string) + ")"
	}
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitLiteral(ctx *parser.LiteralContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitBasicLit(ctx *parser.BasicLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitOperandName(ctx *parser.OperandNameContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitCompositeLit(ctx *parser.CompositeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitLiteralType(ctx *parser.LiteralTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitLiteralValue(ctx *parser.LiteralValueContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.ElementList() == nil {
		return "{}"
	}
	return "{\n" + this.VisitChildren(ctx, delegate).(string) + "\n}"
}
func (this *OgVisitor) VisitElementList(ctx *parser.ElementListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitKeyedElement(ctx *parser.KeyedElementContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := ""
	if ctx.Key() != nil {
		r += this.VisitKey(ctx.Key().(*parser.KeyContext), delegate).(string) + ":"
	}
	r += this.VisitElement(ctx.Element().(*parser.ElementContext), delegate).(string) + ",\n"
	return r
}
func (this *OgVisitor) VisitKey(ctx *parser.KeyContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.IDENTIFIER() != nil {
		return ctx.GetText()
	}
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitElement(ctx *parser.ElementContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitStructType(ctx *parser.StructTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	idx := ""
	res := ""
	if ctx.IDENTIFIER() != nil {
		idx = ctx.IDENTIFIER().GetText()
	}
	if len(ctx.AllFieldDecl()) != 0 {
		res = this.VisitChildren(ctx, delegate).(string)
	}
	return idx + " struct {\n" + res + "}"
}
func (this *OgVisitor) VisitFieldDecl(ctx *parser.FieldDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.IdentifierList() != nil {
		idList := this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).(string)
		type_ := this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(string)
		tag := ""
		if ctx.STRING_LIT() != nil {
			tag = ctx.STRING_LIT().GetText()
		}
		return idList + type_ + " " + tag + "\n"
	} else {
		return ctx.GetText()
	}
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitAnonymousField(ctx *parser.AnonymousFieldContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitFunctionLit(ctx *parser.FunctionLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitSelector(ctx *parser.SelectorContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "." + ctx.IDENTIFIER().GetText()
}
func (this *OgVisitor) VisitIndex(ctx *parser.IndexContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "[" + this.VisitChildren(ctx, delegate).(string) + "]"
}
func (this *OgVisitor) VisitSlice(ctx *parser.SliceContext, delegate antlr.ParseTreeVisitor) interface{} {
	txt := ctx.GetText()
	splited := strings.Split(txt, ":")
	res := "["
	i := 0
	if len(splited) == 2 {
		if splited[0] != "[" {
			res += this.VisitExpression(ctx.Expression(i).(*parser.ExpressionContext), delegate).(string)
			i++
		}
		res += ":"
		if splited[1] != "]" {
			res += this.VisitExpression(ctx.Expression(i).(*parser.ExpressionContext), delegate).(string)
			i++
		}
	}
	if len(splited) == 3 {
		if splited[0] != "[" {
			res += this.VisitExpression(ctx.Expression(i).(*parser.ExpressionContext), delegate).(string)
			i++
		}
		res += ":"
		if splited[1] != "" {
			res += this.VisitExpression(ctx.Expression(i).(*parser.ExpressionContext), delegate).(string)
			i++
		}
		res += ":"
		if splited[2] != "]" {
			res += this.VisitExpression(ctx.Expression(i).(*parser.ExpressionContext), delegate).(string)
			i++
		}
	}
	return res + "]"
}
func (this *OgVisitor) VisitTypeAssertion(ctx *parser.TypeAssertionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ".(" + this.VisitChildren(ctx, delegate).(string) + ")"
}
func (this *OgVisitor) VisitArguments(ctx *parser.ArgumentsContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetChildCount() == 2 {
		return "()"
	}
	return "(" + this.VisitChildren(ctx, delegate).(string) + ")"
}
func (this *OgVisitor) VisitMethodExpr(ctx *parser.MethodExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitReceiverType(ctx *parser.ReceiverTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitExpression(ctx *parser.ExpressionContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetChildCount() > 1 {
		exp1 := this.VisitExpression(ctx.Expression(0).(*parser.ExpressionContext), delegate).(string)
		op := this.VisitBinary_op(ctx.Binary_op().(*parser.Binary_opContext), delegate).(string)
		exp2 := this.VisitExpression(ctx.Expression(1).(*parser.ExpressionContext), delegate).(string)
		return exp1 + op + exp2
	}
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *OgVisitor) VisitUnary_op(ctx *parser.Unary_opContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitConversion(ctx *parser.ConversionContext, delegate antlr.ParseTreeVisitor) interface{} {
	t := this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(string)
	exp := this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
	return t + "(" + exp + ")"
}
func (this *OgVisitor) VisitEos(ctx *parser.EosContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
