package checknoglobals

import (
	"flag"
	"strconv"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckNoGlobals(t *testing.T) {
	testdata := analysistest.TestData()
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.Bool("t", true, "")

	analyzer := Analyzer()
	analyzer.Flags = *flags

	for i := 0; i <= 10; i++ {
		dir := strconv.Itoa(i)
		t.Run(dir, func(t *testing.T) {
			analysistest.Run(t, testdata, analyzer, dir)
		})
	}
}
