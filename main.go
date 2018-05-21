package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagPrintHelp := flag.Bool("help", false, "")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gochecknoglobals [path] [path] ...\n")
	}
	flag.Parse()

	if *flagPrintHelp {
		flag.Usage()
		return
	}

	paths := flag.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}

	for _, path := range paths {
		messages, err := checkNoGlobals(path)
		for _, message := range messages {
			fmt.Fprintf(os.Stdout, "%s\n", message)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
		}
	}
}
