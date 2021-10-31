package code

import (
	"errors"
)

var (
	// Those are not errors
	myVar = 1 // want "myVar is a global variable"

	// Those are fake errors which are not allowed since they are basic literals and
	// can be converted o const.
	errFakeErrorUnexported = 1 // want "errFakeErrorUnexported is a global variable, should be a const"
	ErrFakeErrorExported   = 1 // want "ErrFakeErrorExported is a global variable, should be a const"

	// Those errors are not named correctly
	myErrVar   = errors.New("myErrVar")   // want "myErrVar is a global variable"
	myVarErr   = errors.New("myVarErr")   // want "myVarErr is a global variable"
	myVarError = errors.New("myVarErr")   // want "myVarError is a global variable"
	customErr  = customError{"customErr"} // want "customErr is a global variable"

	// Those are actual errors which should be ignored
	errUnexported       = errors.New("errUnexported")
	ErrExported         = errors.New("ErrExported")
	errCustomUnexported = customError{"errCustomUnexported"}
	ErrCustomExported   = customError{"ErrCustomExported"}

	// Those actual errors have a declared error type
	declaredErr error = errors.New("declaredErr") // want "declaredErr is a global variable"
	errDeclared error = errors.New("errDeclared")
)

type customError struct{ e string }

func (e *customError) Error() string { return e.e }
