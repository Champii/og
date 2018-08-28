package og

import (
	"fmt"
	"strings"
)

func parseAst(ast *INI) string {
	return fmt.Sprint(
		Pack(ast.Pack),
		ProgBody_(ast.ProgBody),
	)
}

func ProgBody_(pbody *ProgBody) string {
	if pbody == nil {
		return ""
	}

	return fmt.Sprint(
		Imports(pbody.Imports),
		TopLevels(pbody.TopLevel),
	)
}

func Pack(pack string) string {
	return fmt.Sprintln("package", pack)
}

func Imports(imports []string) string {
	if len(imports) == 0 {
		return "\n"
	}

	res := []string{fmt.Sprintln("import (")}

	for _, imp := range imports {
		res = append(res, fmt.Sprint("\"", imp, "\"\n"))
	}

	res = append(res, fmt.Sprintln(")"))

	return strings.Join(res, "")
}

func TopLevels(top []*TopLevel) string {
	if len(top) == 0 {
		return ""
	}

	res := []string{}

	for _, t := range top {
		if t.Struct != nil {
			res = append(res, Struct_(t.Struct))
		}
		if t.Func != nil {
			res = append(res, Func_(t.Func))
		}
		res = append(res, "\n")
	}

	return strings.Join(res, "")
}

func Struct_(s *Struct) string {
	res := []string{fmt.Sprintln("type", s.Name, "struct {")}

	for _, field := range s.Fields {
		res = append(res, fmt.Sprint(field.Name, " ", field.Type))

		if field.Tag != nil {
			res = append(res, fmt.Sprint("`", *field.Tag, "`"))
		}

		res = append(res, "\n")
	}

	res = append(res, "}\n")

	return strings.Join(res, "")
}

func Func_(s *Func) string {
	res := []string{fmt.Sprint("func ", s.Name, "(")}

	for _, arg := range s.Args {
		res = append(res, Arg_(arg))
		res = append(res, ",")
	}

	if len(res) > 1 {
		res = res[0 : len(res)-1]
	}

	res = append(res, ") ")

	if s.ReturnType != nil {
		res = append(res, Type_(s.ReturnType))
	}

	res = append(res, "{\n")

	for _, stmt := range s.Body {
		res = append(res, Stmt_(stmt))
	}

	res = append(res, "}")

	return strings.Join(res, "")
}

func Type_(t *Type) string {
	res := []string{}

	for _, a := range t.Array {
		res = append(res, a)
	}

	res = append(res, t.Type)

	return strings.Join(res, "")
}

func Arg_(a *Arg) string {
	res := []string{fmt.Sprint(a.Name, " ")}

	res = append(res, Type_(a.Type))

	return strings.Join(res, "")
}

func Stmt_(s *Stmt) string {
	if s.IdentOrVarDecl != nil {
		return fmt.Sprintln(IdentOrVarDecl_(s.IdentOrVarDecl))
	}

	if s.If != nil {
		return fmt.Sprintln(If_(s.If))
	}

	if s.For != nil {
		return fmt.Sprintln(For_(s.For))
	}

	if s.GoRoutine != nil {
		return fmt.Sprintln(GoRoutine_(s.GoRoutine))
	}

	if s.Return != nil {
		return fmt.Sprintln("return ", Value_(s.Return))
	}

	return ""
}

func For_(f *For) string {
	res := []string{"for "}

	res = append(res, f.Iterator)

	if len(f.Value) > 0 {
		res = append(res, fmt.Sprint(", ", f.Value))
	}

	res = append(res, " := range ")
	res = append(res, f.Source)
	res = append(res, "{\n")

	for _, stmt := range f.Body {
		res = append(res, Stmt_(stmt))
	}

	res = append(res, "}\n")

	return strings.Join(res, "")
}

func If_(i *If) string {
	res := []string{"if "}

	res = append(res, Predicat_(i.Predicat))

	res = append(res, "{\n")

	for _, stmt := range i.Body {
		res = append(res, Stmt_(stmt))
	}

	if i.ElseIf != nil {
		if i.ElseIf.If != nil {
			res = append(res, "} else ")
			res = append(res, If_(i.ElseIf.If))
		} else if i.ElseIf.Else != nil {
			res = append(res, Else_(i.ElseIf.Else))
		}
	} else {
		res = append(res, "}\n")
	}

	return strings.Join(res, "")
}

