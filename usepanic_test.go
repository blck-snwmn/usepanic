package usepanic

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyze(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "main", "other", "foo")
}
