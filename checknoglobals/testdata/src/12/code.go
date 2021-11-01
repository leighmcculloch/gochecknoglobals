package code

import "errors"

// myVar is just a bad named global var.
var myVar = 1 // want "myVar is a global variable, should be a const"

// ErrNotFound is an error and should be OK.
var ErrNotFound = errors.New("this is error")

// ErrIsNotErr is an error and should be OK.
var ErrIsNotErr = 1 // want "ErrIsNotErr is a global variable, should be a const"
