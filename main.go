package main // import "4d63.com/gochecknoglobals"

import (
	"os"

	"4d63.com/gochecknoglobals/checknoglobals"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "./...")
	}

	singlechecker.Main(checknoglobals.Analyzer())
}
