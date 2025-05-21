package code

import (
	"errors"
)

var (
	// Those are not errors
	myVar     = 1
	myVarFunc = func() int { return 1 }()

	// Fake errors
	errFakeErrorUnexported     = 1
	ErrFakeErrorExported       = 1
	fakeUnexportedErr          = 1
	FakeExportedErr            = 1
	errFakeErrorUnexportedFunc = func() int { return 1 }()
	ErrFakeErrorExportedFunc   = func() int { return 1 }()
	fakeUnexportedErrFunc      = func() int { return 1 }()
	FakeExportedErrFunc        = func() int { return 1 }()

	// Those errors are not named correctly
	myErrVar       = errors.New("myErrVar")
	myVarErr       = errors.New("myVarErr")
	myVarError     = errors.New("myVarErr")
	customErr      = customError{"customErr"}
	myErrVarFunc   = func() error { return errors.New("myErrVarFunc") }
	myVarErrFunc   = func() error { return errors.New("myVarErrFunc") }
	myVarErrorFunc = func() error { return errors.New("myVarErrorFunc") }
	customErrFunc  = func() error { return errors.New("customErrFunc") }

	// Those are actual errors which should be ignored
	errUnexported              = errors.New("errUnexported")
	ErrExported                = errors.New("ErrExported")
	errCustomUnexported        = &customError{"errCustomUnexported"}
	ErrCustomExported          = &customError{"ErrCustomExported"}
	errUnexportedFunc          = func() error { return errors.New("errUnexportedFunc") }
	ErrExportedFunc            = func() error { return errors.New("ErrExportedFunc") }
	errCustomUnexportedFunc    = func() *customError { return &customError{"errCustomUnexportedFunc"} }
	errCustomUnexportedFuncErr = func() error { return &customError{"errCustomUnexportedFuncErr"} }
	ErrCustomExportedFunc      = func() *customError { return &customError{"ErrCustomExportedFunc"} }
	ErrCustomExportedFuncErr   = func() error { return &customError{"ErrCustomExportedFuncErr"} }

	// Those errors do not really implement the error interface
	errCustomNonPointerUnexported     = customError{"errCustomNonPointerUnexported"}
	ErrCustomNonPointerExported       = customError{"ErrCustomNonPointerExported"}
	errCustomNonPointerUnexportedFunc = func() customError { return customError{"errCustomNonPointerUnexportedFunc"} }
	ErrCustomNonPointerExportedFunc   = func() customError { return customError{"ErrCustomNonPointerExportedFunc"} }

	// Those actual errors have a declared error type
	declaredErr     error = errors.New("declaredErr")
	errDeclared     error = errors.New("errDeclared")
	declaredErrFunc error = func() error { return errors.New("declaredErrFunc") }
	errDeclaredFunc error = func() error { return errors.New("errDeclaredFunc") }
)

type customError struct{ e string }

func (e *customError) Error() string { return e.e }

func (e *customError) SomeOtherMethod() string { return e.e }
