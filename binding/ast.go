package binding

import (
	"go/ast"
	"go/parser"
	"go/token"
)

var currentPkg string

func getASTFromFile(path string) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return
	}
	currentPkg = file.Name.Name
	ast.Walk(&bindingHelper, file)
}
