package walker

import (
	"github.com/champii/og/lib/common"
)

type FuncSig struct {
	name      string
	arguments map[string]string
	returns   map[string]string
}

func NewFuncSig() *FuncSig {
	return &FuncSig{
		arguments: make(map[string]string),
		returns:   make(map[string]string),
	}
}

type Scope struct {
	vars  map[string]string
	funcs map[string]*FuncSig
}

func NewScope() *Scope {
	return &Scope{vars: make(map[string]string)}
}

type Stack struct {
	scopes []*Scope
}

func (this *Stack) PushScope() {
	this.scopes = append([]*Scope{NewScope()}, this.scopes...)
}
func (this *Stack) PopScope() {
	this.scopes = this.scopes[1:]
}
func (this *Stack) AddVar(name, t string) bool {
	if _, ok := this.GetVar(name); ok {
		return false
	}
	this.scopes[0].vars[name] = t
	return true
}
func (this *Stack) GetVar(name string) (string, bool) {
	for _, scope := range this.scopes {
		if t, ok := scope.vars[name]; ok {
			return t, true
		}
	}
	return "", false
}
func (this *Stack) AddFunc(name string, f *FuncSig) bool {
	if r := this.GetFunc(name); r != nil {
		return false
	}
	this.scopes[0].funcs[name] = f
	return true
}
func (this *Stack) GetFunc(name string) *FuncSig {
	for _, scope := range this.scopes {
		if f, ok := scope.funcs[name]; ok {
			return f
		}
	}
	return nil
}

type TypeChecker struct {
	AstWalker
	stack *Stack
}

func (this *TypeChecker) VarSpec(n common.INode) common.INode {
	return n
}
func (this *TypeChecker) Assignment(n common.INode) common.INode {
	return n
}
func (this *TypeChecker) BeforeBlock(n common.INode) {
	this.stack.PushScope()
}
func (this *TypeChecker) AfterBlock(n common.INode) {
	this.stack.PopScope()
}
func (this *TypeChecker) Each(n common.INode) common.INode {
	return n
}
func TypeCheck(ast common.INode) {
	t := TypeChecker{stack: &Stack{}}
	t.stack.PushScope()
	t.type_ = &t
	t.Walk(ast)
}
