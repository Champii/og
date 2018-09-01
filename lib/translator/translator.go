package translator

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"og/parser"
)

type GolangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (this *GolangVisitor) Aggregate(resultSoFar, childResult interface{}) interface{} {
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
func (this *GolangVisitor) VisitSourceFile(ctx *parser.SourceFileContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitPackageClause(ctx *parser.PackageClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "package " + ctx.IDENTIFIER().GetText() + "\n"
}
func (this *GolangVisitor) VisitImportDecl(ctx *parser.ImportDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "import (\n" + this.VisitChildren(ctx, delegate).(string) + ")\n"
}
func (this *GolangVisitor) VisitImportSpec(ctx *parser.ImportSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitImportPath(ctx *parser.ImportPathContext, delegate antlr.ParseTreeVisitor) interface{} {
	txt := ctx.GetText()
	if txt[0] == '"' {
		return txt + "\n"
	} else {
		return "\"" + txt + "\"\n"
	}
}
func (this *GolangVisitor) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitDeclaration(ctx *parser.DeclarationContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitConstDecl(ctx *parser.ConstDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitConstSpec(ctx *parser.ConstSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *GolangVisitor) VisitExpressionList(ctx *parser.ExpressionListContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
	if ctx.GetChildCount() > 1 {
		r += "," + this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(string)
	}
	return r
}
func (this *GolangVisitor) VisitTypeDecl(ctx *parser.TypeDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "type " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.IDENTIFIER().GetText() + " " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "\nfunc " + ctx.IDENTIFIER().GetText() + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitFunction(ctx *parser.FunctionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitMethodDecl(ctx *parser.MethodDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "\nfunc " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitReceiver(ctx *parser.ReceiverContext, delegate antlr.ParseTreeVisitor) interface{} {
	class := ctx.IDENTIFIER(0).GetText()
	method := ctx.IDENTIFIER(1).GetText()
	return "(this *" + class + ") " + method
}
func (this *GolangVisitor) VisitVarDecl(ctx *parser.VarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "var " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitVarSpec(ctx *parser.VarSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitBlock(ctx *parser.BlockContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "{\n" + this.VisitChildren(ctx, delegate).(string) + "}"
}
func (this *GolangVisitor) VisitStatementList(ctx *parser.StatementListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitStatement(ctx *parser.StatementContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate).(string) + "\n"
}
func (this *GolangVisitor) VisitStatementNoBlock(ctx *parser.StatementNoBlockContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "{" + this.VisitChildren(ctx, delegate).(string) + "}"
}
func (this *GolangVisitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitSendStmt(ctx *parser.SendStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitIncDecStmt(ctx *parser.IncDecStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *GolangVisitor) VisitAssignment(ctx *parser.AssignmentContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitAssign_op(ctx *parser.Assign_opContext, delegate antlr.ParseTreeVisitor) interface{} {
	if len(ctx.GetText()) == 1 {
		return "="
	}
	return ctx.GetText()
}
func (this *GolangVisitor) VisitBinary_op(ctx *parser.Binary_opContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *GolangVisitor) VisitShortVarDecl(ctx *parser.ShortVarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	idList := this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).(string)
	exprList := this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(string)
	return idList + " := " + exprList
}
func (this *GolangVisitor) VisitEmptyStmt(ctx *parser.EmptyStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitLabeledStmt(ctx *parser.LabeledStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "return " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitGotoStmt(ctx *parser.GotoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitFallthroughStmt(ctx *parser.FallthroughStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitDeferStmt(ctx *parser.DeferStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "defer " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitIfStmt(ctx *parser.IfStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
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
func (this *GolangVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitExprSwitchStmt(ctx *parser.ExprSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
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
func (this *GolangVisitor) VisitExprCaseClause(ctx *parser.ExprCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	sCase := this.VisitExprSwitchCase(ctx.ExprSwitchCase().(*parser.ExprSwitchCaseContext), delegate).(string)
	stmts := this.VisitStatementList(ctx.StatementList().(*parser.StatementListContext), delegate).(string)
	return sCase + ":" + stmts
}
func (this *GolangVisitor) VisitExprSwitchCase(ctx *parser.ExprSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetText() == "_" {
		return "default"
	}
	return "case " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitTypeSwitchStmt(ctx *parser.TypeSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := "switch "
	r += this.VisitTypeSwitchGuard(ctx.TypeSwitchGuard().(*parser.TypeSwitchGuardContext), delegate).(string)
	r += "{"
	for _, c := range ctx.AllTypeCaseClause() {
		r += this.VisitTypeCaseClause(c.(*parser.TypeCaseClauseContext), delegate).(string)
	}
	r += "}"
	return r
}
func (this *GolangVisitor) VisitTypeSwitchGuard(ctx *parser.TypeSwitchGuardContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := ""
	if ctx.IDENTIFIER() != nil {
		r += ctx.IDENTIFIER().GetText() + "="
	}
	expr := this.VisitPrimaryExpr(ctx.PrimaryExpr().(*parser.PrimaryExprContext), delegate).(string)
	return r + expr + ".(type)"
}
func (this *GolangVisitor) VisitTypeCaseClause(ctx *parser.TypeCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	sCase := this.VisitTypeSwitchCase(ctx.TypeSwitchCase().(*parser.TypeSwitchCaseContext), delegate).(string)
	stmts := this.VisitStatementList(ctx.StatementList().(*parser.StatementListContext), delegate).(string)
	return sCase + ":" + stmts
}
func (this *GolangVisitor) VisitTypeSwitchCase(ctx *parser.TypeSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetText() == "_" {
		return "default"
	}
	return "case " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitTypeList(ctx *parser.TypeListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitSelectStmt(ctx *parser.SelectStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitCommClause(ctx *parser.CommClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitCommCase(ctx *parser.CommCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitRecvStmt(ctx *parser.RecvStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitForStmt(ctx *parser.ForStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "for " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitForClause(ctx *parser.ForClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitRangeClause(ctx *parser.RangeClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := ""
	if ctx.IdentifierList() != nil {
		r = this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).(string)
	} else if ctx.ExpressionList() != nil {
		r = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(string)
	}
	return r + " := range " + this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
}
func (this *GolangVisitor) VisitGoStmt(ctx *parser.GoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	expr := this.VisitChildren(ctx, delegate).(string)
	if ctx.Function() != nil {
		return "go func " + expr + "()"
	}
	return "go " + expr
}
func (this *GolangVisitor) VisitType_(ctx *parser.Type_Context, delegate antlr.ParseTreeVisitor) interface{} {
	return " " + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitTypeName(ctx *parser.TypeNameContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *GolangVisitor) VisitTypeLit(ctx *parser.TypeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitArrayType(ctx *parser.ArrayTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitArrayLength(ctx *parser.ArrayLengthContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitElementType(ctx *parser.ElementTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitPointerType(ctx *parser.PointerTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "*" + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitInterfaceType(ctx *parser.InterfaceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "interface{}"
}
func (this *GolangVisitor) VisitSliceType(ctx *parser.SliceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "[]" + this.VisitChildren(ctx, delegate).(string)
}
func (this *GolangVisitor) VisitMapType(ctx *parser.MapTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitChannelType(ctx *parser.ChannelTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitMethodSpec(ctx *parser.MethodSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitFunctionType(ctx *parser.FunctionTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitSignature(ctx *parser.SignatureContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitResult(ctx *parser.ResultContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "(" + ctx.GetText() + ")"
}
func (this *GolangVisitor) VisitParameters(ctx *parser.ParametersContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.ParameterList() == nil {
		return "()"
	}
	return "(" + this.VisitChildren(ctx, delegate).(string) + ")"
}
func (this *GolangVisitor) VisitParameterList(ctx *parser.ParameterListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitParameterDecl(ctx *parser.ParameterDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate).(string) + ","
}
func (this *GolangVisitor) VisitOperand(ctx *parser.OperandContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.Expression() != nil {
		return "(" + this.VisitChildren(ctx, delegate).(string) + ")"
	}
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitLiteral(ctx *parser.LiteralContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitBasicLit(ctx *parser.BasicLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *GolangVisitor) VisitOperandName(ctx *parser.OperandNameContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *GolangVisitor) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitCompositeLit(ctx *parser.CompositeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitLiteralType(ctx *parser.LiteralTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitLiteralValue(ctx *parser.LiteralValueContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.ElementList() == nil {
		return "{}"
	}
	return "{\n" + this.VisitChildren(ctx, delegate).(string) + "\n}"
}
func (this *GolangVisitor) VisitElementList(ctx *parser.ElementListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitKeyedElement(ctx *parser.KeyedElementContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := ""
	if ctx.Key() != nil {
		r += this.VisitKey(ctx.Key().(*parser.KeyContext), delegate).(string) + ":"
	}
	r += this.VisitElement(ctx.Element().(*parser.ElementContext), delegate).(string) + ",\n"
	return r
}
func (this *GolangVisitor) VisitKey(ctx *parser.KeyContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.IDENTIFIER() != nil {
		return ctx.GetText()
	}
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitElement(ctx *parser.ElementContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitStructType(ctx *parser.StructTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "struct {\n" + this.VisitChildren(ctx, delegate).(string) + "}\n"
}
func (this *GolangVisitor) VisitFieldDecl(ctx *parser.FieldDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
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
func (this *GolangVisitor) VisitAnonymousField(ctx *parser.AnonymousFieldContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitFunctionLit(ctx *parser.FunctionLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitSelector(ctx *parser.SelectorContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "." + ctx.IDENTIFIER().GetText()
}
func (this *GolangVisitor) VisitIndex(ctx *parser.IndexContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "[" + this.VisitChildren(ctx, delegate).(string) + "]"
}
func (this *GolangVisitor) VisitSlice(ctx *parser.SliceContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitTypeAssertion(ctx *parser.TypeAssertionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ".(" + this.VisitChildren(ctx, delegate).(string) + ")"
}
func (this *GolangVisitor) VisitArguments(ctx *parser.ArgumentsContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetChildCount() == 2 {
		return "()"
	}
	return "(" + this.VisitChildren(ctx, delegate).(string) + ")"
}
func (this *GolangVisitor) VisitMethodExpr(ctx *parser.MethodExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitReceiverType(ctx *parser.ReceiverTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitExpression(ctx *parser.ExpressionContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.GetChildCount() > 1 {
		exp1 := this.VisitExpression(ctx.Expression(0).(*parser.ExpressionContext), delegate).(string)
		op := this.VisitBinary_op(ctx.Binary_op().(*parser.Binary_opContext), delegate).(string)
		exp2 := this.VisitExpression(ctx.Expression(1).(*parser.ExpressionContext), delegate).(string)
		return exp1 + op + exp2
	}
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}
func (this *GolangVisitor) VisitUnary_op(ctx *parser.Unary_opContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *GolangVisitor) VisitConversion(ctx *parser.ConversionContext, delegate antlr.ParseTreeVisitor) interface{} {
	t := this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(string)
	exp := this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
	return t + "(" + exp + ")"
}
func (this *GolangVisitor) VisitEos(ctx *parser.EosContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitChildren(ctx, delegate)
}

