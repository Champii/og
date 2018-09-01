// Generated from ./parser/Og.g4 by ANTLR 4.7.1.

package parser // Og

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOgListener is a complete listener for a parse tree produced by OgParser.
type BaseOgListener struct{}

var _ OgListener = &BaseOgListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOgListener) VisitTerminal(node antlr.TerminalNode) {}
func (s *BaseOgListener) VisitErrorNode(node antlr.ErrorNode)   {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOgListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOgListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BaseOgListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BaseOgListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterPackageClause is called when production packageClause is entered.
func (s *BaseOgListener) EnterPackageClause(ctx *PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *BaseOgListener) ExitPackageClause(ctx *PackageClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseOgListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseOgListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseOgListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseOgListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *BaseOgListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *BaseOgListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterTopLevelDecl is called when production topLevelDecl is entered.
func (s *BaseOgListener) EnterTopLevelDecl(ctx *TopLevelDeclContext) {}

// ExitTopLevelDecl is called when production topLevelDecl is exited.
func (s *BaseOgListener) ExitTopLevelDecl(ctx *TopLevelDeclContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseOgListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseOgListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *BaseOgListener) EnterConstDecl(ctx *ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *BaseOgListener) ExitConstDecl(ctx *ConstDeclContext) {}

// EnterConstSpec is called when production constSpec is entered.
func (s *BaseOgListener) EnterConstSpec(ctx *ConstSpecContext) {}

// ExitConstSpec is called when production constSpec is exited.
func (s *BaseOgListener) ExitConstSpec(ctx *ConstSpecContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseOgListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseOgListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseOgListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseOgListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseOgListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseOgListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseOgListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseOgListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseOgListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseOgListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseOgListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseOgListener) ExitFunction(ctx *FunctionContext) {}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *BaseOgListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *BaseOgListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *BaseOgListener) EnterReceiver(ctx *ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *BaseOgListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseOgListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseOgListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarSpec is called when production varSpec is entered.
func (s *BaseOgListener) EnterVarSpec(ctx *VarSpecContext) {}

// ExitVarSpec is called when production varSpec is exited.
func (s *BaseOgListener) ExitVarSpec(ctx *VarSpecContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseOgListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseOgListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseOgListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseOgListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatementNoBlock is called when production statementNoBlock is entered.
func (s *BaseOgListener) EnterStatementNoBlock(ctx *StatementNoBlockContext) {}

// ExitStatementNoBlock is called when production statementNoBlock is exited.
func (s *BaseOgListener) ExitStatementNoBlock(ctx *StatementNoBlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseOgListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseOgListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseOgListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseOgListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterExpressionStmt is called when production expressionStmt is entered.
func (s *BaseOgListener) EnterExpressionStmt(ctx *ExpressionStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *BaseOgListener) ExitExpressionStmt(ctx *ExpressionStmtContext) {}

// EnterSendStmt is called when production sendStmt is entered.
func (s *BaseOgListener) EnterSendStmt(ctx *SendStmtContext) {}

// ExitSendStmt is called when production sendStmt is exited.
func (s *BaseOgListener) ExitSendStmt(ctx *SendStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *BaseOgListener) EnterIncDecStmt(ctx *IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *BaseOgListener) ExitIncDecStmt(ctx *IncDecStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseOgListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseOgListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *BaseOgListener) EnterAssign_op(ctx *Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *BaseOgListener) ExitAssign_op(ctx *Assign_opContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *BaseOgListener) EnterShortVarDecl(ctx *ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *BaseOgListener) ExitShortVarDecl(ctx *ShortVarDeclContext) {}

// EnterEmptyStmt is called when production emptyStmt is entered.
func (s *BaseOgListener) EnterEmptyStmt(ctx *EmptyStmtContext) {}

// ExitEmptyStmt is called when production emptyStmt is exited.
func (s *BaseOgListener) ExitEmptyStmt(ctx *EmptyStmtContext) {}

// EnterLabeledStmt is called when production labeledStmt is entered.
func (s *BaseOgListener) EnterLabeledStmt(ctx *LabeledStmtContext) {}

// ExitLabeledStmt is called when production labeledStmt is exited.
func (s *BaseOgListener) ExitLabeledStmt(ctx *LabeledStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseOgListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseOgListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseOgListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseOgListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseOgListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseOgListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterGotoStmt is called when production gotoStmt is entered.
func (s *BaseOgListener) EnterGotoStmt(ctx *GotoStmtContext) {}

// ExitGotoStmt is called when production gotoStmt is exited.
func (s *BaseOgListener) ExitGotoStmt(ctx *GotoStmtContext) {}

// EnterFallthroughStmt is called when production fallthroughStmt is entered.
func (s *BaseOgListener) EnterFallthroughStmt(ctx *FallthroughStmtContext) {}

// ExitFallthroughStmt is called when production fallthroughStmt is exited.
func (s *BaseOgListener) ExitFallthroughStmt(ctx *FallthroughStmtContext) {}

// EnterDeferStmt is called when production deferStmt is entered.
func (s *BaseOgListener) EnterDeferStmt(ctx *DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *BaseOgListener) ExitDeferStmt(ctx *DeferStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseOgListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseOgListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *BaseOgListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *BaseOgListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterExprSwitchStmt is called when production exprSwitchStmt is entered.
func (s *BaseOgListener) EnterExprSwitchStmt(ctx *ExprSwitchStmtContext) {}

// ExitExprSwitchStmt is called when production exprSwitchStmt is exited.
func (s *BaseOgListener) ExitExprSwitchStmt(ctx *ExprSwitchStmtContext) {}

// EnterExprCaseClause is called when production exprCaseClause is entered.
func (s *BaseOgListener) EnterExprCaseClause(ctx *ExprCaseClauseContext) {}

// ExitExprCaseClause is called when production exprCaseClause is exited.
func (s *BaseOgListener) ExitExprCaseClause(ctx *ExprCaseClauseContext) {}

// EnterExprSwitchCase is called when production exprSwitchCase is entered.
func (s *BaseOgListener) EnterExprSwitchCase(ctx *ExprSwitchCaseContext) {}

// ExitExprSwitchCase is called when production exprSwitchCase is exited.
func (s *BaseOgListener) ExitExprSwitchCase(ctx *ExprSwitchCaseContext) {}

// EnterTypeSwitchStmt is called when production typeSwitchStmt is entered.
func (s *BaseOgListener) EnterTypeSwitchStmt(ctx *TypeSwitchStmtContext) {}

// ExitTypeSwitchStmt is called when production typeSwitchStmt is exited.
func (s *BaseOgListener) ExitTypeSwitchStmt(ctx *TypeSwitchStmtContext) {}

// EnterTypeSwitchGuard is called when production typeSwitchGuard is entered.
func (s *BaseOgListener) EnterTypeSwitchGuard(ctx *TypeSwitchGuardContext) {}

// ExitTypeSwitchGuard is called when production typeSwitchGuard is exited.
func (s *BaseOgListener) ExitTypeSwitchGuard(ctx *TypeSwitchGuardContext) {}

// EnterTypeCaseClause is called when production typeCaseClause is entered.
func (s *BaseOgListener) EnterTypeCaseClause(ctx *TypeCaseClauseContext) {}

// ExitTypeCaseClause is called when production typeCaseClause is exited.
func (s *BaseOgListener) ExitTypeCaseClause(ctx *TypeCaseClauseContext) {}

// EnterTypeSwitchCase is called when production typeSwitchCase is entered.
func (s *BaseOgListener) EnterTypeSwitchCase(ctx *TypeSwitchCaseContext) {}

// ExitTypeSwitchCase is called when production typeSwitchCase is exited.
func (s *BaseOgListener) ExitTypeSwitchCase(ctx *TypeSwitchCaseContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseOgListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseOgListener) ExitTypeList(ctx *TypeListContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseOgListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseOgListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterCommClause is called when production commClause is entered.
func (s *BaseOgListener) EnterCommClause(ctx *CommClauseContext) {}

// ExitCommClause is called when production commClause is exited.
func (s *BaseOgListener) ExitCommClause(ctx *CommClauseContext) {}

// EnterCommCase is called when production commCase is entered.
func (s *BaseOgListener) EnterCommCase(ctx *CommCaseContext) {}

// ExitCommCase is called when production commCase is exited.
func (s *BaseOgListener) ExitCommCase(ctx *CommCaseContext) {}

// EnterRecvStmt is called when production recvStmt is entered.
func (s *BaseOgListener) EnterRecvStmt(ctx *RecvStmtContext) {}

// ExitRecvStmt is called when production recvStmt is exited.
func (s *BaseOgListener) ExitRecvStmt(ctx *RecvStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseOgListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseOgListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterForClause is called when production forClause is entered.
func (s *BaseOgListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *BaseOgListener) ExitForClause(ctx *ForClauseContext) {}

// EnterRangeClause is called when production rangeClause is entered.
func (s *BaseOgListener) EnterRangeClause(ctx *RangeClauseContext) {}

// ExitRangeClause is called when production rangeClause is exited.
func (s *BaseOgListener) ExitRangeClause(ctx *RangeClauseContext) {}

// EnterGoStmt is called when production goStmt is entered.
func (s *BaseOgListener) EnterGoStmt(ctx *GoStmtContext) {}

// ExitGoStmt is called when production goStmt is exited.
func (s *BaseOgListener) ExitGoStmt(ctx *GoStmtContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseOgListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseOgListener) ExitType_(ctx *Type_Context) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseOgListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseOgListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *BaseOgListener) EnterTypeLit(ctx *TypeLitContext) {}

// ExitTypeLit is called when production typeLit is exited.
func (s *BaseOgListener) ExitTypeLit(ctx *TypeLitContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseOgListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseOgListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterArrayLength is called when production arrayLength is entered.
func (s *BaseOgListener) EnterArrayLength(ctx *ArrayLengthContext) {}

// ExitArrayLength is called when production arrayLength is exited.
func (s *BaseOgListener) ExitArrayLength(ctx *ArrayLengthContext) {}

// EnterElementType is called when production elementType is entered.
func (s *BaseOgListener) EnterElementType(ctx *ElementTypeContext) {}

// ExitElementType is called when production elementType is exited.
func (s *BaseOgListener) ExitElementType(ctx *ElementTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BaseOgListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BaseOgListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BaseOgListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BaseOgListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterSliceType is called when production sliceType is entered.
func (s *BaseOgListener) EnterSliceType(ctx *SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *BaseOgListener) ExitSliceType(ctx *SliceTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseOgListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseOgListener) ExitMapType(ctx *MapTypeContext) {}

// EnterChannelType is called when production channelType is entered.
func (s *BaseOgListener) EnterChannelType(ctx *ChannelTypeContext) {}

// ExitChannelType is called when production channelType is exited.
func (s *BaseOgListener) ExitChannelType(ctx *ChannelTypeContext) {}

// EnterMethodSpec is called when production methodSpec is entered.
func (s *BaseOgListener) EnterMethodSpec(ctx *MethodSpecContext) {}

// ExitMethodSpec is called when production methodSpec is exited.
func (s *BaseOgListener) ExitMethodSpec(ctx *MethodSpecContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseOgListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseOgListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterSignature is called when production signature is entered.
func (s *BaseOgListener) EnterSignature(ctx *SignatureContext) {}

// ExitSignature is called when production signature is exited.
func (s *BaseOgListener) ExitSignature(ctx *SignatureContext) {}

// EnterResult is called when production result is entered.
func (s *BaseOgListener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *BaseOgListener) ExitResult(ctx *ResultContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseOgListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseOgListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseOgListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseOgListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *BaseOgListener) EnterParameterDecl(ctx *ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *BaseOgListener) ExitParameterDecl(ctx *ParameterDeclContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseOgListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseOgListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseOgListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseOgListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BaseOgListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BaseOgListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BaseOgListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BaseOgListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *BaseOgListener) EnterQualifiedIdent(ctx *QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *BaseOgListener) ExitQualifiedIdent(ctx *QualifiedIdentContext) {}

// EnterCompositeLit is called when production compositeLit is entered.
func (s *BaseOgListener) EnterCompositeLit(ctx *CompositeLitContext) {}

// ExitCompositeLit is called when production compositeLit is exited.
func (s *BaseOgListener) ExitCompositeLit(ctx *CompositeLitContext) {}

// EnterLiteralType is called when production literalType is entered.
func (s *BaseOgListener) EnterLiteralType(ctx *LiteralTypeContext) {}

// ExitLiteralType is called when production literalType is exited.
func (s *BaseOgListener) ExitLiteralType(ctx *LiteralTypeContext) {}

// EnterLiteralValue is called when production literalValue is entered.
func (s *BaseOgListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production literalValue is exited.
func (s *BaseOgListener) ExitLiteralValue(ctx *LiteralValueContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseOgListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseOgListener) ExitElementList(ctx *ElementListContext) {}

// EnterKeyedElement is called when production keyedElement is entered.
func (s *BaseOgListener) EnterKeyedElement(ctx *KeyedElementContext) {}

// ExitKeyedElement is called when production keyedElement is exited.
func (s *BaseOgListener) ExitKeyedElement(ctx *KeyedElementContext) {}

// EnterKey is called when production key is entered.
func (s *BaseOgListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseOgListener) ExitKey(ctx *KeyContext) {}

// EnterElement is called when production element is entered.
func (s *BaseOgListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseOgListener) ExitElement(ctx *ElementContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseOgListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseOgListener) ExitStructType(ctx *StructTypeContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BaseOgListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BaseOgListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterAnonymousField is called when production anonymousField is entered.
func (s *BaseOgListener) EnterAnonymousField(ctx *AnonymousFieldContext) {}

// ExitAnonymousField is called when production anonymousField is exited.
func (s *BaseOgListener) ExitAnonymousField(ctx *AnonymousFieldContext) {}

// EnterFunctionLit is called when production functionLit is entered.
func (s *BaseOgListener) EnterFunctionLit(ctx *FunctionLitContext) {}

// ExitFunctionLit is called when production functionLit is exited.
func (s *BaseOgListener) ExitFunctionLit(ctx *FunctionLitContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseOgListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseOgListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseOgListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseOgListener) ExitSelector(ctx *SelectorContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseOgListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseOgListener) ExitIndex(ctx *IndexContext) {}

// EnterSlice is called when production slice is entered.
func (s *BaseOgListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production slice is exited.
func (s *BaseOgListener) ExitSlice(ctx *SliceContext) {}

// EnterTypeAssertion is called when production typeAssertion is entered.
func (s *BaseOgListener) EnterTypeAssertion(ctx *TypeAssertionContext) {}

// ExitTypeAssertion is called when production typeAssertion is exited.
func (s *BaseOgListener) ExitTypeAssertion(ctx *TypeAssertionContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseOgListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseOgListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterMethodExpr is called when production methodExpr is entered.
func (s *BaseOgListener) EnterMethodExpr(ctx *MethodExprContext) {}

// ExitMethodExpr is called when production methodExpr is exited.
func (s *BaseOgListener) ExitMethodExpr(ctx *MethodExprContext) {}

// EnterReceiverType is called when production receiverType is entered.
func (s *BaseOgListener) EnterReceiverType(ctx *ReceiverTypeContext) {}

// ExitReceiverType is called when production receiverType is exited.
func (s *BaseOgListener) ExitReceiverType(ctx *ReceiverTypeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOgListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOgListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBinary_op is called when production binary_op is entered.
func (s *BaseOgListener) EnterBinary_op(ctx *Binary_opContext) {}

// ExitBinary_op is called when production binary_op is exited.
func (s *BaseOgListener) ExitBinary_op(ctx *Binary_opContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseOgListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseOgListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterUnary_op is called when production unary_op is entered.
func (s *BaseOgListener) EnterUnary_op(ctx *Unary_opContext) {}

// ExitUnary_op is called when production unary_op is exited.
func (s *BaseOgListener) ExitUnary_op(ctx *Unary_opContext) {}

// EnterConversion is called when production conversion is entered.
func (s *BaseOgListener) EnterConversion(ctx *ConversionContext) {}

// ExitConversion is called when production conversion is exited.
func (s *BaseOgListener) ExitConversion(ctx *ConversionContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseOgListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseOgListener) ExitEos(ctx *EosContext) {}
