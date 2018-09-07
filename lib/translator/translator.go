package translator

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/champii/og/lib/ast"
	"github.com/champii/og/parser"
	"strings"
)

type OgVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (this OgVisitor) Aggregate(resultSoFar interface{}, childResult interface{}) interface{} {
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
	node := &SourceFile{
		Node:    NewNode(ctx),
		Package: this.VisitPackageClause(ctx.PackageClause().(*parser.PackageClauseContext), delegate).(*Package),
	}
	if ctx.ImportDecl(0) != nil {
		node.Import = this.VisitImportDecl(ctx.ImportDecl(0).(*parser.ImportDeclContext), delegate).(*Import)
	}
	res := []*TopLevel{}
	bodies := ctx.AllTopLevelDecl()
	for _, spec := range bodies {
		res = append(res, this.VisitTopLevelDecl(spec.(*parser.TopLevelDeclContext), delegate).(*TopLevel))
	}
	node.TopLevels = res
	return node
}
func (this *OgVisitor) VisitPackageClause(ctx *parser.PackageClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &Package{
		Node: NewNode(ctx),
		Name: ctx.IDENTIFIER().GetText(),
	}
}
func (this *OgVisitor) VisitImportDecl(ctx *parser.ImportDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &Import{
		Node:  NewNode(ctx),
		Items: this.VisitImportBody(ctx.ImportBody().(*parser.ImportBodyContext), delegate).([]*ImportSpec),
	}
}
func (this *OgVisitor) VisitImportBody(ctx *parser.ImportBodyContext, delegate antlr.ParseTreeVisitor) interface{} {
	res := []*ImportSpec{}
	bodies := ctx.AllImportSpec()
	for _, spec := range bodies {
		res = append(res, this.VisitImportSpec(spec.(*parser.ImportSpecContext), delegate).(*ImportSpec))
	}
	return res
}
func (this *OgVisitor) VisitImportSpec(ctx *parser.ImportSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ImportSpec{Node: NewNode(ctx)}
	if ctx.ImportPath() != nil {
		node.Path = this.VisitImportPath(ctx.ImportPath().(*parser.ImportPathContext), delegate).(string)
	}
	if ctx.IDENTIFIER() != nil {
		node.Alias = ctx.IDENTIFIER().GetText()
	} else if strings.Contains(ctx.GetText(), ":") {
		node.Alias = "."
	}
	return node
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
	node := &TopLevel{Node: NewNode(ctx)}
	if ctx.Declaration() != nil {
		node.Declaration = this.VisitDeclaration(ctx.Declaration().(*parser.DeclarationContext), delegate).(*Declaration)
	}
	if ctx.FunctionDecl() != nil {
		node.FunctionDecl = this.VisitFunctionDecl(ctx.FunctionDecl().(*parser.FunctionDeclContext), delegate).(*FunctionDecl)
	}
	if ctx.MethodDecl() != nil {
		node.MethodDecl = this.VisitMethodDecl(ctx.MethodDecl().(*parser.MethodDeclContext), delegate).(*MethodDecl)
	}
	return node
}
func (this *OgVisitor) VisitDeclaration(ctx *parser.DeclarationContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Declaration{Node: NewNode(ctx)}
	if ctx.ConstDecl() != nil {
		node.ConstDecl = this.VisitConstDecl(ctx.ConstDecl().(*parser.ConstDeclContext), delegate).(*ConstDecl)
	}
	if ctx.TypeDecl() != nil {
		node.TypeDecl = this.VisitTypeDecl(ctx.TypeDecl().(*parser.TypeDeclContext), delegate).(*TypeDecl)
	}
	if ctx.VarDecl() != nil {
		node.VarDecl = this.VisitVarDecl(ctx.VarDecl().(*parser.VarDeclContext), delegate).(*VarDecl)
	}
	return node
}
func (this *OgVisitor) VisitConstDecl(ctx *parser.ConstDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ConstDecl{Node: NewNode(ctx)}
	res := []*ConstSpec{}
	bodies := ctx.AllConstSpec()
	for _, spec := range bodies {
		res = append(res, this.VisitConstSpec(spec.(*parser.ConstSpecContext), delegate).(*ConstSpec))
	}
	node.ConstSpecs = res
	return node
}
func (this *OgVisitor) VisitConstSpec(ctx *parser.ConstSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ConstSpec{Node: NewNode(ctx)}
	if ctx.IdentifierList() != nil {
		node.IdentifierList = this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).([]string)
	}
	if ctx.Type_() != nil {
		node.Type = this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type)
	}
	if ctx.ExpressionList() != nil {
		node.ExpressionList = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(*ExpressionList)
	}
	return node
}
func (this *OgVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext, delegate antlr.ParseTreeVisitor) interface{} {
	return strings.Split(ctx.GetText(), ",")
}
func (this *OgVisitor) VisitExpressionList(ctx *parser.ExpressionListContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ExpressionList{Node: NewNode(ctx)}
	res := []*Expression{}
	bodies := ctx.AllExpression()
	for _, spec := range bodies {
		res = append(res, this.VisitExpression(spec.(*parser.ExpressionContext), delegate).(*Expression))
	}
	node.Expressions = res
	return node
}
func (this *OgVisitor) VisitTypeDecl(ctx *parser.TypeDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &TypeDecl{Node: NewNode(ctx)}
	if len(ctx.AllTypeSpec()) > 0 {
		res := []*TypeSpec{}
		bodies := ctx.AllTypeSpec()
		for _, spec := range bodies {
			res = append(res, this.VisitTypeSpec(spec.(*parser.TypeSpecContext), delegate).(*TypeSpec))
		}
		node.TypeSpecs = res
	}
	if ctx.StructType() != nil {
		node.StructType = this.VisitStructType(ctx.StructType().(*parser.StructTypeContext), delegate).(*StructType)
	}
	if ctx.InterfaceType() != nil {
		node.InterfaceType = this.VisitInterfaceType(ctx.InterfaceType().(*parser.InterfaceTypeContext), delegate).(*InterfaceType)
	}
	return node
}
func (this *OgVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &TypeSpec{
		Node: NewNode(ctx),
		Name: ctx.IDENTIFIER().GetText(),
		Type: this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type),
	}
}
func (this *OgVisitor) VisitFunctionDecl(ctx *parser.FunctionDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &FunctionDecl{
		Node: NewNode(ctx),
		Name: ctx.IDENTIFIER().GetText(),
	}
	if ctx.Function() != nil {
		node.Function = this.VisitFunction(ctx.Function().(*parser.FunctionContext), delegate).(*Function)
	}
	if ctx.Signature() != nil {
		node.Signature = this.VisitSignature(ctx.Signature().(*parser.SignatureContext), delegate).(*Signature)
	}
	return node
}
func (this *OgVisitor) VisitFunction(ctx *parser.FunctionContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Function{
		Node:      NewNode(ctx),
		Signature: this.VisitSignature(ctx.Signature().(*parser.SignatureContext), delegate).(*Signature),
	}
	if ctx.Block() != nil {
		node.Block = this.VisitBlock(ctx.Block().(*parser.BlockContext), delegate).(*Block)
	}
	if ctx.StatementNoBlock() != nil {
		node.Block = this.VisitStatementNoBlock(ctx.StatementNoBlock().(*parser.StatementNoBlockContext), delegate).(*Block)
	}
	return node
}
func (this *OgVisitor) VisitMethodDecl(ctx *parser.MethodDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &MethodDecl{
		Node:     NewNode(ctx),
		Receiver: this.VisitReceiver(ctx.Receiver().(*parser.ReceiverContext), delegate).(*Receiver),
	}
	if ctx.Function() != nil {
		node.Function = this.VisitFunction(ctx.Function().(*parser.FunctionContext), delegate).(*Function)
	}
	if ctx.Signature() != nil {
		node.Signature = this.VisitSignature(ctx.Signature().(*parser.SignatureContext), delegate).(*Signature)
	}
	return node
}
func (this *OgVisitor) VisitReceiver(ctx *parser.ReceiverContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &Receiver{
		Node:              NewNode(ctx),
		Package:           ctx.IDENTIFIER(0).GetText(),
		IsPointerReceiver: strings.Contains(ctx.GetText(), "*"),
		Method:            ctx.IDENTIFIER(1).GetText(),
	}
}
func (this *OgVisitor) VisitVarDecl(ctx *parser.VarDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &VarDecl{Node: NewNode(ctx)}
	res := []*VarSpec{}
	bodies := ctx.AllVarSpec()
	for _, spec := range bodies {
		res = append(res, this.VisitVarSpec(spec.(*parser.VarSpecContext), delegate).(*VarSpec))
	}
	node.VarSpecs = res
	return node
}
func (this *OgVisitor) VisitVarSpec(ctx *parser.VarSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &VarSpec{
		Node:           NewNode(ctx),
		IdentifierList: this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).([]string),
	}
	if ctx.Type_() != nil {
		node.Type = this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type)
	}
	if ctx.ExpressionList() != nil {
		node.ExpressionList = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(*ExpressionList)
	}
	return node
}
func (this *OgVisitor) VisitBlock(ctx *parser.BlockContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &Block{
		Node:       NewNode(ctx),
		Statements: this.VisitStatementList(ctx.StatementList().(*parser.StatementListContext), delegate).([]*Statement),
	}
}
func (this *OgVisitor) VisitStatementList(ctx *parser.StatementListContext, delegate antlr.ParseTreeVisitor) interface{} {
	res := []*Statement{}
	bodies := ctx.AllStatement()
	for _, spec := range bodies {
		res = append(res, this.VisitStatement(spec.(*parser.StatementContext), delegate).(*Statement))
	}
	return res
}
func (this *OgVisitor) VisitStatement(ctx *parser.StatementContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Statement{Node: NewNode(ctx)}
	if ctx.SimpleStmt() != nil {
		node.SimpleStmt = this.VisitSimpleStmt(ctx.SimpleStmt().(*parser.SimpleStmtContext), delegate).(*SimpleStmt)
	}
	if ctx.LabeledStmt() != nil {
		node.LabeledStmt = this.VisitLabeledStmt(ctx.LabeledStmt().(*parser.LabeledStmtContext), delegate).(*LabeledStmt)
	}
	if ctx.GoStmt() != nil {
		node.GoStmt = this.VisitGoStmt(ctx.GoStmt().(*parser.GoStmtContext), delegate).(*GoStmt)
	}
	if ctx.ReturnStmt() != nil {
		node.ReturnStmt = this.VisitReturnStmt(ctx.ReturnStmt().(*parser.ReturnStmtContext), delegate).(*ReturnStmt)
	}
	if ctx.BreakStmt() != nil {
		node.BreakStmt = this.VisitBreakStmt(ctx.BreakStmt().(*parser.BreakStmtContext), delegate).(*BreakStmt)
	}
	if ctx.ContinueStmt() != nil {
		node.ContinueStmt = this.VisitContinueStmt(ctx.ContinueStmt().(*parser.ContinueStmtContext), delegate).(*ContinueStmt)
	}
	if ctx.GotoStmt() != nil {
		node.GotoStmt = this.VisitGotoStmt(ctx.GotoStmt().(*parser.GotoStmtContext), delegate).(*GotoStmt)
	}
	if ctx.FallthroughStmt() != nil {
		node.FallthroughStmt = this.VisitFallthroughStmt(ctx.FallthroughStmt().(*parser.FallthroughStmtContext), delegate).(*FallthroughStmt)
	}
	if ctx.IfStmt() != nil {
		node.IfStmt = this.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext), delegate).(*IfStmt)
	}
	if ctx.SwitchStmt() != nil {
		node.SwitchStmt = this.VisitSwitchStmt(ctx.SwitchStmt().(*parser.SwitchStmtContext), delegate).(*SwitchStmt)
	}
	if ctx.SelectStmt() != nil {
		node.SelectStmt = this.VisitSelectStmt(ctx.SelectStmt().(*parser.SelectStmtContext), delegate).(*SelectStmt)
	}
	if ctx.ForStmt() != nil {
		node.ForStmt = this.VisitForStmt(ctx.ForStmt().(*parser.ForStmtContext), delegate).(*ForStmt)
	}
	if ctx.Block() != nil {
		node.Block = this.VisitBlock(ctx.Block().(*parser.BlockContext), delegate).(*Block)
	}
	if ctx.DeferStmt() != nil {
		node.DeferStmt = this.VisitDeferStmt(ctx.DeferStmt().(*parser.DeferStmtContext), delegate).(*DeferStmt)
	}
	return node
}
func (this *OgVisitor) VisitStatementNoBlock(ctx *parser.StatementNoBlockContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Block{Node: NewNode(ctx)}
	res := &Statement{}
	if ctx.SimpleStmt() != nil {
		res.SimpleStmt = this.VisitSimpleStmt(ctx.SimpleStmt().(*parser.SimpleStmtContext), delegate).(*SimpleStmt)
	}
	if ctx.LabeledStmt() != nil {
		res.LabeledStmt = this.VisitLabeledStmt(ctx.LabeledStmt().(*parser.LabeledStmtContext), delegate).(*LabeledStmt)
	}
	if ctx.GoStmt() != nil {
		res.GoStmt = this.VisitGoStmt(ctx.GoStmt().(*parser.GoStmtContext), delegate).(*GoStmt)
	}
	if ctx.ReturnStmt() != nil {
		res.ReturnStmt = this.VisitReturnStmt(ctx.ReturnStmt().(*parser.ReturnStmtContext), delegate).(*ReturnStmt)
	}
	if ctx.BreakStmt() != nil {
		res.BreakStmt = this.VisitBreakStmt(ctx.BreakStmt().(*parser.BreakStmtContext), delegate).(*BreakStmt)
	}
	if ctx.ContinueStmt() != nil {
		res.ContinueStmt = this.VisitContinueStmt(ctx.ContinueStmt().(*parser.ContinueStmtContext), delegate).(*ContinueStmt)
	}
	if ctx.GotoStmt() != nil {
		res.GotoStmt = this.VisitGotoStmt(ctx.GotoStmt().(*parser.GotoStmtContext), delegate).(*GotoStmt)
	}
	if ctx.FallthroughStmt() != nil {
		res.FallthroughStmt = this.VisitFallthroughStmt(ctx.FallthroughStmt().(*parser.FallthroughStmtContext), delegate).(*FallthroughStmt)
	}
	if ctx.IfStmt() != nil {
		res.IfStmt = this.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext), delegate).(*IfStmt)
	}
	if ctx.SwitchStmt() != nil {
		res.SwitchStmt = this.VisitSwitchStmt(ctx.SwitchStmt().(*parser.SwitchStmtContext), delegate).(*SwitchStmt)
	}
	if ctx.SelectStmt() != nil {
		res.SelectStmt = this.VisitSelectStmt(ctx.SelectStmt().(*parser.SelectStmtContext), delegate).(*SelectStmt)
	}
	if ctx.ForStmt() != nil {
		res.ForStmt = this.VisitForStmt(ctx.ForStmt().(*parser.ForStmtContext), delegate).(*ForStmt)
	}
	if ctx.DeferStmt() != nil {
		res.DeferStmt = this.VisitDeferStmt(ctx.DeferStmt().(*parser.DeferStmtContext), delegate).(*DeferStmt)
	}
	node.Statements = []*Statement{res}
	return node
}
func (this *OgVisitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &SimpleStmt{Node: NewNode(ctx)}
	if ctx.SendStmt() != nil {
		node.SendStmt = this.VisitSendStmt(ctx.SendStmt().(*parser.SendStmtContext), delegate).(*SendStmt)
	}
	if ctx.ExpressionStmt() != nil {
		node.Expression = this.VisitExpressionStmt(ctx.ExpressionStmt().(*parser.ExpressionStmtContext), delegate).(*Expression)
	}
	if ctx.IncDecStmt() != nil {
		node.IncDecStmt = this.VisitIncDecStmt(ctx.IncDecStmt().(*parser.IncDecStmtContext), delegate).(*IncDecStmt)
	}
	if ctx.ShortVarDecl() != nil {
		node.ShortVarDecl = this.VisitShortVarDecl(ctx.ShortVarDecl().(*parser.ShortVarDeclContext), delegate).(*ShortVarDecl)
	}
	if ctx.Assignment() != nil {
		node.Assignment = this.VisitAssignment(ctx.Assignment().(*parser.AssignmentContext), delegate).(*Assignment)
	}
	if ctx.EmptyStmt() != nil {
		node.EmptyStmt = true
	}
	return node
}
func (this *OgVisitor) VisitExpressionStmt(ctx *parser.ExpressionStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
}
func (this *OgVisitor) VisitSendStmt(ctx *parser.SendStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &SendStmt{
		Node:  NewNode(ctx),
		Left:  this.VisitExpression(ctx.Expression(0).(*parser.ExpressionContext), delegate).(*Expression),
		Right: this.VisitExpression(ctx.Expression(1).(*parser.ExpressionContext), delegate).(*Expression),
	}
}
func (this *OgVisitor) VisitIncDecStmt(ctx *parser.IncDecStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &IncDecStmt{
		Node:       NewNode(ctx),
		Expression: this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression),
		IsInc:      strings.Contains(ctx.GetText(), "++"),
	}
}
func (this *OgVisitor) VisitAssignment(ctx *parser.AssignmentContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &Assignment{
		Node:  NewNode(ctx),
		Left:  this.VisitExpressionList(ctx.ExpressionList(0).(*parser.ExpressionListContext), delegate).(*ExpressionList),
		Op:    this.VisitAssign_op(ctx.Assign_op().(*parser.Assign_opContext), delegate).(string),
		Right: this.VisitExpressionList(ctx.ExpressionList(1).(*parser.ExpressionListContext), delegate).(*ExpressionList),
	}
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
	node := &ShortVarDecl{
		Node:        NewNode(ctx),
		Identifiers: this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).([]string),
	}
	if ctx.ExpressionList() != nil {
		node.Expressions = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(*ExpressionList)
	} else if ctx.StatementNoBlock() != nil {
		node.Statement = this.VisitStatementNoBlock(ctx.StatementNoBlock().(*parser.StatementNoBlockContext), delegate).(*Block).Statements[0]
	}
	return node
}
func (this *OgVisitor) VisitEmptyStmt(ctx *parser.EmptyStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "\n"
}
func (this *OgVisitor) VisitLabeledStmt(ctx *parser.LabeledStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &LabeledStmt{
		Node:      NewNode(ctx),
		Name:      ctx.IDENTIFIER().GetText(),
		Statement: this.VisitStatement(ctx.Statement().(*parser.StatementContext), delegate).(*Statement),
	}
}
func (this *OgVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ReturnStmt{Node: NewNode(ctx)}
	if ctx.ExpressionList() != nil {
		node.Expressions = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(*ExpressionList)
	}
	return node
}
func (this *OgVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &BreakStmt{Node: NewNode(ctx)}
	if ctx.IDENTIFIER() != nil {
		node.Name = ctx.IDENTIFIER().GetText()
	}
	return node
}
func (this *OgVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ContinueStmt{Node: NewNode(ctx)}
	if ctx.IDENTIFIER() != nil {
		node.Name = ctx.IDENTIFIER().GetText()
	}
	return node
}
func (this *OgVisitor) VisitGotoStmt(ctx *parser.GotoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &GotoStmt{
		Node: NewNode(ctx),
		Name: ctx.IDENTIFIER().GetText(),
	}
}
func (this *OgVisitor) VisitFallthroughStmt(ctx *parser.FallthroughStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &FallthroughStmt{Node: NewNode(ctx)}
}
func (this *OgVisitor) VisitDeferStmt(ctx *parser.DeferStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &DeferStmt{
		Node:       NewNode(ctx),
		Expression: this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression),
	}
}
func (this *OgVisitor) VisitIfStmt(ctx *parser.IfStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &IfStmt{Node: NewNode(ctx)}
	if ctx.SimpleStmt() != nil {
		node.SimpleStmt = this.VisitSimpleStmt(ctx.SimpleStmt().(*parser.SimpleStmtContext), delegate).(*SimpleStmt)
	}
	node.Expression = this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
	if ctx.Block(0) != nil {
		node.Block = this.VisitBlock(ctx.Block(0).(*parser.BlockContext), delegate).(*Block)
	} else if ctx.StatementNoBlock(0) != nil {
		node.Block = this.VisitStatementNoBlock(ctx.StatementNoBlock(0).(*parser.StatementNoBlockContext), delegate).(*Block)
	}
	if ctx.Block(1) != nil {
		node.BlockElse = this.VisitBlock(ctx.Block(1).(*parser.BlockContext), delegate).(*Block)
	} else if ctx.StatementNoBlock(1) != nil {
		node.BlockElse = this.VisitStatementNoBlock(ctx.StatementNoBlock(1).(*parser.StatementNoBlockContext), delegate).(*Block)
	} else if ctx.IfStmt() != nil {
		node.IfStmt = this.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext), delegate).(*IfStmt)
	}
	return node
}
func (this *OgVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &SwitchStmt{Node: NewNode(ctx)}
	if ctx.ExprSwitchStmt() != nil {
		node.ExprSwitchStmt = this.VisitExprSwitchStmt(ctx.ExprSwitchStmt().(*parser.ExprSwitchStmtContext), delegate).(*ExprSwitchStmt)
	}
	if ctx.TypeSwitchStmt() != nil {
		node.TypeSwitchStmt = this.VisitTypeSwitchStmt(ctx.TypeSwitchStmt().(*parser.TypeSwitchStmtContext), delegate).(*TypeSwitchStmt)
	}
	return node
}
func (this *OgVisitor) VisitExprSwitchStmt(ctx *parser.ExprSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ExprSwitchStmt{Node: NewNode(ctx)}
	if ctx.SimpleStmt() != nil {
		node.SimpleStmt = this.VisitSimpleStmt(ctx.SimpleStmt().(*parser.SimpleStmtContext), delegate).(*SimpleStmt)
	}
	if ctx.Expression() != nil {
		node.Expression = this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
	}
	res := []*ExprCaseClause{}
	bodies := ctx.AllExprCaseClause()
	for _, spec := range bodies {
		res = append(res, this.VisitExprCaseClause(spec.(*parser.ExprCaseClauseContext), delegate).(*ExprCaseClause))
	}
	node.ExprCaseClauses = res
	return node
}
func (this *OgVisitor) VisitExprCaseClause(ctx *parser.ExprCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &ExprCaseClause{
		Node:           NewNode(ctx),
		ExprSwitchCase: this.VisitExprSwitchCase(ctx.ExprSwitchCase().(*parser.ExprSwitchCaseContext), delegate).(*ExprSwitchCase),
		Statements:     this.VisitStatementList(ctx.StatementList().(*parser.StatementListContext), delegate).([]*Statement),
	}
}
func (this *OgVisitor) VisitExprSwitchCase(ctx *parser.ExprSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ExprSwitchCase{Node: NewNode(ctx)}
	if ctx.GetText() == "_" {
		node.IsDefault = true
		return node
	}
	if ctx.ExpressionList() != nil {
		node.Expressions = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(*ExpressionList)
	}
	return node
}
func (this *OgVisitor) VisitTypeSwitchStmt(ctx *parser.TypeSwitchStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &TypeSwitchStmt{
		Node:            NewNode(ctx),
		TypeSwitchGuard: this.VisitTypeSwitchGuard(ctx.TypeSwitchGuard().(*parser.TypeSwitchGuardContext), delegate).(*TypeSwitchGuard),
	}
	if ctx.SimpleStmt() != nil {
		node.SimpleStmt = this.VisitSimpleStmt(ctx.SimpleStmt().(*parser.SimpleStmtContext), delegate).(*SimpleStmt)
	}
	res := []*TypeCaseClause{}
	bodies := ctx.AllTypeCaseClause()
	for _, spec := range bodies {
		res = append(res, this.VisitTypeCaseClause(spec.(*parser.TypeCaseClauseContext), delegate).(*TypeCaseClause))
	}
	node.TypeCaseClauses = res
	return node
}
func (this *OgVisitor) VisitTypeSwitchGuard(ctx *parser.TypeSwitchGuardContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &TypeSwitchGuard{
		Node:        NewNode(ctx),
		PrimaryExpr: this.VisitPrimaryExpr(ctx.PrimaryExpr().(*parser.PrimaryExprContext), delegate).(*PrimaryExpr),
	}
	if ctx.IDENTIFIER() != nil {
		node.Name = ctx.IDENTIFIER().GetText()
	}
	return node
}
func (this *OgVisitor) VisitTypeCaseClause(ctx *parser.TypeCaseClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &TypeCaseClause{
		Node:           NewNode(ctx),
		TypeSwitchCase: this.VisitTypeSwitchCase(ctx.TypeSwitchCase().(*parser.TypeSwitchCaseContext), delegate).(*TypeSwitchCase),
		Statements:     this.VisitStatementList(ctx.StatementList().(*parser.StatementListContext), delegate).([]*Statement),
	}
}
func (this *OgVisitor) VisitTypeSwitchCase(ctx *parser.TypeSwitchCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &TypeSwitchCase{Node: NewNode(ctx)}
	if ctx.TypeList() != nil {
		node.Types = this.VisitTypeList(ctx.TypeList().(*parser.TypeListContext), delegate).([]*Type)
	}
	return node
}
func (this *OgVisitor) VisitTypeList(ctx *parser.TypeListContext, delegate antlr.ParseTreeVisitor) interface{} {
	res := []*Type{}
	bodies := ctx.AllType_()
	for _, spec := range bodies {
		res = append(res, this.VisitType_(spec.(*parser.Type_Context), delegate).(*Type))
	}
	return res
}
func (this *OgVisitor) VisitSelectStmt(ctx *parser.SelectStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &SelectStmt{Node: NewNode(ctx)}
	res := []*CommClause{}
	bodies := ctx.AllCommClause()
	for _, spec := range bodies {
		res = append(res, this.VisitCommClause(spec.(*parser.CommClauseContext), delegate).(*CommClause))
	}
	node.CommClauses = res
	return node
}
func (this *OgVisitor) VisitCommClause(ctx *parser.CommClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &CommClause{
		Node:     NewNode(ctx),
		CommCase: this.VisitCommCase(ctx.CommCase().(*parser.CommCaseContext), delegate).(*CommCase),
	}
	if ctx.Block() != nil {
		node.Block = this.VisitBlock(ctx.Block().(*parser.BlockContext), delegate).(*Block)
	} else if ctx.StatementNoBlock() != nil {
		node.Block = this.VisitStatementNoBlock(ctx.StatementNoBlock().(*parser.StatementNoBlockContext), delegate).(*Block)
	}
	return node
}
func (this *OgVisitor) VisitCommCase(ctx *parser.CommCaseContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &CommCase{Node: NewNode(ctx)}
	if ctx.GetText() == "_" {
		node.IsDefault = true
		return node
	}
	if ctx.SendStmt() != nil {
		node.SendStmt = this.VisitSendStmt(ctx.SendStmt().(*parser.SendStmtContext), delegate).(*SendStmt)
	}
	if ctx.RecvStmt() != nil {
		node.RecvStmt = this.VisitRecvStmt(ctx.RecvStmt().(*parser.RecvStmtContext), delegate).(*RecvStmt)
	}
	return node
}
func (this *OgVisitor) VisitRecvStmt(ctx *parser.RecvStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &RecvStmt{
		Node:       NewNode(ctx),
		Expression: this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression),
	}
	if ctx.ExpressionList() != nil {
		node.Expressions = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(*ExpressionList)
	}
	if ctx.IdentifierList() != nil {
		node.Identifiers = this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).([]string)
	}
	return node
}
func (this *OgVisitor) VisitForStmt(ctx *parser.ForStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ForStmt{
		Node:  NewNode(ctx),
		Block: this.VisitBlock(ctx.Block().(*parser.BlockContext), delegate).(*Block),
	}
	if ctx.ForClause() != nil {
		node.ForClause = this.VisitForClause(ctx.ForClause().(*parser.ForClauseContext), delegate).(*ForClause)
	}
	if ctx.RangeClause() != nil {
		node.RangeClause = this.VisitRangeClause(ctx.RangeClause().(*parser.RangeClauseContext), delegate).(*RangeClause)
	}
	if ctx.Expression() != nil {
		node.Expression = this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
	}
	return node
}
func (this *OgVisitor) VisitForClause(ctx *parser.ForClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ForClause{Node: NewNode(ctx)}
	if ctx.SimpleStmt(0) != nil {
		node.LeftSimpleStmt = this.VisitSimpleStmt(ctx.SimpleStmt(0).(*parser.SimpleStmtContext), delegate).(*SimpleStmt)
	}
	if ctx.Expression() != nil {
		node.Expression = this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
	}
	if ctx.SimpleStmt(1) != nil {
		node.RightSimpleStmt = this.VisitSimpleStmt(ctx.SimpleStmt(1).(*parser.SimpleStmtContext), delegate).(*SimpleStmt)
	}
	return node
}
func (this *OgVisitor) VisitRangeClause(ctx *parser.RangeClauseContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &RangeClause{
		Node:       NewNode(ctx),
		Expression: this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression),
	}
	if ctx.ExpressionList() != nil {
		node.Expressions = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(*ExpressionList)
	}
	if ctx.IdentifierList() != nil {
		node.Identifiers = this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).([]string)
	}
	return node
}
func (this *OgVisitor) VisitGoStmt(ctx *parser.GoStmtContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &GoStmt{Node: NewNode(ctx)}
	if ctx.Expression() != nil {
		node.Expression = this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
	}
	if ctx.Function() != nil {
		node.Function = this.VisitFunction(ctx.Function().(*parser.FunctionContext), delegate).(*Function)
	}
	return node
}
func (this *OgVisitor) VisitType_(ctx *parser.Type_Context, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Type{Node: NewNode(ctx)}
	if ctx.TypeName() != nil {
		node.TypeName = this.VisitTypeName(ctx.TypeName().(*parser.TypeNameContext), delegate).(string)
	}
	if ctx.TypeLit() != nil {
		node.TypeLit = this.VisitTypeLit(ctx.TypeLit().(*parser.TypeLitContext), delegate).(*TypeLit)
	}
	if ctx.Type_() != nil {
		node.Type = this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type)
	}
	return node
}
func (this *OgVisitor) VisitTypeName(ctx *parser.TypeNameContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitTypeLit(ctx *parser.TypeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &TypeLit{Node: NewNode(ctx)}
	if ctx.ArrayType() != nil {
		node.ArrayType = this.VisitArrayType(ctx.ArrayType().(*parser.ArrayTypeContext), delegate).(*ArrayType)
	}
	if ctx.StructType() != nil {
		node.StructType = this.VisitStructType(ctx.StructType().(*parser.StructTypeContext), delegate).(*StructType)
	}
	if ctx.PointerType() != nil {
		node.PointerType = this.VisitPointerType(ctx.PointerType().(*parser.PointerTypeContext), delegate).(*PointerType)
	}
	if ctx.FunctionType() != nil {
		node.FunctionType = this.VisitFunctionType(ctx.FunctionType().(*parser.FunctionTypeContext), delegate).(*FunctionType)
	}
	if ctx.InterfaceType() != nil {
		node.InterfaceType = this.VisitInterfaceType(ctx.InterfaceType().(*parser.InterfaceTypeContext), delegate).(*InterfaceType)
	}
	if ctx.SliceType() != nil {
		node.SliceType = this.VisitSliceType(ctx.SliceType().(*parser.SliceTypeContext), delegate).(*SliceType)
	}
	if ctx.MapType() != nil {
		node.MapType = this.VisitMapType(ctx.MapType().(*parser.MapTypeContext), delegate).(*MapType)
	}
	if ctx.ChannelType() != nil {
		node.ChannelType = this.VisitChannelType(ctx.ChannelType().(*parser.ChannelTypeContext), delegate).(*ChannelType)
	}
	return node
}
func (this *OgVisitor) VisitArrayType(ctx *parser.ArrayTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &ArrayType{
		Node:        NewNode(ctx),
		Length:      this.VisitArrayLength(ctx.ArrayLength().(*parser.ArrayLengthContext), delegate).(*Expression),
		ElementType: this.VisitElementType(ctx.ElementType().(*parser.ElementTypeContext), delegate).(*Type),
	}
}
func (this *OgVisitor) VisitArrayLength(ctx *parser.ArrayLengthContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
}
func (this *OgVisitor) VisitElementType(ctx *parser.ElementTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type)
}
func (this *OgVisitor) VisitPointerType(ctx *parser.PointerTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &PointerType{
		Node: NewNode(ctx),
		Type: this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type),
	}
}
func (this *OgVisitor) VisitInterfaceType(ctx *parser.InterfaceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &InterfaceType{Node: NewNode(ctx)}
	if ctx.IDENTIFIER() != nil {
		node.Name = ctx.IDENTIFIER().GetText()
	}
	res := []*MethodSpec{}
	bodies := ctx.AllMethodSpec()
	for _, spec := range bodies {
		res = append(res, this.VisitMethodSpec(spec.(*parser.MethodSpecContext), delegate).(*MethodSpec))
	}
	node.MethodSpecs = res
	return node
}
func (this *OgVisitor) VisitSliceType(ctx *parser.SliceTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &SliceType{
		Node: NewNode(ctx),
		Type: this.VisitElementType(ctx.ElementType().(*parser.ElementTypeContext), delegate).(*Type),
	}
}
func (this *OgVisitor) VisitMapType(ctx *parser.MapTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &MapType{
		Node:      NewNode(ctx),
		InnerType: this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type),
		OuterType: this.VisitElementType(ctx.ElementType().(*parser.ElementTypeContext), delegate).(*Type),
	}
}
func (this *OgVisitor) VisitChannelType(ctx *parser.ChannelTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &ChannelType{
		Node:        NewNode(ctx),
		ChannelDecl: ctx.ChannelDecl().GetText(),
		Type:        this.VisitElementType(ctx.ElementType().(*parser.ElementTypeContext), delegate).(*Type),
	}
}
func (this *OgVisitor) VisitMethodSpec(ctx *parser.MethodSpecContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &MethodSpec{Node: NewNode(ctx)}
	if ctx.IDENTIFIER() != nil {
		node.Name = ctx.IDENTIFIER().GetText()
	}
	if ctx.Parameters() != nil {
		node.Parameters = this.VisitParameters(ctx.Parameters().(*parser.ParametersContext), delegate).(*Parameters)
	}
	if ctx.TypeName() != nil {
		node.Type = ctx.TypeName().GetText()
	}
	if ctx.Result() != nil {
		node.Result = this.VisitResult(ctx.Result().(*parser.ResultContext), delegate).(*Result)
	}
	return node
}
func (this *OgVisitor) VisitFunctionType(ctx *parser.FunctionTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &FunctionType{
		Node:      NewNode(ctx),
		Signature: this.VisitSignature(ctx.Signature().(*parser.SignatureContext), delegate).(*Signature),
	}
}
func (this *OgVisitor) VisitSignature(ctx *parser.SignatureContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Signature{Node: NewNode(ctx)}
	if ctx.Parameters() != nil {
		node.Parameters = this.VisitParameters(ctx.Parameters().(*parser.ParametersContext), delegate).(*Parameters)
	}
	if ctx.Result() != nil {
		node.Result = this.VisitResult(ctx.Result().(*parser.ResultContext), delegate).(*Result)
	}
	return node
}
func (this *OgVisitor) VisitResult(ctx *parser.ResultContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Result{Node: NewNode(ctx)}
	res := []*Type{}
	bodies := ctx.AllType_()
	for _, spec := range bodies {
		res = append(res, this.VisitType_(spec.(*parser.Type_Context), delegate).(*Type))
	}
	node.Types = res
	return node
}
func (this *OgVisitor) VisitParameters(ctx *parser.ParametersContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.ParameterList() != nil {
		return this.VisitParameterList(ctx.ParameterList().(*parser.ParameterListContext), delegate).(*Parameters)
	}
	return &Parameters{}
}
func (this *OgVisitor) VisitParameterList(ctx *parser.ParameterListContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Parameters{Node: NewNode(ctx)}
	res := []*Parameter{}
	bodies := ctx.AllParameterDecl()
	for _, spec := range bodies {
		res = append(res, this.VisitParameterDecl(spec.(*parser.ParameterDeclContext), delegate).(*Parameter))
	}
	node.List = res
	return node
}
func (this *OgVisitor) VisitParameterDecl(ctx *parser.ParameterDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Parameter{
		Node:       NewNode(ctx),
		Names:      this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).([]string),
		Type:       this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type),
		IsVariadic: ctx.RestOp() != nil,
	}
	return node
}
func (this *OgVisitor) VisitRestOp(ctx *parser.RestOpContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitOperand(ctx *parser.OperandContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Operand{Node: NewNode(ctx)}
	if ctx.Literal() != nil {
		node.Literal = this.VisitLiteral(ctx.Literal().(*parser.LiteralContext), delegate).(*Literal)
	}
	if ctx.OperandName() != nil {
		node.OperandName = this.VisitOperandName(ctx.OperandName().(*parser.OperandNameContext), delegate).(*OperandName)
	}
	if ctx.MethodExpr() != nil {
		node.MethodExpr = this.VisitMethodExpr(ctx.MethodExpr().(*parser.MethodExprContext), delegate).(*MethodExpr)
	}
	if ctx.Expression() != nil {
		node.Expression = this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
	}
	return node
}
func (this *OgVisitor) VisitLiteral(ctx *parser.LiteralContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Literal{Node: NewNode(ctx)}
	if ctx.BasicLit() != nil {
		node.Basic = this.VisitBasicLit(ctx.BasicLit().(*parser.BasicLitContext), delegate).(string)
	}
	if ctx.CompositeLit() != nil {
		node.Composite = this.VisitCompositeLit(ctx.CompositeLit().(*parser.CompositeLitContext), delegate).(*CompositeLit)
	}
	return node
}
func (this *OgVisitor) VisitBasicLit(ctx *parser.BasicLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitOperandName(ctx *parser.OperandNameContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &OperandName{Node: NewNode(ctx)}
	if ctx.This_() != nil {
		node.Name = "this"
	} else if ctx.QualifiedIdent() != nil {
		node.Name = this.VisitQualifiedIdent(ctx.QualifiedIdent().(*parser.QualifiedIdentContext), delegate).(string)
	} else {
		node.Name = ctx.GetText()
	}
	return node
}
func (this *OgVisitor) VisitThis_(ctx *parser.This_Context, delegate antlr.ParseTreeVisitor) interface{} {
	return "this"
}
func (this *OgVisitor) VisitQualifiedIdent(ctx *parser.QualifiedIdentContext, delegate antlr.ParseTreeVisitor) interface{} {
	if ctx.This_() != nil {
		return "this." + ctx.IDENTIFIER(0).GetText()
	}
	return ctx.IDENTIFIER(0).GetText() + "." + ctx.IDENTIFIER(1).GetText()
}
func (this *OgVisitor) VisitCompositeLit(ctx *parser.CompositeLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &CompositeLit{
		Node:         NewNode(ctx),
		LiteralType:  this.VisitLiteralType(ctx.LiteralType().(*parser.LiteralTypeContext), delegate).(*LiteralType),
		LiteralValue: this.VisitLiteralValue(ctx.LiteralValue().(*parser.LiteralValueContext), delegate).(*LiteralValue),
	}
}
func (this *OgVisitor) VisitLiteralType(ctx *parser.LiteralTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &LiteralType{Node: NewNode(ctx)}
	if ctx.StructType() != nil {
		node.Struct = this.VisitStructType(ctx.StructType().(*parser.StructTypeContext), delegate).(*StructType)
	}
	if ctx.ArrayType() != nil {
		node.Array = this.VisitArrayType(ctx.ArrayType().(*parser.ArrayTypeContext), delegate).(*ArrayType)
	}
	if ctx.ElementType() != nil {
		node.Element = this.VisitElementType(ctx.ElementType().(*parser.ElementTypeContext), delegate).(*Type)
	}
	if ctx.SliceType() != nil {
		node.Slice = this.VisitSliceType(ctx.SliceType().(*parser.SliceTypeContext), delegate).(*SliceType)
	}
	if ctx.MapType() != nil {
		node.Map = this.VisitMapType(ctx.MapType().(*parser.MapTypeContext), delegate).(*MapType)
	}
	if ctx.TypeName() != nil {
		node.Type = this.VisitTypeName(ctx.TypeName().(*parser.TypeNameContext), delegate).(string)
	}
	return node
}
func (this *OgVisitor) VisitLiteralValue(ctx *parser.LiteralValueContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &LiteralValue{Node: NewNode(ctx)}
	if ctx.ElementList() != nil {
		node.Elements = this.VisitElementList(ctx.ElementList().(*parser.ElementListContext), delegate).([]*KeyedElement)
	}
	return node
}
func (this *OgVisitor) VisitElementList(ctx *parser.ElementListContext, delegate antlr.ParseTreeVisitor) interface{} {
	res := []*KeyedElement{}
	bodies := ctx.AllKeyedElement()
	for _, spec := range bodies {
		res = append(res, this.VisitKeyedElement(spec.(*parser.KeyedElementContext), delegate).(*KeyedElement))
	}
	return res
}
func (this *OgVisitor) VisitKeyedElement(ctx *parser.KeyedElementContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &KeyedElement{
		Node:    NewNode(ctx),
		Element: this.VisitElement(ctx.Element().(*parser.ElementContext), delegate).(*Element),
	}
	if ctx.Key() != nil {
		node.Key = this.VisitKey(ctx.Key().(*parser.KeyContext), delegate).(*Key)
	}
	return node
}
func (this *OgVisitor) VisitKey(ctx *parser.KeyContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Key{Node: NewNode(ctx)}
	if ctx.IDENTIFIER() != nil {
		node.Name = ctx.IDENTIFIER().GetText()
	}
	if ctx.Expression() != nil {
		node.Expression = this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
	}
	if ctx.LiteralValue() != nil {
		node.LiteralValue = this.VisitLiteralValue(ctx.LiteralValue().(*parser.LiteralValueContext), delegate).(*LiteralValue)
	}
	return node
}
func (this *OgVisitor) VisitElement(ctx *parser.ElementContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Element{Node: NewNode(ctx)}
	if ctx.Expression() != nil {
		node.Expression = this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression)
	}
	if ctx.LiteralValue() != nil {
		node.LiteralValue = this.VisitLiteralValue(ctx.LiteralValue().(*parser.LiteralValueContext), delegate).(*LiteralValue)
	}
	return node
}
func (this *OgVisitor) VisitStructType(ctx *parser.StructTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &StructType{Node: NewNode(ctx)}
	if ctx.IDENTIFIER() != nil {
		node.Name = ctx.IDENTIFIER().GetText()
	}
	res := []*FieldDecl{}
	bodies := ctx.AllFieldDecl()
	for _, spec := range bodies {
		res = append(res, this.VisitFieldDecl(spec.(*parser.FieldDeclContext), delegate).(*FieldDecl))
	}
	node.Fields = res
	return node
}
func (this *OgVisitor) VisitFieldDecl(ctx *parser.FieldDeclContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &FieldDecl{Node: NewNode(ctx)}
	if ctx.IdentifierList() != nil {
		node.Identifiers = this.VisitIdentifierList(ctx.IdentifierList().(*parser.IdentifierListContext), delegate).([]string)
	}
	if ctx.Type_() != nil {
		node.Type = this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type)
	}
	if ctx.AnonymousField() != nil {
		node.Anonymous = this.VisitAnonymousField(ctx.AnonymousField().(*parser.AnonymousFieldContext), delegate).(*AnonymousField)
	}
	if ctx.STRING_LIT() != nil {
		node.Tag = ctx.STRING_LIT().GetText()
	}
	if ctx.InlineStructMethod() != nil {
		node.InlineStructMethod = this.VisitInlineStructMethod(ctx.InlineStructMethod().(*parser.InlineStructMethodContext), delegate).(*InlineStructMethod)
	}
	return node
}
func (this *OgVisitor) VisitInlineStructMethod(ctx *parser.InlineStructMethodContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &InlineStructMethod{
		Node:              NewNode(ctx),
		FunctionDecl:      this.VisitFunctionDecl(ctx.FunctionDecl().(*parser.FunctionDeclContext), delegate).(*FunctionDecl),
		IsPointerReceiver: strings.Contains(ctx.GetText(), "*"),
	}
}
func (this *OgVisitor) VisitAnonymousField(ctx *parser.AnonymousFieldContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &AnonymousField{
		Node:              NewNode(ctx),
		Type:              this.VisitTypeName(ctx.TypeName().(*parser.TypeNameContext), delegate).(string),
		IsPointerReceiver: strings.Contains(ctx.GetText(), "*"),
	}
}
func (this *OgVisitor) VisitFunctionLit(ctx *parser.FunctionLitContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &FunctionLit{
		Node:     NewNode(ctx),
		Function: this.VisitFunction(ctx.Function().(*parser.FunctionContext), delegate).(*Function),
	}
}
func (this *OgVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &PrimaryExpr{Node: NewNode(ctx)}
	if ctx.Operand() != nil {
		node.Operand = this.VisitOperand(ctx.Operand().(*parser.OperandContext), delegate).(*Operand)
	}
	if ctx.Conversion() != nil {
		node.Conversion = this.VisitConversion(ctx.Conversion().(*parser.ConversionContext), delegate).(*Conversion)
	}
	if ctx.PrimaryExpr() != nil {
		node.PrimaryExpr = this.VisitPrimaryExpr(ctx.PrimaryExpr().(*parser.PrimaryExprContext), delegate).(*PrimaryExpr)
	}
	if ctx.Selector() != nil {
		node.Selector = this.VisitSelector(ctx.Selector().(*parser.SelectorContext), delegate).(string)
	}
	if ctx.Index() != nil {
		node.Index = this.VisitIndex(ctx.Index().(*parser.IndexContext), delegate).(*Index)
	}
	if ctx.Slice() != nil {
		node.Slice = this.VisitSlice(ctx.Slice().(*parser.SliceContext), delegate).(*Slice)
	}
	if ctx.TypeAssertion() != nil {
		node.TypeAssertion = this.VisitTypeAssertion(ctx.TypeAssertion().(*parser.TypeAssertionContext), delegate).(*TypeAssertion)
	}
	if ctx.Arguments() != nil {
		node.Arguments = this.VisitArguments(ctx.Arguments().(*parser.ArgumentsContext), delegate).(*Arguments)
	}
	return node
}
func (this *OgVisitor) VisitSelector(ctx *parser.SelectorContext, delegate antlr.ParseTreeVisitor) interface{} {
	return "." + ctx.IDENTIFIER().GetText()
}
func (this *OgVisitor) VisitIndex(ctx *parser.IndexContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &Index{
		Node:       NewNode(ctx),
		Expression: this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression),
	}
}
func (this *OgVisitor) VisitSlice(ctx *parser.SliceContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Slice{Node: NewNode(ctx)}
	dist := strings.Split(ctx.GetText(), ":")
	if len(dist) == 3 {
		if len(dist[0]) <= 1 {
			node.RightExpr = this.VisitExpression(ctx.Expression(1).(*parser.ExpressionContext), delegate).(*Expression)
			node.MiddleExpr = this.VisitExpression(ctx.Expression(0).(*parser.ExpressionContext), delegate).(*Expression)
		} else {
			node.RightExpr = this.VisitExpression(ctx.Expression(2).(*parser.ExpressionContext), delegate).(*Expression)
			node.MiddleExpr = this.VisitExpression(ctx.Expression(1).(*parser.ExpressionContext), delegate).(*Expression)
			if ctx.Expression(0) != nil {
				node.LeftExpr = this.VisitExpression(ctx.Expression(0).(*parser.ExpressionContext), delegate).(*Expression)
			}
		}
	} else {
		if len(dist[0]) <= 1 {
			if ctx.Expression(0) != nil {
				node.MiddleExpr = this.VisitExpression(ctx.Expression(0).(*parser.ExpressionContext), delegate).(*Expression)
			}
		} else {
			if ctx.Expression(0) != nil {
				node.LeftExpr = this.VisitExpression(ctx.Expression(0).(*parser.ExpressionContext), delegate).(*Expression)
			}
			if ctx.Expression(1) != nil {
				node.MiddleExpr = this.VisitExpression(ctx.Expression(1).(*parser.ExpressionContext), delegate).(*Expression)
			}
		}
	}
	return node
}
func (this *OgVisitor) VisitTypeAssertion(ctx *parser.TypeAssertionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &TypeAssertion{
		Node: NewNode(ctx),
		Type: this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type),
	}
}
func (this *OgVisitor) VisitArguments(ctx *parser.ArgumentsContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Arguments{
		Node:       NewNode(ctx),
		IsVariadic: ctx.RestOp() != nil,
	}
	if ctx.Type_() != nil {
		node.Type = this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type)
	}
	if ctx.ExpressionList() != nil {
		node.Expressions = this.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext), delegate).(*ExpressionList)
	}
	return node
}
func (this *OgVisitor) VisitMethodExpr(ctx *parser.MethodExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &MethodExpr{
		Node:         NewNode(ctx),
		ReceiverType: this.VisitReceiverType(ctx.ReceiverType().(*parser.ReceiverTypeContext), delegate).(*ReceiverType),
		Name:         ctx.IDENTIFIER().GetText(),
	}
}
func (this *OgVisitor) VisitReceiverType(ctx *parser.ReceiverTypeContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &ReceiverType{
		Node:      NewNode(ctx),
		IsPointer: strings.Contains(ctx.GetText(), "*"),
	}
	if ctx.TypeName() != nil {
		node.Type = this.VisitTypeName(ctx.TypeName().(*parser.TypeNameContext), delegate).(string)
	}
	if ctx.ReceiverType() != nil {
		node.ReceiverType = this.VisitReceiverType(ctx.ReceiverType().(*parser.ReceiverTypeContext), delegate).(*ReceiverType)
	}
	return node
}
func (this *OgVisitor) VisitExpression(ctx *parser.ExpressionContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Expression{Node: NewNode(ctx)}
	if ctx.UnaryExpr() != nil {
		node.UnaryExpr = this.VisitUnaryExpr(ctx.UnaryExpr().(*parser.UnaryExprContext), delegate).(*UnaryExpr)
	}
	if ctx.FunctionLit() != nil {
		node.FunctionLit = this.VisitFunctionLit(ctx.FunctionLit().(*parser.FunctionLitContext), delegate).(*FunctionLit)
	}
	if ctx.Expression(0) != nil {
		node.LeftExpression = this.VisitExpression(ctx.Expression(0).(*parser.ExpressionContext), delegate).(*Expression)
	}
	if ctx.Binary_op() != nil {
		node.Op = ctx.Binary_op().GetText()
	}
	if ctx.Expression(1) != nil {
		node.RightExpression = this.VisitExpression(ctx.Expression(1).(*parser.ExpressionContext), delegate).(*Expression)
	}
	return node
}
func (this *OgVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &UnaryExpr{Node: NewNode(ctx)}
	if ctx.PrimaryExpr() != nil {
		node.PrimaryExpr = this.VisitPrimaryExpr(ctx.PrimaryExpr().(*parser.PrimaryExprContext), delegate).(*PrimaryExpr)
	}
	if ctx.Unary_op() != nil {
		node.Op = ctx.Unary_op().GetText()
	}
	if ctx.UnaryExpr() != nil {
		node.UnaryExpr = this.VisitUnaryExpr(ctx.UnaryExpr().(*parser.UnaryExprContext), delegate).(*UnaryExpr)
	}
	return node
}
func (this *OgVisitor) VisitUnary_op(ctx *parser.Unary_opContext, delegate antlr.ParseTreeVisitor) interface{} {
	return ctx.GetText()
}
func (this *OgVisitor) VisitConversion(ctx *parser.ConversionContext, delegate antlr.ParseTreeVisitor) interface{} {
	return &Conversion{
		Node:       NewNode(ctx),
		Type:       this.VisitType_(ctx.Type_().(*parser.Type_Context), delegate).(*Type),
		Expression: this.VisitExpression(ctx.Expression().(*parser.ExpressionContext), delegate).(*Expression),
	}
}
func (this *OgVisitor) VisitInterp(ctx *parser.InterpContext, delegate antlr.ParseTreeVisitor) interface{} {
	node := &Interpret{Node: NewNode(ctx)}
	if ctx.Statement() != nil {
		node.Statement = this.VisitStatement(ctx.Statement().(*parser.StatementContext), delegate).(*Statement)
	}
	if ctx.TopLevelDecl() != nil {
		node.TopLevel = this.VisitTopLevelDecl(ctx.TopLevelDecl().(*parser.TopLevelDeclContext), delegate).(*TopLevel)
	}
	return node
}
