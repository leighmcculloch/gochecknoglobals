package code

import (
	"errors"
)

// Those are not errors
var myVar = 1                             // want "myVar is a global variable"
var myVarFunc = func() int { return 1 }() // want "myVarFunc is a global variable"

// Fake errors
var errFakeErrorUnexported = 1                             // want "errFakeErrorUnexported is a global variable"
var ErrFakeErrorExported = 1                               // want "ErrFakeErrorExported is a global variable"
var fakeUnexportedErr = 1                                  // want "fakeUnexportedErr is a global variable"
var FakeExportedErr = 1                                    // want "FakeExportedErr is a global variable"
var errFakeErrorUnexportedFunc = func() int { return 1 }() // want "errFakeErrorUnexportedFunc is a global variable"
var ErrFakeErrorExportedFunc = func() int { return 1 }()   // want "ErrFakeErrorExportedFunc is a global variable"
var fakeUnexportedErrFunc = func() int { return 1 }()      // want "fakeUnexportedErrFunc is a global variable"
var FakeExportedErrFunc = func() int { return 1 }()        // want "FakeExportedErrFunc is a global variable"

// Those errors are not named correctly
var myErrVar = errors.New("myErrVar")                                       // want "myErrVar is a global variable"
var myVarErr = errors.New("myVarErr")                                       // want "myVarErr is a global variable"
var myVarError = errors.New("myVarErr")                                     // want "myVarError is a global variable"
var customErr = customError{"customErr"}                                    // want "customErr is a global variable"
var myErrVarFunc = func() error { return errors.New("myErrVarFunc") }()     // want "myErrVarFunc is a global variable"
var myVarErrFunc = func() error { return errors.New("myVarErrFunc") }()     // want "myVarErrFunc is a global variable"
var myVarErrorFunc = func() error { return errors.New("myVarErrorFunc") }() // want "myVarErrorFunc is a global variable"
var customErrFunc = func() error { return errors.New("customErrFunc") }()   // want "customErrFunc is a global variable"

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
var errCustomNonPointerUnexported = customError{"errCustomNonPointerUnexported"}                                         // want "errCustomNonPointerUnexported is a global variable"
var ErrCustomNonPointerExported = customError{"ErrCustomNonPointerExported"}                                             // want "ErrCustomNonPointerExported is a global variable"
var errCustomNonPointerUnexportedFunc = func() customError { return customError{"errCustomNonPointerUnexportedFunc"} }() // want "errCustomNonPointerUnexportedFunc is a global variable"
var ErrCustomNonPointerExportedFunc = func() customError { return customError{"ErrCustomNonPointerExportedFunc"} }()     // want "ErrCustomNonPointerExportedFunc is a global variable"

// Those actual errors have a declared error type
var declaredErr error = errors.New("declaredErr") // want "declaredErr is a global variable"
var errDeclared error = errors.New("errDeclared")
var declaredErrFunc error = func() error { return errors.New("declaredErrFunc") }() // want "declaredErrFunc is a global variable"
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
