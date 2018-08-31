package translator

import "Og/parser"

type Translator struct {
	parser.BaseGolangListener

	Out []string
}

func (this *Translator) Push(str string) {
	this.Out = append(this.Out, str)
}

func (this *Translator) ExitPackageClause(ctx *parser.PackageClauseContext) {
	this.Push("package " + ctx.IDENTIFIER().GetText() + "\n")
}

func (this *Translator) EnterImportDecl(ctx *parser.ImportDeclContext) {
	this.Push("import (\n")
}

func (this *Translator) ExitImportDecl(ctx *parser.ImportDeclContext) {
	this.Push(")\n")
}

func (this *Translator) ExitImportSpec(ctx *parser.ImportSpecContext) {
	this.Push("\n")
	//TODO: manage renamed export
}

func (this *Translator) ExitImportPath(ctx *parser.ImportPathContext) {
	txt := ctx.GetText()

	if txt[0] == '"' {
		this.Push(txt)
	} else {
		this.Push("\"" + txt + "\"")
	}
}

// func (this *Translator) ExitTopLevelDecl(ctx *parser.TopLevelDeclContext) {
// 	this.Push(ctx.GetText())
// }

func (this *Translator) EnterFunctionDecl(ctx *parser.FunctionDeclContext) {
	this.Push("func")
	this.Push(ctx.IDENTIFIER().GetText())
}

func (this *Translator) EnterParameters(ctx *parser.ParametersContext) {
	this.Push("(")
}

func (this *Translator) ExitParameters(ctx *parser.ParametersContext) {
	this.Push(")")
}

func (this *Translator) ExitParameterDecl(ctx *parser.ParameterDeclContext) {
	this.Push(",")
}

func (this *Translator) ExitIdentifierList(ctx *parser.IdentifierListContext) {
	this.Push(ctx.GetText())
}

func (this *Translator) ExitType_(ctx *parser.Type_Context) {
	this.Push(ctx.GetText())
}

func (this *Translator) EnterBlock(ctx *parser.BlockContext) {
	this.Push("{\n")
}

func (this *Translator) ExitBlock(ctx *parser.BlockContext) {
	this.Push("}\n")
}

func (this *Translator) EnterOperand(ctx *parser.OperandContext) {
	this.Push(ctx.GetText())
}

func (this *Translator) EnterSelector(ctx *parser.SelectorContext) {
	this.Push(ctx.GetText())
}

func (this *Translator) EnterArguments(ctx *parser.ArgumentsContext) {
	this.Push("(")
}

func (this *Translator) ExitArguments(ctx *parser.ArgumentsContext) {
	this.Push(")")
}
