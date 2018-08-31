// Code generated from Golang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Golang

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGolangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGolangVisitor) VisitSourceFile(ctx *SourceFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitPackageClause(ctx *PackageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitImportPath(ctx *ImportPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitConstSpec(ctx *ConstSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeDecl(ctx *TypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitFunctionDecl(ctx *FunctionDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitMethodDecl(ctx *MethodDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitReceiver(ctx *ReceiverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitVarSpec(ctx *VarSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitExpressionStmt(ctx *ExpressionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitSendStmt(ctx *SendStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitIncDecStmt(ctx *IncDecStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitAssign_op(ctx *Assign_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitShortVarDecl(ctx *ShortVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitEmptyStmt(ctx *EmptyStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitLabeledStmt(ctx *LabeledStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitGotoStmt(ctx *GotoStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitFallthroughStmt(ctx *FallthroughStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitDeferStmt(ctx *DeferStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitExprSwitchStmt(ctx *ExprSwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitExprCaseClause(ctx *ExprCaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitExprSwitchCase(ctx *ExprSwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeSwitchStmt(ctx *TypeSwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeSwitchGuard(ctx *TypeSwitchGuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeCaseClause(ctx *TypeCaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeSwitchCase(ctx *TypeSwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitSelectStmt(ctx *SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitCommClause(ctx *CommClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitCommCase(ctx *CommCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitRecvStmt(ctx *RecvStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitForClause(ctx *ForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitRangeClause(ctx *RangeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitGoStmt(ctx *GoStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeLit(ctx *TypeLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitArrayLength(ctx *ArrayLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitElementType(ctx *ElementTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitPointerType(ctx *PointerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitSliceType(ctx *SliceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitChannelType(ctx *ChannelTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitMethodSpec(ctx *MethodSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitSignature(ctx *SignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitResult(ctx *ResultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitParameterDecl(ctx *ParameterDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitBasicLit(ctx *BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitOperandName(ctx *OperandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitQualifiedIdent(ctx *QualifiedIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitCompositeLit(ctx *CompositeLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitLiteralType(ctx *LiteralTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitLiteralValue(ctx *LiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitElementList(ctx *ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitKeyedElement(ctx *KeyedElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitFieldDecl(ctx *FieldDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitAnonymousField(ctx *AnonymousFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitFunctionLit(ctx *FunctionLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitSlice(ctx *SliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitTypeAssertion(ctx *TypeAssertionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitMethodExpr(ctx *MethodExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitReceiverType(ctx *ReceiverTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitConversion(ctx *ConversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGolangVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}
