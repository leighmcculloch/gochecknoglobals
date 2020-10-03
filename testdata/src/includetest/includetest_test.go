package includetest

import "testing"

var Foo = 1 // want "Foo is a global variable"

func TestSome(t *testing.T) {}
