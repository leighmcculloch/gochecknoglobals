package code

// This is a global variable declaration which should be allowed
var myGlobal = 5

// This is another global variable declaration which should be allowed
var anotherGlobal = "hello"

func modifyGlobal() {
	myGlobal = 10 // want "global variable myGlobal is being mutated"
}

func incrementGlobal() {
	myGlobal++ // want "global variable myGlobal is being mutated"
}

func modifyGlobalWithCompoundAssignment() {
	myGlobal += 20 // want "global variable myGlobal is being mutated"
}

func modifyAnotherGlobal() {
	anotherGlobal = "world" // want "global variable anotherGlobal is being mutated"
}

// This local variable with the same name as a global should not be reported
func localVariable() {
	myGlobal := 15 // This is a local variable, not the global one
	myGlobal = 20  // This should not be reported
}