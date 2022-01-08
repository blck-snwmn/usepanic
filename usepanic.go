package usepanic

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer check if panic is used except for the specified package
var Analyzer = &analysis.Analyzer{
	Name:     "usepanic",
	Doc:      "check if panic is used except for the specified package",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// TODO use command line args
	okPackage := map[string]struct{}{"main": {}}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.File)(nil),
		(*ast.Ident)(nil),
	}
	inspect.Nodes(nodeFilter, func(n ast.Node, push bool) (proceed bool) {
		// only first visit
		if !push {
			return false
		}
		switch n := n.(type) {
		case *ast.File:
			packageName := n.Name.String()
			if _, ok := okPackage[packageName]; ok {
				// if packageName exsit in `okPackage`, visit no children node.
				return false
			}
		case *ast.Ident:
			if n.Name == "panic" {
				pass.Reportf(n.Pos(), "don't use `panic` in this package.")
			}
		}
		return true
	})
	return nil, nil
}
