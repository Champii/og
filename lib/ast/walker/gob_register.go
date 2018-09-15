package walker

import (
	"encoding/gob"
	"github.com/champii/og/lib/ast"
)

type GobRegister struct {
	AstWalker
}

func (this *GobRegister) Each(n ast.INode) {
	gob.Register(n)
}
func RunGobRegister(tree ast.INode) {
	p := GobRegister{}
	gob.Register(tree)
	p.type_ = &p
	p.Walk(tree)
}
