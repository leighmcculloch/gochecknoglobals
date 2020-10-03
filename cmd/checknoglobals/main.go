package main

import (
	"4d63.com/gochecknoglobals"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gochecknoglobals.Analyzer)
}
