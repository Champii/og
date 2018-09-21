package walker

import (
	"encoding/gob"
	"github.com/champii/og/lib/common"
)

type GobRegister struct {
	AstWalker
}

func (this *GobRegister) Each(n common.INode) {
	gob.Register(n)
}
func RunGobRegister(tree common.INode) {
	p := GobRegister{}
	gob.Register(tree)
	p.type_ = &p
	p.Walk(tree)
}
