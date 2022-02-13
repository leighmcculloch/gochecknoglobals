package code

import (
	"errors"
)

var (
	// Those are not errors
	myVar = 1 // want "myVar is a global variable"

	// Fake errors
	errFakeErrorUnexported = 1 // want "errFakeErrorUnexported is a global variable"
	ErrFakeErrorExported   = 1 // want "ErrFakeErrorExported is a global variable"
	fakeUnexportedErr      = 1 // want "fakeUnexportedErr is a global variable"
	FakeExportedErr        = 1 // want "FakeExportedErr is a global variable"

	// Those errors are not named correctly
	myErrVar   = errors.New("myErrVar")   // want "myErrVar is a global variable"
	myVarErr   = errors.New("myVarErr")   // want "myVarErr is a global variable"
	myVarError = errors.New("myVarErr")   // want "myVarError is a global variable"
	customErr  = customError{"customErr"} // want "customErr is a global variable"

	// Those are actual errors which should be ignored
	errUnexported       = errors.New("errUnexported")
	ErrExported         = errors.New("ErrExported")
	errCustomUnexported = &customError{"errCustomUnexported"}
	ErrCustomExported   = &customError{"ErrCustomExported"}

	// Those errors do not really implement the error interface
	errCustomNonPointerUnexported = customError{"errCustomNonPointerUnexported"} // want "errCustomNonPointerUnexported is a global variable"
	ErrCustomNonPointerExported   = customError{"ErrCustomNonPointerExported"}   // want "ErrCustomNonPointerExported is a global variable"

	// Those actual errors have a declared error type
	declaredErr error = errors.New("declaredErr") // want "declaredErr is a global variable"
	errDeclared error = errors.New("errDeclared")
)

type customError struct{ e string }

func (e *customError) Error() string { return e.e }

func (e *customError) SomeOtherMethod() string { return e.e }
