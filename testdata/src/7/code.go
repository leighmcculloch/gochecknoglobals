package code

import (
	"errors"
)

// Those are not errors
var myVar = 1 // want "myVar is a global variable"

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

type customError struct{ e string }

func (e *customError) Error() string { return e.e }
