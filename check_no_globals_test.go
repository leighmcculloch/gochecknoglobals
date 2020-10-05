package gochecknoglobals

import (
	"flag"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

// testfilesWorking is a constant to symbolise that we cannot yet run
// analysistest on test files. This is because the test files will get compiled
// to files containing global variables and thus the test cannot pass.
// I've confirmed by running this manually that it is in fact working, however
// the extra diagnostics we get makes the test not able to pass. Until this is
// fixed we cannot run tests on any package containing tests.
//
// An upstream issue has been filed:
// https://github.com/golang/go/issues/41771
const testfilesWorking = false

func TestCheckNoGlobals(t *testing.T) {
	testdata := analysistest.TestData()
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.Bool("t", true, "")

	analyzer := Analyzer
	analyzer.Flags = *flags

	if testfilesWorking {
		analysistest.Run(t, testdata, analyzer, "1")
		analysistest.Run(t, testdata, analyzer, "2")
	}

	analysistest.Run(t, testdata, analyzer, "0")
	analysistest.Run(t, testdata, analyzer, "3")
	analysistest.Run(t, testdata, analyzer, "4")
	analysistest.Run(t, testdata, analyzer, "5")
	analysistest.Run(t, testdata, analyzer, "6")
	analysistest.Run(t, testdata, analyzer, "7")
	analysistest.Run(t, testdata, analyzer, "8")
	analysistest.Run(t, testdata, analyzer, "9")
	analysistest.Run(t, testdata, analyzer, "10")
}
