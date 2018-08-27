package main

import (
	"fmt"
	"strings"
)

func parseAst(ast *INI) string {
	return fmt.Sprint(
		Pack(ast.Pack),
		Imports(ast.Imports),
		TopLevels(ast.TopLevel),
	)
}

func Pack(pack string) string {
	return fmt.Sprintln("package", pack)
}

func Imports(imports []string) string {
	res := []string{fmt.Sprintln("import (")}

	for _, imp := range imports {
		res = append(res, fmt.Sprint("\"", imp, "\"\n"))
	}

	res = append(res, fmt.Sprintln(")"))

	return strings.Join(res, "")
}

func TopLevels(top []*TopLevel) string {
	res := []string{}

	for _, t := range top {
		for _, s := range t.Structs {
			res = append(res, Struct_(s))
		}
		for _, f := range t.Funcs {
			res = append(res, Func_(f))
		}
	}

	return strings.Join(res, "")
}

func Struct_(s *Struct) string {
	res := []string{fmt.Sprintln("type", s.Name, "struct {")}

	for _, field := range s.Fields {
		res = append(res, fmt.Sprintln(field.Name, field.Type))
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

	res = append(res, ") {\n")

	for _, stmt := range s.Body {
		res = append(res, Stmt_(stmt))
	}

	res = append(res, "}\n")

	return strings.Join(res, "")
}

func Arg_(a *Arg) string {
	return fmt.Sprint(a.Name, " ", a.Type)
}

func Stmt_(s *Stmt) string {
	res := []string{NestedProperty_(s.Ident)}

	if s.VarDecl != nil {
		res = append(res, fmt.Sprintln(":=", Value_(s.VarDecl.Value)))
	}

	if s.FuncCall != nil {
		res = append(res, fmt.Sprintln(FuncCall_(s.FuncCall)))
	}

	return strings.Join(res, "")
}

func Value_(v *Value) string {
	if len(v.String) > 0 {
		return fmt.Sprint("\"", v.String, "\"")
	}

	if v.Ident != nil {
		return NestedProperty_(v.Ident)
	}

	return fmt.Sprint(v.Number)
}

func NestedProperty_(n *NestedProperty) string {
	res := []string{n.Ident}

	if n.Nested != nil {
		res = append(res, ".")
		res = append(res, NestedProperty_(n.Nested))
	}

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
