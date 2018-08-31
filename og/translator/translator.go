// Generated from Golang.g4 by ANTLR 4.7.1.

package translator // Golang
import (
	"Og/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type GolangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

// func (v *GolangVisitor) Init() interface{} {
// 	return nil
// }

// func (v *GolangVisitor) VisitNext(node antlr.Tree, resultSoFar interface{}) bool {
// 	return true
// }

func (v *GolangVisitor) Aggregate(resultSoFar, childResult interface{}) interface{} {
	switch childResult.(type) {
	case string:
		switch resultSoFar.(type) {
		case string:
			return resultSoFar.(string) + childResult.(string)
		default:
			return childResult
		}
	default:
		return resultSoFar
	}

	return nil
}

// func (v *GolangVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
// 	return nil
// }

// func (v *GolangVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
// 	return nil
// }

func (v *GolangVisitor) VisitSourceFile(ctx *parser.SourceFileContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitPackageClause(ctx *parser.PackageClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "package " + ctx.IDENTIFIER().GetText() + "\n"
}

func (v *GolangVisitor) VisitImportDecl(ctx *parser.ImportDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "import (\n" + v.VisitChildren(ctx, delegate).(string) + ")\n"
}

func (v *GolangVisitor) VisitImportSpec(ctx *parser.ImportSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitImportPath(ctx *parser.ImportPathContext, delegate antlr.ParseTreeVisitor) interface{} {
	txt := ctx.GetText()

	if txt[0] == '"' {
		return txt + "\n"
	} else {
		return "\"" + txt + "\"\n"
	}
}

func (v *GolangVisitor) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitDeclaration(ctx *parser.DeclarationContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitConstDecl(ctx *parser.ConstDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitConstSpec(ctx *parser.ConstSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}

func (v *GolangVisitor) VisitExpressionList(ctx *parser.ExpressionListContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := ctx.Expression().GetText()

	if ctx.GetChildCount() > 1 {
		r += "," + v.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(string)
	}

	return r
}

func (v *GolangVisitor) VisitTypeDecl(ctx *parser.TypeDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "type " + v.VisitChildren(ctx, delegate).(string)
}

func (v *GolangVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	// TODO: allow for multiple types declarations
	return ctx.IDENTIFIER().GetText() + " " + v.VisitChildren(ctx, delegate).(string)
}

func (v *GolangVisitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "\nfunc " + ctx.IDENTIFIER().GetText() + v.VisitChildren(ctx, delegate).(string)
}

func (v *GolangVisitor) VisitFunction(ctx *parser.FunctionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitMethodDecl(ctx *parser.MethodDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitReceiver(ctx *parser.ReceiverContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitVarDecl(ctx *parser.VarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitVarSpec(ctx *parser.VarSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitBlock(ctx *parser.BlockContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "{\n" + v.VisitChildren(ctx, delegate).(string) + "}"
	// return "{}\n"
}

func (v *GolangVisitor) VisitStatementList(ctx *parser.StatementListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitStatement(ctx *parser.StatementContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate).(string) + "\n"
}

func (v *GolangVisitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitSendStmt(ctx *parser.SendStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitIncDecStmt(ctx *parser.IncDecStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitAssignment(ctx *parser.AssignmentContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitAssign_op(ctx *parser.Assign_opContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitShortVarDecl(ctx *parser.ShortVarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	idList := v.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).(string)

	exprList := v.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(string)

	return idList + " := " + exprList
}

func (v *GolangVisitor) VisitEmptyStmt(ctx *parser.EmptyStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitLabeledStmt(ctx *parser.LabeledStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "return " + v.VisitChildren(ctx, delegate).(string)
}

func (v *GolangVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitGotoStmt(ctx *parser.GotoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitFallthroughStmt(ctx *parser.FallthroughStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitDeferStmt(ctx *parser.DeferStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitIfStmt(ctx *parser.IfStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	r := "if "

	if ctx.SimpleStmt() != nil {
		r += v.VisitSimpleStmt(ctx.SimpleStmt().(*parser.SimpleStmtContext), delegate).(string) + ";"
	}

	r += v.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(string)
	r += v.VisitBlock(ctx.Block(0).(*parser.BlockContext), delegate).(string)

	if ctx.Block(1) != nil {
		r += "else " + v.VisitBlock(ctx.Block(1).(*parser.BlockContext), delegate).(string)
	} else if ctx.IfStmt() != nil {
		r += "else " + v.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext), delegate).(string)
	}

	return r
}

func (v *GolangVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitExprSwitchStmt(ctx *parser.ExprSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitExprCaseClause(ctx *parser.ExprCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitExprSwitchCase(ctx *parser.ExprSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitTypeSwitchStmt(ctx *parser.TypeSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitTypeSwitchGuard(ctx *parser.TypeSwitchGuardContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitTypeCaseClause(ctx *parser.TypeCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitTypeSwitchCase(ctx *parser.TypeSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitTypeList(ctx *parser.TypeListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitSelectStmt(ctx *parser.SelectStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitCommClause(ctx *parser.CommClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitCommCase(ctx *parser.CommCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitRecvStmt(ctx *parser.RecvStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitForStmt(ctx *parser.ForStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitForClause(ctx *parser.ForClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitRangeClause(ctx *parser.RangeClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitGoStmt(ctx *parser.GoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitType_(ctx *parser.Type_Context, delegate antlr.ParseTreeVisitor) interface{} {
	return " " + v.VisitChildren(ctx, delegate).(string)
}

func (v *GolangVisitor) VisitTypeName(ctx *parser.TypeNameContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}

func (v *GolangVisitor) VisitTypeLit(ctx *parser.TypeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitArrayType(ctx *parser.ArrayTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitArrayLength(ctx *parser.ArrayLengthContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitElementType(ctx *parser.ElementTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitPointerType(ctx *parser.PointerTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "*" + v.VisitChildren(ctx, delegate).(string)
}

func (v *GolangVisitor) VisitInterfaceType(ctx *parser.InterfaceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitSliceType(ctx *parser.SliceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "[]" + v.VisitChildren(ctx, delegate).(string)
}

func (v *GolangVisitor) VisitMapType(ctx *parser.MapTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitChannelType(ctx *parser.ChannelTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitMethodSpec(ctx *parser.MethodSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitFunctionType(ctx *parser.FunctionTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitSignature(ctx *parser.SignatureContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitResult(ctx *parser.ResultContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "(" + ctx.GetText() + ")"
}

func (v *GolangVisitor) VisitParameters(ctx *parser.ParametersContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.ParameterList() == nil {
		return "()"
	}

	return "(" + v.VisitChildren(ctx, delegate).(string) + ")"
}

func (v *GolangVisitor) VisitParameterList(ctx *parser.ParameterListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitParameterDecl(ctx *parser.ParameterDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate).(string) + ","
}

func (v *GolangVisitor) VisitOperand(ctx *parser.OperandContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}

func (v *GolangVisitor) VisitLiteral(ctx *parser.LiteralContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitBasicLit(ctx *parser.BasicLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitOperandName(ctx *parser.OperandNameContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}

func (v *GolangVisitor) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitCompositeLit(ctx *parser.CompositeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitLiteralType(ctx *parser.LiteralTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitLiteralValue(ctx *parser.LiteralValueContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitElementList(ctx *parser.ElementListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitKeyedElement(ctx *parser.KeyedElementContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitKey(ctx *parser.KeyContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitElement(ctx *parser.ElementContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitStructType(ctx *parser.StructTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "struct {\n" + v.VisitChildren(ctx, delegate).(string) + "}\n"
}

func (v *GolangVisitor) VisitFieldDecl(ctx *parser.FieldDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.IdentifierList() != nil {
		idList := v.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).(string)
		type_ := v.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(string)

		tag := ""

		if ctx.STRING_LIT() != nil {
			tag = ctx.STRING_LIT().GetText()
		}

		return idList + type_ + " " + tag + "\n"
	} else {
		return ctx.GetText()
	}

	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitAnonymousField(ctx *parser.AnonymousFieldContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitFunctionLit(ctx *parser.FunctionLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitSelector(ctx *parser.SelectorContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "." + ctx.IDENTIFIER().GetText()
}

func (v *GolangVisitor) VisitIndex(ctx *parser.IndexContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitSlice(ctx *parser.SliceContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitTypeAssertion(ctx *parser.TypeAssertionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitArguments(ctx *parser.ArgumentsContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "(" + v.VisitChildren(ctx, delegate).(string) + ")"
}

func (v *GolangVisitor) VisitMethodExpr(ctx *parser.MethodExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitReceiverType(ctx *parser.ReceiverTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitExpression(ctx *parser.ExpressionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}

func (v *GolangVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitConversion(ctx *parser.ConversionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}

func (v *GolangVisitor) VisitEos(ctx *parser.EosContext, delegate antlr.ParseTreeVisitor) interface{} {
	return v.VisitChildren(ctx, delegate)
}
