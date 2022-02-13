package code

import (
	"errors"
)

var (
	// Those are not errors
	myVar     = 1                         // want "myVar is a global variable"
	myVarFunc = func() int { return 1 }() // want "myVarFunc is a global variable"

	// Fake errors
	errFakeErrorUnexported     = 1                         // want "errFakeErrorUnexported is a global variable"
	ErrFakeErrorExported       = 1                         // want "ErrFakeErrorExported is a global variable"
	fakeUnexportedErr          = 1                         // want "fakeUnexportedErr is a global variable"
	FakeExportedErr            = 1                         // want "FakeExportedErr is a global variable"
	errFakeErrorUnexportedFunc = func() int { return 1 }() // want "errFakeErrorUnexportedFunc is a global variable"
	ErrFakeErrorExportedFunc   = func() int { return 1 }() // want "ErrFakeErrorExportedFunc is a global variable"
	fakeUnexportedErrFunc      = func() int { return 1 }() // want "fakeUnexportedErrFunc is a global variable"
	FakeExportedErrFunc        = func() int { return 1 }() // want "FakeExportedErrFunc is a global variable"

	// Those errors are not named correctly
	myErrVar       = errors.New("myErrVar")                                 // want "myErrVar is a global variable"
	myVarErr       = errors.New("myVarErr")                                 // want "myVarErr is a global variable"
	myVarError     = errors.New("myVarErr")                                 // want "myVarError is a global variable"
	customErr      = customError{"customErr"}                               // want "customErr is a global variable"
	myErrVarFunc   = func() error { return errors.New("myErrVarFunc") }()   // want "myErrVarFunc is a global variable"
	myVarErrFunc   = func() error { return errors.New("myVarErrFunc") }()   // want "myVarErrFunc is a global variable"
	myVarErrorFunc = func() error { return errors.New("myVarErrorFunc") }() // want "myVarErrorFunc is a global variable"
	customErrFunc  = func() error { return errors.New("customErrFunc") }()  // want "customErrFunc is a global variable"

	// Those are actual errors which should be ignored
	errUnexported              = errors.New("errUnexported")
	ErrExported                = errors.New("ErrExported")
	errCustomUnexported        = &customError{"errCustomUnexported"}
	ErrCustomExported          = &customError{"ErrCustomExported"}
	errUnexportedFunc          = func() error { return errors.New("errUnexportedFunc") }()
	ErrExportedFunc            = func() error { return errors.New("ErrExportedFunc") }()
	errCustomUnexportedFunc    = func() *customError { return &customError{"errCustomUnexportedFunc"} }()
	errCustomUnexportedFuncErr = func() error { return &customError{"errCustomUnexportedFuncErr"} }()
	ErrCustomExportedFunc      = func() *customError { return &customError{"ErrCustomExportedFunc"} }()
	ErrCustomExportedFuncErr   = func() error { return &customError{"ErrCustomExportedFuncErr"} }()

	// Those errors do not really implement the error interface
	errCustomNonPointerUnexported     = customError{"errCustomNonPointerUnexported"}                                     // want "errCustomNonPointerUnexported is a global variable"
	ErrCustomNonPointerExported       = customError{"ErrCustomNonPointerExported"}                                       // want "ErrCustomNonPointerExported is a global variable"
	errCustomNonPointerUnexportedFunc = func() customError { return customError{"errCustomNonPointerUnexportedFunc"} }() // want "errCustomNonPointerUnexportedFunc is a global variable"
	ErrCustomNonPointerExportedFunc   = func() customError { return customError{"ErrCustomNonPointerExportedFunc"} }()   // want "ErrCustomNonPointerExportedFunc is a global variable"

	// Those actual errors have a declared error type
	declaredErr     error = errors.New("declaredErr") // want "declaredErr is a global variable"
	errDeclared     error = errors.New("errDeclared")
	declaredErrFunc error = func() error { return errors.New("declaredErrFunc") }() // want "declaredErrFunc is a global variable"
	errDeclaredFunc error = func() error { return errors.New("errDeclaredFunc") }()
)

type customError struct{ e string }

func (e *customError) Error() string { return e.e }

func (e *customError) SomeOtherMethod() string { return e.e }
