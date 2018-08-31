// Generated from Golang.g4 by ANTLR 4.7.1.

package parser // Golang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGolangListener is a complete listener for a parse tree produced by GolangParser.
type BaseGolangListener struct{}

var _ GolangListener = &BaseGolangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGolangListener) VisitTerminal(node antlr.TerminalNode) {}
func (s *BaseGolangListener) VisitErrorNode(node antlr.ErrorNode)   {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGolangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGolangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BaseGolangListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BaseGolangListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterPackageClause is called when production packageClause is entered.
func (s *BaseGolangListener) EnterPackageClause(ctx *PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *BaseGolangListener) ExitPackageClause(ctx *PackageClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseGolangListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseGolangListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseGolangListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseGolangListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *BaseGolangListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *BaseGolangListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterTopLevelDecl is called when production topLevelDecl is entered.
func (s *BaseGolangListener) EnterTopLevelDecl(ctx *TopLevelDeclContext) {}

// ExitTopLevelDecl is called when production topLevelDecl is exited.
func (s *BaseGolangListener) ExitTopLevelDecl(ctx *TopLevelDeclContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseGolangListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseGolangListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *BaseGolangListener) EnterConstDecl(ctx *ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *BaseGolangListener) ExitConstDecl(ctx *ConstDeclContext) {}

// EnterConstSpec is called when production constSpec is entered.
func (s *BaseGolangListener) EnterConstSpec(ctx *ConstSpecContext) {}

// ExitConstSpec is called when production constSpec is exited.
func (s *BaseGolangListener) ExitConstSpec(ctx *ConstSpecContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseGolangListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseGolangListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseGolangListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseGolangListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseGolangListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseGolangListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseGolangListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseGolangListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseGolangListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseGolangListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseGolangListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseGolangListener) ExitFunction(ctx *FunctionContext) {}

// EnterMethodDecl is called when production methodDecl is entered.
func (s *BaseGolangListener) EnterMethodDecl(ctx *MethodDeclContext) {}

// ExitMethodDecl is called when production methodDecl is exited.
func (s *BaseGolangListener) ExitMethodDecl(ctx *MethodDeclContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *BaseGolangListener) EnterReceiver(ctx *ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *BaseGolangListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseGolangListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseGolangListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarSpec is called when production varSpec is entered.
func (s *BaseGolangListener) EnterVarSpec(ctx *VarSpecContext) {}

// ExitVarSpec is called when production varSpec is exited.
func (s *BaseGolangListener) ExitVarSpec(ctx *VarSpecContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGolangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGolangListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseGolangListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseGolangListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatementNoBlock is called when production statementNoBlock is entered.
func (s *BaseGolangListener) EnterStatementNoBlock(ctx *StatementNoBlockContext) {}

// ExitStatementNoBlock is called when production statementNoBlock is exited.
func (s *BaseGolangListener) ExitStatementNoBlock(ctx *StatementNoBlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGolangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGolangListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseGolangListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseGolangListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterExpressionStmt is called when production expressionStmt is entered.
func (s *BaseGolangListener) EnterExpressionStmt(ctx *ExpressionStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *BaseGolangListener) ExitExpressionStmt(ctx *ExpressionStmtContext) {}

// EnterSendStmt is called when production sendStmt is entered.
func (s *BaseGolangListener) EnterSendStmt(ctx *SendStmtContext) {}

// ExitSendStmt is called when production sendStmt is exited.
func (s *BaseGolangListener) ExitSendStmt(ctx *SendStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *BaseGolangListener) EnterIncDecStmt(ctx *IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *BaseGolangListener) ExitIncDecStmt(ctx *IncDecStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseGolangListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseGolangListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *BaseGolangListener) EnterAssign_op(ctx *Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *BaseGolangListener) ExitAssign_op(ctx *Assign_opContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *BaseGolangListener) EnterShortVarDecl(ctx *ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *BaseGolangListener) ExitShortVarDecl(ctx *ShortVarDeclContext) {}

// EnterEmptyStmt is called when production emptyStmt is entered.
func (s *BaseGolangListener) EnterEmptyStmt(ctx *EmptyStmtContext) {}

// ExitEmptyStmt is called when production emptyStmt is exited.
func (s *BaseGolangListener) ExitEmptyStmt(ctx *EmptyStmtContext) {}

// EnterLabeledStmt is called when production labeledStmt is entered.
func (s *BaseGolangListener) EnterLabeledStmt(ctx *LabeledStmtContext) {}

// ExitLabeledStmt is called when production labeledStmt is exited.
func (s *BaseGolangListener) ExitLabeledStmt(ctx *LabeledStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseGolangListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseGolangListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseGolangListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseGolangListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseGolangListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseGolangListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterGotoStmt is called when production gotoStmt is entered.
func (s *BaseGolangListener) EnterGotoStmt(ctx *GotoStmtContext) {}

// ExitGotoStmt is called when production gotoStmt is exited.
func (s *BaseGolangListener) ExitGotoStmt(ctx *GotoStmtContext) {}

// EnterFallthroughStmt is called when production fallthroughStmt is entered.
func (s *BaseGolangListener) EnterFallthroughStmt(ctx *FallthroughStmtContext) {}

// ExitFallthroughStmt is called when production fallthroughStmt is exited.
func (s *BaseGolangListener) ExitFallthroughStmt(ctx *FallthroughStmtContext) {}

// EnterDeferStmt is called when production deferStmt is entered.
func (s *BaseGolangListener) EnterDeferStmt(ctx *DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *BaseGolangListener) ExitDeferStmt(ctx *DeferStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseGolangListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseGolangListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *BaseGolangListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *BaseGolangListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterExprSwitchStmt is called when production exprSwitchStmt is entered.
func (s *BaseGolangListener) EnterExprSwitchStmt(ctx *ExprSwitchStmtContext) {}

// ExitExprSwitchStmt is called when production exprSwitchStmt is exited.
func (s *BaseGolangListener) ExitExprSwitchStmt(ctx *ExprSwitchStmtContext) {}

// EnterExprCaseClause is called when production exprCaseClause is entered.
func (s *BaseGolangListener) EnterExprCaseClause(ctx *ExprCaseClauseContext) {}

// ExitExprCaseClause is called when production exprCaseClause is exited.
func (s *BaseGolangListener) ExitExprCaseClause(ctx *ExprCaseClauseContext) {}

// EnterExprSwitchCase is called when production exprSwitchCase is entered.
func (s *BaseGolangListener) EnterExprSwitchCase(ctx *ExprSwitchCaseContext) {}

// ExitExprSwitchCase is called when production exprSwitchCase is exited.
func (s *BaseGolangListener) ExitExprSwitchCase(ctx *ExprSwitchCaseContext) {}

// EnterTypeSwitchStmt is called when production typeSwitchStmt is entered.
func (s *BaseGolangListener) EnterTypeSwitchStmt(ctx *TypeSwitchStmtContext) {}

// ExitTypeSwitchStmt is called when production typeSwitchStmt is exited.
func (s *BaseGolangListener) ExitTypeSwitchStmt(ctx *TypeSwitchStmtContext) {}

// EnterTypeSwitchGuard is called when production typeSwitchGuard is entered.
func (s *BaseGolangListener) EnterTypeSwitchGuard(ctx *TypeSwitchGuardContext) {}

// ExitTypeSwitchGuard is called when production typeSwitchGuard is exited.
func (s *BaseGolangListener) ExitTypeSwitchGuard(ctx *TypeSwitchGuardContext) {}

// EnterTypeCaseClause is called when production typeCaseClause is entered.
func (s *BaseGolangListener) EnterTypeCaseClause(ctx *TypeCaseClauseContext) {}

// ExitTypeCaseClause is called when production typeCaseClause is exited.
func (s *BaseGolangListener) ExitTypeCaseClause(ctx *TypeCaseClauseContext) {}

// EnterTypeSwitchCase is called when production typeSwitchCase is entered.
func (s *BaseGolangListener) EnterTypeSwitchCase(ctx *TypeSwitchCaseContext) {}

// ExitTypeSwitchCase is called when production typeSwitchCase is exited.
func (s *BaseGolangListener) ExitTypeSwitchCase(ctx *TypeSwitchCaseContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseGolangListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseGolangListener) ExitTypeList(ctx *TypeListContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseGolangListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseGolangListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterCommClause is called when production commClause is entered.
func (s *BaseGolangListener) EnterCommClause(ctx *CommClauseContext) {}

// ExitCommClause is called when production commClause is exited.
func (s *BaseGolangListener) ExitCommClause(ctx *CommClauseContext) {}

// EnterCommCase is called when production commCase is entered.
func (s *BaseGolangListener) EnterCommCase(ctx *CommCaseContext) {}

// ExitCommCase is called when production commCase is exited.
func (s *BaseGolangListener) ExitCommCase(ctx *CommCaseContext) {}

// EnterRecvStmt is called when production recvStmt is entered.
func (s *BaseGolangListener) EnterRecvStmt(ctx *RecvStmtContext) {}

// ExitRecvStmt is called when production recvStmt is exited.
func (s *BaseGolangListener) ExitRecvStmt(ctx *RecvStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseGolangListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseGolangListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterForClause is called when production forClause is entered.
func (s *BaseGolangListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *BaseGolangListener) ExitForClause(ctx *ForClauseContext) {}

// EnterRangeClause is called when production rangeClause is entered.
func (s *BaseGolangListener) EnterRangeClause(ctx *RangeClauseContext) {}

// ExitRangeClause is called when production rangeClause is exited.
func (s *BaseGolangListener) ExitRangeClause(ctx *RangeClauseContext) {}

// EnterGoStmt is called when production goStmt is entered.
func (s *BaseGolangListener) EnterGoStmt(ctx *GoStmtContext) {}

// ExitGoStmt is called when production goStmt is exited.
func (s *BaseGolangListener) ExitGoStmt(ctx *GoStmtContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseGolangListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseGolangListener) ExitType_(ctx *Type_Context) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseGolangListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseGolangListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *BaseGolangListener) EnterTypeLit(ctx *TypeLitContext) {}

// ExitTypeLit is called when production typeLit is exited.
func (s *BaseGolangListener) ExitTypeLit(ctx *TypeLitContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseGolangListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseGolangListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterArrayLength is called when production arrayLength is entered.
func (s *BaseGolangListener) EnterArrayLength(ctx *ArrayLengthContext) {}

// ExitArrayLength is called when production arrayLength is exited.
func (s *BaseGolangListener) ExitArrayLength(ctx *ArrayLengthContext) {}

// EnterElementType is called when production elementType is entered.
func (s *BaseGolangListener) EnterElementType(ctx *ElementTypeContext) {}

// ExitElementType is called when production elementType is exited.
func (s *BaseGolangListener) ExitElementType(ctx *ElementTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BaseGolangListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BaseGolangListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BaseGolangListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BaseGolangListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterSliceType is called when production sliceType is entered.
func (s *BaseGolangListener) EnterSliceType(ctx *SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *BaseGolangListener) ExitSliceType(ctx *SliceTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseGolangListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseGolangListener) ExitMapType(ctx *MapTypeContext) {}

// EnterChannelType is called when production channelType is entered.
func (s *BaseGolangListener) EnterChannelType(ctx *ChannelTypeContext) {}

// ExitChannelType is called when production channelType is exited.
func (s *BaseGolangListener) ExitChannelType(ctx *ChannelTypeContext) {}

// EnterMethodSpec is called when production methodSpec is entered.
func (s *BaseGolangListener) EnterMethodSpec(ctx *MethodSpecContext) {}

// ExitMethodSpec is called when production methodSpec is exited.
func (s *BaseGolangListener) ExitMethodSpec(ctx *MethodSpecContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseGolangListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseGolangListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterSignature is called when production signature is entered.
func (s *BaseGolangListener) EnterSignature(ctx *SignatureContext) {}

// ExitSignature is called when production signature is exited.
func (s *BaseGolangListener) ExitSignature(ctx *SignatureContext) {}

// EnterResult is called when production result is entered.
func (s *BaseGolangListener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *BaseGolangListener) ExitResult(ctx *ResultContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseGolangListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseGolangListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseGolangListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseGolangListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *BaseGolangListener) EnterParameterDecl(ctx *ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *BaseGolangListener) ExitParameterDecl(ctx *ParameterDeclContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseGolangListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseGolangListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseGolangListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseGolangListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BaseGolangListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BaseGolangListener) ExitBasicLit(ctx *BasicLitContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BaseGolangListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BaseGolangListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *BaseGolangListener) EnterQualifiedIdent(ctx *QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *BaseGolangListener) ExitQualifiedIdent(ctx *QualifiedIdentContext) {}

// EnterCompositeLit is called when production compositeLit is entered.
func (s *BaseGolangListener) EnterCompositeLit(ctx *CompositeLitContext) {}

// ExitCompositeLit is called when production compositeLit is exited.
func (s *BaseGolangListener) ExitCompositeLit(ctx *CompositeLitContext) {}

// EnterLiteralType is called when production literalType is entered.
func (s *BaseGolangListener) EnterLiteralType(ctx *LiteralTypeContext) {}

// ExitLiteralType is called when production literalType is exited.
func (s *BaseGolangListener) ExitLiteralType(ctx *LiteralTypeContext) {}

// EnterLiteralValue is called when production literalValue is entered.
func (s *BaseGolangListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production literalValue is exited.
func (s *BaseGolangListener) ExitLiteralValue(ctx *LiteralValueContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseGolangListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseGolangListener) ExitElementList(ctx *ElementListContext) {}

// EnterKeyedElement is called when production keyedElement is entered.
func (s *BaseGolangListener) EnterKeyedElement(ctx *KeyedElementContext) {}

// ExitKeyedElement is called when production keyedElement is exited.
func (s *BaseGolangListener) ExitKeyedElement(ctx *KeyedElementContext) {}

// EnterKey is called when production key is entered.
func (s *BaseGolangListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseGolangListener) ExitKey(ctx *KeyContext) {}

// EnterElement is called when production element is entered.
func (s *BaseGolangListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseGolangListener) ExitElement(ctx *ElementContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseGolangListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseGolangListener) ExitStructType(ctx *StructTypeContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BaseGolangListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BaseGolangListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterAnonymousField is called when production anonymousField is entered.
func (s *BaseGolangListener) EnterAnonymousField(ctx *AnonymousFieldContext) {}

// ExitAnonymousField is called when production anonymousField is exited.
func (s *BaseGolangListener) ExitAnonymousField(ctx *AnonymousFieldContext) {}

// EnterFunctionLit is called when production functionLit is entered.
func (s *BaseGolangListener) EnterFunctionLit(ctx *FunctionLitContext) {}

// ExitFunctionLit is called when production functionLit is exited.
func (s *BaseGolangListener) ExitFunctionLit(ctx *FunctionLitContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseGolangListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseGolangListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseGolangListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseGolangListener) ExitSelector(ctx *SelectorContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseGolangListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseGolangListener) ExitIndex(ctx *IndexContext) {}

// EnterSlice is called when production slice is entered.
func (s *BaseGolangListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production slice is exited.
func (s *BaseGolangListener) ExitSlice(ctx *SliceContext) {}

// EnterTypeAssertion is called when production typeAssertion is entered.
func (s *BaseGolangListener) EnterTypeAssertion(ctx *TypeAssertionContext) {}

// ExitTypeAssertion is called when production typeAssertion is exited.
func (s *BaseGolangListener) ExitTypeAssertion(ctx *TypeAssertionContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGolangListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGolangListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterMethodExpr is called when production methodExpr is entered.
func (s *BaseGolangListener) EnterMethodExpr(ctx *MethodExprContext) {}

// ExitMethodExpr is called when production methodExpr is exited.
func (s *BaseGolangListener) ExitMethodExpr(ctx *MethodExprContext) {}

// EnterReceiverType is called when production receiverType is entered.
func (s *BaseGolangListener) EnterReceiverType(ctx *ReceiverTypeContext) {}

// ExitReceiverType is called when production receiverType is exited.
func (s *BaseGolangListener) ExitReceiverType(ctx *ReceiverTypeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGolangListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGolangListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBinary_op is called when production binary_op is entered.
func (s *BaseGolangListener) EnterBinary_op(ctx *Binary_opContext) {}

// ExitBinary_op is called when production binary_op is exited.
func (s *BaseGolangListener) ExitBinary_op(ctx *Binary_opContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseGolangListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseGolangListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterUnary_op is called when production unary_op is entered.
func (s *BaseGolangListener) EnterUnary_op(ctx *Unary_opContext) {}

// ExitUnary_op is called when production unary_op is exited.
func (s *BaseGolangListener) ExitUnary_op(ctx *Unary_opContext) {}

// EnterConversion is called when production conversion is entered.
func (s *BaseGolangListener) EnterConversion(ctx *ConversionContext) {}

// ExitConversion is called when production conversion is exited.
func (s *BaseGolangListener) ExitConversion(ctx *ConversionContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseGolangListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseGolangListener) ExitEos(ctx *EosContext) {}
