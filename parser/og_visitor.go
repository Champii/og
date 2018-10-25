// Generated from ./parser/Og.g4 by ANTLR 4.7.1.

package parser // Og

import "github.com/champii/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by OgParser.
type OgVisitor interface {
	antlr.ParseTreeVisitor
}

type SourceFileVisitor interface {
	VisitSourceFile(ctx *SourceFileContext, delegate antlr.ParseTreeVisitor) interface{}
}
type InterpVisitor interface {
	VisitInterp(ctx *InterpContext, delegate antlr.ParseTreeVisitor) interface{}
}
type PackageClauseVisitor interface {
	VisitPackageClause(ctx *PackageClauseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ImportDeclVisitor interface {
	VisitImportDecl(ctx *ImportDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ImportBodyVisitor interface {
	VisitImportBody(ctx *ImportBodyContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ImportSpecVisitor interface {
	VisitImportSpec(ctx *ImportSpecContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ImportPathVisitor interface {
	VisitImportPath(ctx *ImportPathContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TopLevelDeclVisitor interface {
	VisitTopLevelDecl(ctx *TopLevelDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type DeclarationVisitor interface {
	VisitDeclaration(ctx *DeclarationContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ConstDeclVisitor interface {
	VisitConstDecl(ctx *ConstDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ConstSpecVisitor interface {
	VisitConstSpec(ctx *ConstSpecContext, delegate antlr.ParseTreeVisitor) interface{}
}
type IdentifierListVisitor interface {
	VisitIdentifierList(ctx *IdentifierListContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ExpressionListVisitor interface {
	VisitExpressionList(ctx *ExpressionListContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeDeclVisitor interface {
	VisitTypeDecl(ctx *TypeDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeSpecVisitor interface {
	VisitTypeSpec(ctx *TypeSpecContext, delegate antlr.ParseTreeVisitor) interface{}
}
type FunctionDeclVisitor interface {
	VisitFunctionDecl(ctx *FunctionDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type FunctionVisitor interface {
	VisitFunction(ctx *FunctionContext, delegate antlr.ParseTreeVisitor) interface{}
}
type MethodDeclVisitor interface {
	VisitMethodDecl(ctx *MethodDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ReceiverVisitor interface {
	VisitReceiver(ctx *ReceiverContext, delegate antlr.ParseTreeVisitor) interface{}
}
type VarDeclVisitor interface {
	VisitVarDecl(ctx *VarDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type VarSpecVisitor interface {
	VisitVarSpec(ctx *VarSpecContext, delegate antlr.ParseTreeVisitor) interface{}
}
type BlockVisitor interface {
	VisitBlock(ctx *BlockContext, delegate antlr.ParseTreeVisitor) interface{}
}
type StatementListVisitor interface {
	VisitStatementList(ctx *StatementListContext, delegate antlr.ParseTreeVisitor) interface{}
}
type StatementVisitor interface {
	VisitStatement(ctx *StatementContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SimpleStmtVisitor interface {
	VisitSimpleStmt(ctx *SimpleStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SendStmtVisitor interface {
	VisitSendStmt(ctx *SendStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type IncDecStmtVisitor interface {
	VisitIncDecStmt(ctx *IncDecStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type AssignmentVisitor interface {
	VisitAssignment(ctx *AssignmentContext, delegate antlr.ParseTreeVisitor) interface{}
}
type Assign_opVisitor interface {
	VisitAssign_op(ctx *Assign_opContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ShortVarDeclVisitor interface {
	VisitShortVarDecl(ctx *ShortVarDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type EmptyStmtVisitor interface {
	VisitEmptyStmt(ctx *EmptyStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type LabeledStmtVisitor interface {
	VisitLabeledStmt(ctx *LabeledStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ReturnStmtVisitor interface {
	VisitReturnStmt(ctx *ReturnStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type BreakStmtVisitor interface {
	VisitBreakStmt(ctx *BreakStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ContinueStmtVisitor interface {
	VisitContinueStmt(ctx *ContinueStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type GotoStmtVisitor interface {
	VisitGotoStmt(ctx *GotoStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type FallthroughStmtVisitor interface {
	VisitFallthroughStmt(ctx *FallthroughStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type DeferStmtVisitor interface {
	VisitDeferStmt(ctx *DeferStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type IfStmtVisitor interface {
	VisitIfStmt(ctx *IfStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SwitchStmtVisitor interface {
	VisitSwitchStmt(ctx *SwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ExprSwitchStmtVisitor interface {
	VisitExprSwitchStmt(ctx *ExprSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ExprCaseClauseVisitor interface {
	VisitExprCaseClause(ctx *ExprCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ExprSwitchCaseVisitor interface {
	VisitExprSwitchCase(ctx *ExprSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeSwitchStmtVisitor interface {
	VisitTypeSwitchStmt(ctx *TypeSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeSwitchGuardVisitor interface {
	VisitTypeSwitchGuard(ctx *TypeSwitchGuardContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeCaseClauseVisitor interface {
	VisitTypeCaseClause(ctx *TypeCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeSwitchCaseVisitor interface {
	VisitTypeSwitchCase(ctx *TypeSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeListVisitor interface {
	VisitTypeList(ctx *TypeListContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SelectStmtVisitor interface {
	VisitSelectStmt(ctx *SelectStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type CommClauseVisitor interface {
	VisitCommClause(ctx *CommClauseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type CommCaseVisitor interface {
	VisitCommCase(ctx *CommCaseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type RecvStmtVisitor interface {
	VisitRecvStmt(ctx *RecvStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ForStmtVisitor interface {
	VisitForStmt(ctx *ForStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ForClauseVisitor interface {
	VisitForClause(ctx *ForClauseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type RangeClauseVisitor interface {
	VisitRangeClause(ctx *RangeClauseContext, delegate antlr.ParseTreeVisitor) interface{}
}
type GoStmtVisitor interface {
	VisitGoStmt(ctx *GoStmtContext, delegate antlr.ParseTreeVisitor) interface{}
}
type Type_Visitor interface {
	VisitType_(ctx *Type_Context, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeNameVisitor interface {
	VisitTypeName(ctx *TypeNameContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeLitVisitor interface {
	VisitTypeLit(ctx *TypeLitContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ArrayTypeVisitor interface {
	VisitArrayType(ctx *ArrayTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ArrayLengthVisitor interface {
	VisitArrayLength(ctx *ArrayLengthContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ElementTypeVisitor interface {
	VisitElementType(ctx *ElementTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type PointerTypeVisitor interface {
	VisitPointerType(ctx *PointerTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type InterfaceTypeVisitor interface {
	VisitInterfaceType(ctx *InterfaceTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SliceTypeVisitor interface {
	VisitSliceType(ctx *SliceTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type MapTypeVisitor interface {
	VisitMapType(ctx *MapTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ChannelTypeVisitor interface {
	VisitChannelType(ctx *ChannelTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ChannelDeclVisitor interface {
	VisitChannelDecl(ctx *ChannelDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type MethodSpecVisitor interface {
	VisitMethodSpec(ctx *MethodSpecContext, delegate antlr.ParseTreeVisitor) interface{}
}
type FunctionTypeVisitor interface {
	VisitFunctionType(ctx *FunctionTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SignatureVisitor interface {
	VisitSignature(ctx *SignatureContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TemplateSpecVisitor interface {
	VisitTemplateSpec(ctx *TemplateSpecContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ResultVisitor interface {
	VisitResult(ctx *ResultContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ParametersVisitor interface {
	VisitParameters(ctx *ParametersContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ParameterListVisitor interface {
	VisitParameterList(ctx *ParameterListContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ParameterDeclVisitor interface {
	VisitParameterDecl(ctx *ParameterDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type RestOpVisitor interface {
	VisitRestOp(ctx *RestOpContext, delegate antlr.ParseTreeVisitor) interface{}
}
type OperandVisitor interface {
	VisitOperand(ctx *OperandContext, delegate antlr.ParseTreeVisitor) interface{}
}
type LiteralVisitor interface {
	VisitLiteral(ctx *LiteralContext, delegate antlr.ParseTreeVisitor) interface{}
}
type BasicLitVisitor interface {
	VisitBasicLit(ctx *BasicLitContext, delegate antlr.ParseTreeVisitor) interface{}
}
type OperandNameVisitor interface {
	VisitOperandName(ctx *OperandNameContext, delegate antlr.ParseTreeVisitor) interface{}
}
type This_Visitor interface {
	VisitThis_(ctx *This_Context, delegate antlr.ParseTreeVisitor) interface{}
}
type QualifiedIdentVisitor interface {
	VisitQualifiedIdent(ctx *QualifiedIdentContext, delegate antlr.ParseTreeVisitor) interface{}
}
type CompositeLitVisitor interface {
	VisitCompositeLit(ctx *CompositeLitContext, delegate antlr.ParseTreeVisitor) interface{}
}
type LiteralTypeVisitor interface {
	VisitLiteralType(ctx *LiteralTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type LiteralValueVisitor interface {
	VisitLiteralValue(ctx *LiteralValueContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ElementListVisitor interface {
	VisitElementList(ctx *ElementListContext, delegate antlr.ParseTreeVisitor) interface{}
}
type KeyedElementVisitor interface {
	VisitKeyedElement(ctx *KeyedElementContext, delegate antlr.ParseTreeVisitor) interface{}
}
type KeyVisitor interface {
	VisitKey(ctx *KeyContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ElementVisitor interface {
	VisitElement(ctx *ElementContext, delegate antlr.ParseTreeVisitor) interface{}
}
type StructTypeVisitor interface {
	VisitStructType(ctx *StructTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type FieldDeclVisitor interface {
	VisitFieldDecl(ctx *FieldDeclContext, delegate antlr.ParseTreeVisitor) interface{}
}
type InlineStructMethodVisitor interface {
	VisitInlineStructMethod(ctx *InlineStructMethodContext, delegate antlr.ParseTreeVisitor) interface{}
}
type AnonymousFieldVisitor interface {
	VisitAnonymousField(ctx *AnonymousFieldContext, delegate antlr.ParseTreeVisitor) interface{}
}
type FunctionLitVisitor interface {
	VisitFunctionLit(ctx *FunctionLitContext, delegate antlr.ParseTreeVisitor) interface{}
}
type PrimaryExprVisitor interface {
	VisitPrimaryExpr(ctx *PrimaryExprContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SecondaryExprVisitor interface {
	VisitSecondaryExpr(ctx *SecondaryExprContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SelectorVisitor interface {
	VisitSelector(ctx *SelectorContext, delegate antlr.ParseTreeVisitor) interface{}
}
type IndexVisitor interface {
	VisitIndex(ctx *IndexContext, delegate antlr.ParseTreeVisitor) interface{}
}
type SliceVisitor interface {
	VisitSlice(ctx *SliceContext, delegate antlr.ParseTreeVisitor) interface{}
}
type TypeAssertionVisitor interface {
	VisitTypeAssertion(ctx *TypeAssertionContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ArgumentsVisitor interface {
	VisitArguments(ctx *ArgumentsContext, delegate antlr.ParseTreeVisitor) interface{}
}
type MethodExprVisitor interface {
	VisitMethodExpr(ctx *MethodExprContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ReceiverTypeVisitor interface {
	VisitReceiverType(ctx *ReceiverTypeContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ExpressionVisitor interface {
	VisitExpression(ctx *ExpressionContext, delegate antlr.ParseTreeVisitor) interface{}
}
type UnaryExprVisitor interface {
	VisitUnaryExpr(ctx *UnaryExprContext, delegate antlr.ParseTreeVisitor) interface{}
}
type ConversionVisitor interface {
	VisitConversion(ctx *ConversionContext, delegate antlr.ParseTreeVisitor) interface{}
}
type EosVisitor interface {
	VisitEos(ctx *EosContext, delegate antlr.ParseTreeVisitor) interface{}
}
