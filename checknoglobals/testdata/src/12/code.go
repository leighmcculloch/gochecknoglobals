package code

//go:embed embedfiles/embedfile.txt
var str string // want "str is a global variable"

//go:embed embedfiles/embedfile.txt
var data []byte // want "data is a global variable"

//go:embed embedfiles/embedfile.txt
//
var strExtraCommentLines string // want "strExtraCommentLines is a global variable"

//go:embed embedfiles/embedfile.txt

var strEmptyLines string // want "strEmptyLines is a global variable"

var strEmptyLinesNoComment string // want "strEmptyLinesNoComment is a global variable"

// go : embed that does not match
var strEmptyLinesOtherComment string // want "strEmptyLinesOtherComment is a global variable"

//go:embed embedfiles/embedfile.txt
//

var strExtraCommentLinesAndEmptyLines string // want "strExtraCommentLinesAndEmptyLines is a global variable"

var (
	//go:embed embedfiles/embedfile.txt
	groupedStr string // want "groupedStr is a global variable"

	//go:embed embedfiles/embedfile.txt
	groupedData []byte // want "groupedData is a global variable"

	//go:embed embedfiles/embedfile.txt

	groupedStrEmptyLines string // want "groupedStrEmptyLines is a global variable"

	groupedStrEmptyLinesNoComment string // want "groupedStrEmptyLinesNoComment is a global variable"

	// go : embed that does not match
	groupedStrEmptyLinesOtherComment string // want "groupedStrEmptyLinesOtherComment is a global variable"

	//go:embed embedfiles/embedfile.txt
	//

	groupedStrExtraCommentLinesAndEmptyLines string // want "groupedStrExtraCommentLinesAndEmptyLines is a global variable"
)
