package gochecknoglobals

import (
	"flag"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckNoGlobals(t *testing.T) {
	testdata := analysistest.TestData()
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.Bool("t", true, "")

	analyzer := Analyzer()
	analyzer.Flags = *flags

	analysistest.Run(t, testdata, analyzer, "0")
	analysistest.Run(t, testdata, analyzer, "1")
	analysistest.Run(t, testdata, analyzer, "2")
	analysistest.Run(t, testdata, analyzer, "3")
	analysistest.Run(t, testdata, analyzer, "4")
	analysistest.Run(t, testdata, analyzer, "5")
	analysistest.Run(t, testdata, analyzer, "6")
	analysistest.Run(t, testdata, analyzer, "7")
	analysistest.Run(t, testdata, analyzer, "8")
	analysistest.Run(t, testdata, analyzer, "9")
	analysistest.Run(t, testdata, analyzer, "10")
}
