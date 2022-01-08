package main

import (
	"github/blck-snwmn/usepanic"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(usepanic.Analyzer)
}
