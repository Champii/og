// Code generated from Golang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Golang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GolangParser.
type GolangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GolangParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by GolangParser#packageClause.
	VisitPackageClause(ctx *PackageClauseContext) interface{}

	// Visit a parse tree produced by GolangParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by GolangParser#importPath.
	VisitImportPath(ctx *ImportPathContext) interface{}

	// Visit a parse tree produced by GolangParser#topLevelDecl.
	VisitTopLevelDecl(ctx *TopLevelDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by GolangParser#constDecl.
	VisitConstDecl(ctx *ConstDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#constSpec.
	VisitConstSpec(ctx *ConstSpecContext) interface{}

	// Visit a parse tree produced by GolangParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by GolangParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by GolangParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by GolangParser#functionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by GolangParser#methodDecl.
	VisitMethodDecl(ctx *MethodDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#receiver.
	VisitReceiver(ctx *ReceiverContext) interface{}

	// Visit a parse tree produced by GolangParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#varSpec.
	VisitVarSpec(ctx *VarSpecContext) interface{}

	// Visit a parse tree produced by GolangParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GolangParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by GolangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by GolangParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#expressionStmt.
	VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#sendStmt.
	VisitSendStmt(ctx *SendStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#incDecStmt.
	VisitIncDecStmt(ctx *IncDecStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by GolangParser#assign_op.
	VisitAssign_op(ctx *Assign_opContext) interface{}

	// Visit a parse tree produced by GolangParser#shortVarDecl.
	VisitShortVarDecl(ctx *ShortVarDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#emptyStmt.
	VisitEmptyStmt(ctx *EmptyStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#labeledStmt.
	VisitLabeledStmt(ctx *LabeledStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#gotoStmt.
	VisitGotoStmt(ctx *GotoStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#fallthroughStmt.
	VisitFallthroughStmt(ctx *FallthroughStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#deferStmt.
	VisitDeferStmt(ctx *DeferStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#switchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#exprSwitchStmt.
	VisitExprSwitchStmt(ctx *ExprSwitchStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#exprCaseClause.
	VisitExprCaseClause(ctx *ExprCaseClauseContext) interface{}

	// Visit a parse tree produced by GolangParser#exprSwitchCase.
	VisitExprSwitchCase(ctx *ExprSwitchCaseContext) interface{}

	// Visit a parse tree produced by GolangParser#typeSwitchStmt.
	VisitTypeSwitchStmt(ctx *TypeSwitchStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#typeSwitchGuard.
	VisitTypeSwitchGuard(ctx *TypeSwitchGuardContext) interface{}

	// Visit a parse tree produced by GolangParser#typeCaseClause.
	VisitTypeCaseClause(ctx *TypeCaseClauseContext) interface{}

	// Visit a parse tree produced by GolangParser#typeSwitchCase.
	VisitTypeSwitchCase(ctx *TypeSwitchCaseContext) interface{}

	// Visit a parse tree produced by GolangParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by GolangParser#selectStmt.
	VisitSelectStmt(ctx *SelectStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#commClause.
	VisitCommClause(ctx *CommClauseContext) interface{}

	// Visit a parse tree produced by GolangParser#commCase.
	VisitCommCase(ctx *CommCaseContext) interface{}

	// Visit a parse tree produced by GolangParser#recvStmt.
	VisitRecvStmt(ctx *RecvStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#forClause.
	VisitForClause(ctx *ForClauseContext) interface{}

	// Visit a parse tree produced by GolangParser#rangeClause.
	VisitRangeClause(ctx *RangeClauseContext) interface{}

	// Visit a parse tree produced by GolangParser#goStmt.
	VisitGoStmt(ctx *GoStmtContext) interface{}

	// Visit a parse tree produced by GolangParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by GolangParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by GolangParser#typeLit.
	VisitTypeLit(ctx *TypeLitContext) interface{}

	// Visit a parse tree produced by GolangParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#arrayLength.
	VisitArrayLength(ctx *ArrayLengthContext) interface{}

	// Visit a parse tree produced by GolangParser#elementType.
	VisitElementType(ctx *ElementTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#pointerType.
	VisitPointerType(ctx *PointerTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#interfaceType.
	VisitInterfaceType(ctx *InterfaceTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#sliceType.
	VisitSliceType(ctx *SliceTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#channelType.
	VisitChannelType(ctx *ChannelTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#methodSpec.
	VisitMethodSpec(ctx *MethodSpecContext) interface{}

	// Visit a parse tree produced by GolangParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#signature.
	VisitSignature(ctx *SignatureContext) interface{}

	// Visit a parse tree produced by GolangParser#result.
	VisitResult(ctx *ResultContext) interface{}

	// Visit a parse tree produced by GolangParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by GolangParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by GolangParser#parameterDecl.
	VisitParameterDecl(ctx *ParameterDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by GolangParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by GolangParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by GolangParser#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by GolangParser#qualifiedIdent.
	VisitQualifiedIdent(ctx *QualifiedIdentContext) interface{}

	// Visit a parse tree produced by GolangParser#compositeLit.
	VisitCompositeLit(ctx *CompositeLitContext) interface{}

	// Visit a parse tree produced by GolangParser#literalType.
	VisitLiteralType(ctx *LiteralTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#literalValue.
	VisitLiteralValue(ctx *LiteralValueContext) interface{}

	// Visit a parse tree produced by GolangParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by GolangParser#keyedElement.
	VisitKeyedElement(ctx *KeyedElementContext) interface{}

	// Visit a parse tree produced by GolangParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by GolangParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by GolangParser#structType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#fieldDecl.
	VisitFieldDecl(ctx *FieldDeclContext) interface{}

	// Visit a parse tree produced by GolangParser#anonymousField.
	VisitAnonymousField(ctx *AnonymousFieldContext) interface{}

	// Visit a parse tree produced by GolangParser#functionLit.
	VisitFunctionLit(ctx *FunctionLitContext) interface{}

	// Visit a parse tree produced by GolangParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by GolangParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by GolangParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by GolangParser#slice.
	VisitSlice(ctx *SliceContext) interface{}

	// Visit a parse tree produced by GolangParser#typeAssertion.
	VisitTypeAssertion(ctx *TypeAssertionContext) interface{}

	// Visit a parse tree produced by GolangParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by GolangParser#methodExpr.
	VisitMethodExpr(ctx *MethodExprContext) interface{}

	// Visit a parse tree produced by GolangParser#receiverType.
	VisitReceiverType(ctx *ReceiverTypeContext) interface{}

	// Visit a parse tree produced by GolangParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GolangParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by GolangParser#conversion.
	VisitConversion(ctx *ConversionContext) interface{}

	// Visit a parse tree produced by GolangParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
