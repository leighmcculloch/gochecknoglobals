package gochecknoglobals

import (
	"flag"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckNoGlobals(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, Analyzer, "globalcheck")
}

func TestIncludeTests(t *testing.T) {
	// This is not yet testable due to an issue with analysistest.
	// Issue created at https://github.com/golang/go/issues/41771
	// The test is workign but also adds false positives from the cache so it
	// must be skipped for now.
	t.Skip("Not yet testable")

	testdata := analysistest.TestData()
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.Bool("t", false, "")

	analyzer := Analyzer
	analyzer.Flags = *flags

	analysistest.RunWithSuggestedFixes(t, testdata, analyzer, "includetest")
}
