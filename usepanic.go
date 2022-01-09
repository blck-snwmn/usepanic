package usepanic

import (
	"go/ast"
	"strings"

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

type allowPackagesFlags map[string]struct{}

func (apf *allowPackagesFlags) String() string {
	elem := make([]string, 0, len(*apf))
	for k, _ := range *apf {
		elem = append(elem, k)
	}
	return strings.Join(elem, ", ")
}
func (apf *allowPackagesFlags) Set(s string) error {
	t := map[string]struct{}{}
	for _, e := range strings.Split(s, ",") {
		if e == "" {
			continue
		}
		t[e] = struct{}{}
	}
	*apf = t
	return nil
}

var allowPackages allowPackagesFlags

func init() {
	Analyzer.Flags.Var(&allowPackages, "p", "packages allowed to use `panic`")
}

func run(pass *analysis.Pass) (interface{}, error) {
	if len(allowPackages) == 0 {
		allowPackages.Set("main")
	}

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
			if _, ok := allowPackages[packageName]; ok {
				// if packageName exsit in `okPackage`, visit no children node.
				return false
			}
		case *ast.Ident:
			if n.Name == "panic" {
				// report if name is `panic` whether or not built-in function
				pass.Reportf(n.Pos(), "don't use `panic` in this package.")
			}
		}
		return true
	})
	return nil, nil
}
