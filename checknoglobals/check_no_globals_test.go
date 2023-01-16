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
	dir, cleanup, err := analysistest.WriteFiles(map[string]string{
		"file.go": `package code
		import "errors"
		var global = "" // want "global is a global variable"
		var ErrVar = errors.New("myErrVar")
		var myErrVar = errors.New("myErrVar") // want "myErrVar is a global variable"
		var errCustom = customError{}
		type customError struct {}
		func (customError) Error() string { return "custom error" }`,
	})
	if err != nil {
		b.Fatal(err)
	}
	defer cleanup()
	results := analysistest.Run(b, dir, analyzer, "./...")
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
