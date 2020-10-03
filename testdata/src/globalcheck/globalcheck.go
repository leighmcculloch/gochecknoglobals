package globalcheck

import (
	"errors"
	"net/http"
	"regexp"
)

const constant = 0

var myVar = 0 // want "myVar is a global variable"

var theVar = true // want "theVar is a global variable"

var myVar1, myVar2 = 1, 2 // want "myVar1 is a global variable" "myVar2 is a global variable"

var _ = 0

// Those are fake errors which are currently not detected
// because they start with 'err' or 'Err' and we don't
// check if such a variable implements the error interface.
var errFakeErrorUnexported = 1
var ErrFakeErrorExported = 1

// Those errors are not named correctly
var myErrVar = errors.New("myErrVar")    // want "myErrVar is a global variable"
var myVarErr = errors.New("myVarErr")    // want "myVarErr is a global variable"
var myVarError = errors.New("myVarErr")  // want "myVarError is a global variable"
var customErr = customError{"customErr"} // want "customErr is a global variable"

// Those are actual errors which should be ignored
var errUnexported = errors.New("errUnexported")
var ErrExported = errors.New("ErrExported")
var errCustomUnexported = customError{"errCustomUnexported"}
var ErrCustomExported = customError{"ErrCustomExported"}

// Those actual errors have a declared error type
var declaredErr error = errors.New("declaredErr") // want "declaredErr is a global variable"
var errDeclared error = errors.New("errDeclared")

// Check errors in grouped var block
var (
	myVarGroup = 1 // want "myVarGroup is a global variable"

	// Those are fake errors which are currently not detected
	// because they start with 'err' or 'Err' and we don't
	// check if such a variable implements the error interface.
	errFakeErrorUnexportedGroup = 1
	ErrFakeErrorExportedGroup   = 1

	// Those errors are not named correctly
	myErrVarGroup   = errors.New("myErrVar")   // want "myErrVarGroup is a global variable"
	myVarErrGroup   = errors.New("myVarErr")   // want "myVarErrGroup is a global variable"
	myVarErrorGroup = errors.New("myVarErr")   // want "myVarErrorGroup is a global variable"
	customErrGroup  = customError{"customErr"} // want "customErrGroup is a global variable"

	// Those are actual errors which should be ignored
	errUnexportedGroup       = errors.New("errUnexported")
	ErrExportedGroup         = errors.New("ErrExported")
	errCustomUnexportedGroup = customError{"errCustomUnexported"}
	ErrCustomExportedGroup   = customError{"ErrCustomExported"}

	// Those actual errors have a declared error type
	declaredErrGroup error = errors.New("declaredErr") // want "declaredErrGroup is a global variable"
	errDeclaredGroup error = errors.New("errDeclared")
)

var Version string   // want "Version is a global variable"
var version22 string // want "version22 is a global variable"
var version string

type customError struct{ e string }

func (e *customError) Error() string { return e.e }

// IsOnlyDigitsRe is a global regexp that should be OK.
var IsOnlyDigitsRe = regexp.MustCompile(`^\d+$`)

// Testing multiple variable assignments, all allowed.
var (
	PrecompileOne   = regexp.MustCompile(`[a-z]{1,3}`)
	PrecompileTwo   = regexp.MustCompile(`[a-z]{3,6}`)
	PrecompileThree = regexp.MustCompile(`[a-z]{6,9}`)
)

// Testing multiple variable assignments, some unallowed.
var (
	PrecompileFour = regexp.MustCompile(`[a-z]{1,3}`)
	PrecompileFive = regexp.MustCompile(`[a-z]{3,6}`)
	PrecompileSix  = regexp.MustCompile(`[a-z]{6,9}`)
	HTTPClient     = http.Client{} // want "HTTPClient is a global variable"
)

var invalid = 1 //nolint // Can ignore with nolint directive
