// Generated from ./parser/Og.g4 by ANTLR 4.7.1.

package parser // Og

//import "github.com/antlr/antlr4/runtime/Go/antlr"

//type OgVisitor struct {
//    *antlr.BaseParseTreeVisitor
//}

//func (v *OgVisitor) Init() interface{} {
//    return nil
//}

//func (v *OgVisitor) VisitNext(node antlr.Tree, resultSoFar interface{}) bool {
//    return true
//}

//func (v *OgVisitor) Aggregate(resultSoFar, childResult interface{}) interface{} {
//    return childResult
//}

//func (v *OgVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
//    return nil
//}

//func (v *OgVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
//    return nil
//}

//func (v *OgVisitor) VisitSourceFile(ctx *parser.SourceFileContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitPackageClause(ctx *parser.PackageClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitImportDecl(ctx *parser.ImportDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitImportSpec(ctx *parser.ImportSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitImportPath(ctx *parser.ImportPathContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTopLevelDecl(ctx *parser.TopLevelDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitDeclaration(ctx *parser.DeclarationContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitConstDecl(ctx *parser.ConstDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitConstSpec(ctx *parser.ConstSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitExpressionList(ctx *parser.ExpressionListContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeDecl(ctx *parser.TypeDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitFunction(ctx *parser.FunctionContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitMethodDecl(ctx *parser.MethodDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitReceiver(ctx *parser.ReceiverContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitVarDecl(ctx *parser.VarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitVarSpec(ctx *parser.VarSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitBlock(ctx *parser.BlockContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitStatementList(ctx *parser.StatementListContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitStatementNoBlock(ctx *parser.StatementNoBlockContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitStatement(ctx *parser.StatementContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitSendStmt(ctx *parser.SendStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitIncDecStmt(ctx *parser.IncDecStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitAssignment(ctx *parser.AssignmentContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitAssign_op(ctx *parser.Assign_opContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitShortVarDecl(ctx *parser.ShortVarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitEmptyStmt(ctx *parser.EmptyStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitLabeledStmt(ctx *parser.LabeledStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitGotoStmt(ctx *parser.GotoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitFallthroughStmt(ctx *parser.FallthroughStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitDeferStmt(ctx *parser.DeferStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitIfStmt(ctx *parser.IfStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitExprSwitchStmt(ctx *parser.ExprSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitExprCaseClause(ctx *parser.ExprCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitExprSwitchCase(ctx *parser.ExprSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeSwitchStmt(ctx *parser.TypeSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeSwitchGuard(ctx *parser.TypeSwitchGuardContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeCaseClause(ctx *parser.TypeCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeSwitchCase(ctx *parser.TypeSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeList(ctx *parser.TypeListContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitSelectStmt(ctx *parser.SelectStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitCommClause(ctx *parser.CommClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitCommCase(ctx *parser.CommCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitRecvStmt(ctx *parser.RecvStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitForStmt(ctx *parser.ForStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitForClause(ctx *parser.ForClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitRangeClause(ctx *parser.RangeClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitGoStmt(ctx *parser.GoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitType_(ctx *parser.Type_Context, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeName(ctx *parser.TypeNameContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeLit(ctx *parser.TypeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitArrayType(ctx *parser.ArrayTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitArrayLength(ctx *parser.ArrayLengthContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitElementType(ctx *parser.ElementTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitPointerType(ctx *parser.PointerTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitInterfaceType(ctx *parser.InterfaceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitSliceType(ctx *parser.SliceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitMapType(ctx *parser.MapTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitChannelType(ctx *parser.ChannelTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitMethodSpec(ctx *parser.MethodSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitFunctionType(ctx *parser.FunctionTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitSignature(ctx *parser.SignatureContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitResult(ctx *parser.ResultContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitParameters(ctx *parser.ParametersContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitParameterList(ctx *parser.ParameterListContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitParameterDecl(ctx *parser.ParameterDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitOperand(ctx *parser.OperandContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitLiteral(ctx *parser.LiteralContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitBasicLit(ctx *parser.BasicLitContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitOperandName(ctx *parser.OperandNameContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitCompositeLit(ctx *parser.CompositeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitLiteralType(ctx *parser.LiteralTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitLiteralValue(ctx *parser.LiteralValueContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitElementList(ctx *parser.ElementListContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitKeyedElement(ctx *parser.KeyedElementContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitKey(ctx *parser.KeyContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitElement(ctx *parser.ElementContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitStructType(ctx *parser.StructTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitFieldDecl(ctx *parser.FieldDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitAnonymousField(ctx *parser.AnonymousFieldContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitFunctionLit(ctx *parser.FunctionLitContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitSelector(ctx *parser.SelectorContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitIndex(ctx *parser.IndexContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitSlice(ctx *parser.SliceContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitTypeAssertion(ctx *parser.TypeAssertionContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitArguments(ctx *parser.ArgumentsContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitMethodExpr(ctx *parser.MethodExprContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitReceiverType(ctx *parser.ReceiverTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitExpression(ctx *parser.ExpressionContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitBinary_op(ctx *parser.Binary_opContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitUnary_op(ctx *parser.Unary_opContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitConversion(ctx *parser.ConversionContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}

//func (v *OgVisitor) VisitEos(ctx *parser.EosContext, delegate antlr.ParseTreeVisitor) interface{} {
//  // before children
//  r := v.VisitChildren(ctx, delegate)
//  // afer children
//  return r
//}
