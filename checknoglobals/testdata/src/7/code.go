package code

import (
	"errors"
)

// Those are not errors
var myVar = 1
var myVarFunc = func() int { return 1 }()

// Fake errors
var errFakeErrorUnexported = 1
var ErrFakeErrorExported = 1
var fakeUnexportedErr = 1
var FakeExportedErr = 1
var errFakeErrorUnexportedFunc = func() int { return 1 }()
var ErrFakeErrorExportedFunc = func() int { return 1 }()
var fakeUnexportedErrFunc = func() int { return 1 }()
var FakeExportedErrFunc = func() int { return 1 }()

// Those errors are not named correctly
var myErrVar = errors.New("myErrVar")
var myVarErr = errors.New("myVarErr")
var myVarError = errors.New("myVarErr")
var customErr = customError{"customErr"}
var myErrVarFunc = func() error { return errors.New("myErrVarFunc") }()
var myVarErrFunc = func() error { return errors.New("myVarErrFunc") }()
var myVarErrorFunc = func() error { return errors.New("myVarErrorFunc") }()
var customErrFunc = func() error { return errors.New("customErrFunc") }()

// Those are actual errors which should be ignored
var errUnexported = errors.New("errUnexported")
var ErrExported = errors.New("ErrExported")
var errCustomUnexported = &customError{"errCustomUnexported"}
var ErrCustomExported = &customError{"ErrCustomExported"}
var errCustomUnexported2 = customError2{"errCustomUnexported"}
var ErrCustomExported2 = customError2{"ErrCustomExported"}
var errCustomUnexported3 = customError3("errCustomUnexported")
var ErrCustomExported3 = customError3("ErrCustomExported")
var errUnexportedFunc = func() error { return errors.New("errUnexportedFunc") }()
var ErrExportedFunc = func() error { return errors.New("ErrExportedFunc") }()
var errCustomUnexportedFunc = func() *customError { return &customError{"errCustomUnexportedFunc"} }()
var errCustomUnexportedFuncErr = func() error { return &customError{"errCustomUnexportedFuncErr"} }()
var ErrCustomExportedFunc = func() *customError { return &customError{"ErrCustomExportedFunc"} }()
var ErrCustomExportedFuncErr = func() error { return &customError{"ErrCustomExportedFuncErr"} }()

// Those errors do not really implement the error interface
var errCustomNonPointerUnexported = customError{"errCustomNonPointerUnexported"}
var ErrCustomNonPointerExported = customError{"ErrCustomNonPointerExported"}
var errCustomNonPointerUnexportedFunc = func() customError { return customError{"errCustomNonPointerUnexportedFunc"} }()
var ErrCustomNonPointerExportedFunc = func() customError { return customError{"ErrCustomNonPointerExportedFunc"} }()

// Those actual errors have a declared error type
var declaredErr error = errors.New("declaredErr")
var errDeclared error = errors.New("errDeclared")
var declaredErrFunc error = func() error { return errors.New("declaredErrFunc") }()
var errDeclaredFunc error = func() error { return errors.New("errDeclaredFunc") }()

type customError struct{ e string }

func (e *customError) Error() string { return e.e }

func (e *customError) SomeOtherMethod() string { return e.e }

type customError2 struct{ e string }

func (e customError2) Error() string { return e.e }

func (e customError2) SomeOtherMethod() string { return e.e }

type customError3 string

func (e customError3) Error() string { return string(e) }

func (e customError3) SomeOtherMethod() string { return string(e) }
