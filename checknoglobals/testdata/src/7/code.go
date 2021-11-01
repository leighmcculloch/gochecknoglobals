package code

import (
	"errors"
)

// Those are not errors
var myVar = 1 // want "myVar is a global variable"

// Those are fake errors which are not allowed since they are basic literals and
// can be converted o const.
var errFakeErrorUnexported = 1 // want "errFakeErrorUnexported is a global variable, should be a const"
var ErrFakeErrorExported = 1   // want "ErrFakeErrorExported is a global variable, should be a const"

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

type customError struct{ e string }

func (e *customError) Error() string { return e.e }
