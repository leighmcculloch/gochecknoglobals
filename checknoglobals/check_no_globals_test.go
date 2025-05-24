package checknoglobals

import (
	"strconv"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckNoGlobals(t *testing.T) {
	testdata := analysistest.TestData()
	
	// Keep track of the original value to restore it
	originalCheckGlobalDeclarations := CheckGlobalDeclarations
	defer func() {
		CheckGlobalDeclarations = originalCheckGlobalDeclarations
	}()

	// Run tests 0-11 with the original behavior (checking declarations)
	t.Run("OldBehavior", func(t *testing.T) {
		// Set the global flag to use the old behavior
		CheckGlobalDeclarations = true
		
		analyzer := Analyzer()
		
		for i := 0; i <= 11; i++ {
			dir := strconv.Itoa(i)
			t.Run(dir, func(t *testing.T) {
				analysistest.Run(t, testdata, analyzer, dir)
			})
		}
	})
	
	// Run tests 12-14 with the new behavior (checking mutations)
	t.Run("NewBehavior", func(t *testing.T) {
		// Ensure the global flag is set to the new behavior
		CheckGlobalDeclarations = false
		
		analyzer := Analyzer()
		
		for i := 12; i <= 14; i++ {
			dir := strconv.Itoa(i)
			t.Run(dir, func(t *testing.T) {
				analysistest.Run(t, testdata, analyzer, dir)
			})
		}
	})
}

func BenchmarkRun(b *testing.B) {
	// Use the new behavior (checking mutations)
	CheckGlobalDeclarations = false
	
	analyzer := Analyzer()
	dir, cleanup, err := analysistest.WriteFiles(map[string]string{
		"file.go": `package code
		import "errors"
		var global = "" 
		var ErrVar = errors.New("myErrVar")
		var myErrVar = errors.New("myErrVar") 
		var errCustom = customError{}
		
		func modifyGlobal() {
			global = "modified" // want "global variable global is being mutated"
		}
		
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