func Else_(e *Else) string {
	res := []string{"} else "}

	res = append(res, "{\n")

	for _, stmt := range e.Body {
		res = append(res, Stmt_(stmt))
	}

	res = append(res, "}\n")

	return strings.Join(res, "")
}

func Predicat_(p *Predicat) string {
	res := []string{}

	res = append(res, Value_(p.First))
	res = append(res, PredicatOperator_(p.Operator))
	res = append(res, Value_(p.Second))

	return strings.Join(res, "")
}
func PredicatOperator_(o *PredicatOperator) string {
	if len(o.Eq) > 0 {
		return "=="
	}

	if len(o.Neq) > 0 {
		return "!="
	}

	if len(o.Gt) > 0 {
		return ">"
	}

	if len(o.Gte) > 0 {
		return ">="
	}

	if len(o.Lt) > 0 {
		return "<"
	}

	if len(o.Lte) > 0 {
		return "<="
	}

	return ""
}

func IdentOrVarDecl_(s *IdentOrVarDecl) string {
	res := []string{NestedProperty_(s.Ident)}

	if s.VarDecl != nil {
		res = append(res, fmt.Sprintln(":=", Value_(s.VarDecl.Value)))
	}

	return strings.Join(res, "")
}

func Value_(v *Value) string {
	if v.Bool != nil {
		return fmt.Sprint(*v.Bool)
	}

	if v.String != nil {
		return fmt.Sprint("\"", *v.String, "\"")
	}

	if v.Operation != nil {
		return Operation_(v.Operation)
	}

	if v.NestedProperty != nil {
		return NestedProperty_(v.NestedProperty)
	}

	if v.ArrDecl != nil {
		return ArrDecl_(v.ArrDecl)
	}

	return ""
}

func Number_(n *Number) string {
	if n.Int != nil {
		return fmt.Sprint(*n.Int)
	}

	if n.Float != nil {
		return fmt.Sprint(*n.Float)
	}

	return ""
}

func Operation_(o *Operation) string {
	res := []string{Number_(o.First)}

	if o.Op != nil {
		res = append(res, Operator_(o.Op))
	}

	if o.Nested != nil {
		res = append(res, Operation_(o.Nested))
	}

	return strings.Join(res, "")
}

func Operator_(o *Operator) string {
	if o.Plus != nil {
		return *o.Plus
	}

	if o.Less != nil {
		return *o.Less
	}

	if o.Times != nil {
		return *o.Times
	}

	if o.Div != nil {
		return *o.Div
	}

	if o.Mod != nil {
		return *o.Mod
	}

	return ""
}

func NestedProperty_(n *NestedProperty) string {
	res := []string{n.Ident}

	for _, a := range n.ArrAccessOrFuncCall {
		res = append(res, ArrAccessOrFuncCall_(a))
	}

	if n.Nested != nil {
		res = append(res, ".")
		res = append(res, NestedProperty_(n.Nested))
	}

	return strings.Join(res, "")
}

func ArrAccessOrFuncCall_(a *ArrAccessOrFuncCall) string {
	if a.ArrAccess != nil {
		return ArrAccess_(a.ArrAccess)
	}

	if a.FuncCall != nil {
		return FuncCall_(a.FuncCall)
	}

	return ""
}

func ArrAccess_(a *ArrAccess) string {
	res := []string{"["}

	res = append(res, Value_(a.Value))

	res = append(res, "]")

	return strings.Join(res, "")
}

func FuncCall_(f *FuncCall) string {
	res := []string{"("}

	for _, arg := range f.Args {
		res = append(res, Value_(arg))
		res = append(res, ",")
	}

	if len(res) > 1 {
		res = res[0 : len(res)-1]
	}

	res = append(res, ")")

	return strings.Join(res, "")
}

func ArrDecl_(a *ArrDecl) string {
	res := []string{fmt.Sprint("[]", a.Type, "{\n")}

	for _, arg := range a.Values {
		res = append(res, Value_(arg))
		res = append(res, ",\n")
	}

	res = append(res, "}")

	return strings.Join(res, "")
}

func GoRoutine_(g *GoRoutine) string {
	res := []string{"go "}

	if g.Func != nil {
		res = append(res, Func_(g.Func))
		res = append(res, "()")
	}

	if g.Value != nil {
		res = append(res, Value_(g.Value))
	}

	res = append(res, "\n")

	return strings.Join(res, "")
}
