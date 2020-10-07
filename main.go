package main // import "4d63.com/gochecknoglobals"

import (
	"4d63.com/gochecknoglobals/checknoglobals"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(checknoglobals.Analyzer())
}
