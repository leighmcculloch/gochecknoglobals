package code

import (
	"errors"
	"net/http"
	"regexp"
)

// myVar is just a bad named global var.
var myVar = 1 // want "myVar is a global variable"

// ErrNotFound is an error and should be OK.
var ErrNotFound = errors.New("this is error")

// ErrIsNotErr is an error and should be OK.
var ErrIsNotErr = 1

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
