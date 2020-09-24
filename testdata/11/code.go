package code

import (
	"errors"
	"fmt"
)

// Bad name, bad method
var allWrong = 1

// Bad name, OK method
var nameWrong = errors.New("bad variable name")
var nameWrongAgain = fmt.Errorf("bad variable name")

// OK name, bad method.
var ErrCorrect = 1
var ErrAlsoCorrect = fmt.Sprintf("foo")

// Ok name, ok method
var ErrThisIsOk = errors.New("some error")
var ErrThisToo = fmt.Errorf("error: %s", "ok")

// Same in group
var (
	allWrongGroup       = 1
	nameWrongGroup      = errors.New("bad variable name")
	nameWrongAgainGroup = fmt.Errorf("bad variable name")
	ErrCorrectGroup     = 1
	ErrAlsoCorrectGroup = fmt.Sprintf("foo")
	ErrThisIsOkGroup    = errors.New("some error")
	ErrThisTooGroup     = fmt.Errorf("error: %s", "ok")
)
