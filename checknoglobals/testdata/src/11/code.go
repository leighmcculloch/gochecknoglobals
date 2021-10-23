package code

import (
	"embed"
)

//go:embed embedfiles/*
var fileSystem embed.FS

//go:embed embedfiles/embedfile.txt
var str string

//go:embed embedfiles/embedfile.txt
var data []byte

//go:embed embedfiles/embedfile.txt
//
var strExtraCommentLines string

//go:embed embedfiles/embedfile.txt

var strEmptyLines string

var strEmptyLinesNoComment string // want "strEmptyLinesNoComment is a global variable"

// go : embed that does not match
var strEmptyLinesOtherComment string // want "strEmptyLinesOtherComment is a global variable"

//go:embed embedfiles/embedfile.txt
//

var strExtraCommentLinesAndEmptyLines string

var (
	//go:embed embedfiles/embedfile.txt
	groupedStr string

	//go:embed embedfiles/embedfile.txt
	groupedData []byte

	//go:embed embedfiles/embedfile.txt

	groupedStrEmptyLines string

	groupedStrEmptyLinesNoComment string // want "groupedStrEmptyLinesNoComment is a global variable"

	// go : embed that does not match
	groupedStrEmptyLinesOtherComment string // want "groupedStrEmptyLinesOtherComment is a global variable"

	//go:embed embedfiles/embedfile.txt
	//

	groupedStrExtraCommentLinesAndEmptyLines string
)
