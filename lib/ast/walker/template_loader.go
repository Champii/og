package walker

import (
	"github.com/champii/og/lib/ast"
	"github.com/champii/og/lib/common"
	"io/ioutil"
	"os"
	"path"
)

func load(dir string, templates *Templates) {
	templateDir := path.Join(os.Getenv("GOPATH"), "src", dir, ".og")
	_, err := os.Stat(templateDir)
	if err != nil {
		return
	}
	content, err := ioutil.ReadFile(path.Join(templateDir, "template"))
	if err != nil {
		return
	}
	templates.Decode(content)
}
func RunTemplateLoader(tree common.INode, templates *Templates) {
	imports := tree.(*ast.SourceFile).Import
	if imports == nil {
		return
	}
	for _, imp := range imports.Items {
		load(imp.Path[1:len(imp.Path)-2], templates)
	}
}
