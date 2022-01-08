package usepanic

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// Analyzer check if panic is used except for the specified file
var Analyzer = &analysis.Analyzer{
	Name:     "usepanic",
	Doc:      "check if panic is used except for the specified file",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
