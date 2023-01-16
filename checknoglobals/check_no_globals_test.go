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

	for i := 0; i <= 11; i++ {
		dir := strconv.Itoa(i)
		t.Run(dir, func(t *testing.T) {
			analysistest.Run(t, testdata, analyzer, dir)
		})
	}
}

func BenchmarkRun(b *testing.B) {
	analyzer := Analyzer()
	flags := flag.NewFlagSet("", flag.ExitOnError)
	flags.Bool("t", true, "")
	analyzer.Flags = *flags
	testdata := analysistest.TestData()
	results := analysistest.Run(b, testdata, analyzer, "./...")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, r := range results {
			_, err := analyzer.Run(r.Pass)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
