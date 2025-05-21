package mutations

import (
	"errors"
	"regexp"
)

// 1. Simple Mutation
var myGlobal int

func f1() {
	myGlobal = 10 // want "mutation of global variable 'myGlobal'"
}

// 2. No Mutation (Read Only)
var myGlobalRO int

func f2() int {
	return myGlobalRO // Read, no mutation
}

// 3. Mutation of different types (direct assignment)
type MyS struct{ Val int }

var globalStruct MyS
var globalStructPtr *MyS

func f3() {
	globalStruct = MyS{Val: 1}    // want "mutation of global variable 'globalStruct'"
	globalStructPtr = &MyS{Val: 2} // want "mutation of global variable 'globalStructPtr'"

	// Note: Field assignments are not direct mutations of the global variable identifier.
	// globalStruct.Val = 100 // This is not a mutation of 'globalStruct' itself.
}

// 4. Multiple globals, some mutated, some not
var g1 int
var g2 int
var g3 int // Unused/not mutated

func f4() {
	g1 = 1 // want "mutation of global variable 'g1'"
	_ = g2 // Read g2, no mutation
}

// 5. Assignments of various kinds
var gAssignOps int
var gIncDec int // Note: ++/-- are IncDecStmts, not AssignStmts, and might not be caught by current linter

func f5() {
	gAssignOps += 5 // want "mutation of global variable 'gAssignOps'"
	gAssignOps = gAssignOps + 1 // want "mutation of global variable 'gAssignOps'"

	// gIncDec++ // This is an IncDecStmt. If it's not caught, the test will reflect that.
	// Based on the current linter implementation (checking AssignStmt),
	// IncDecStmt won't be flagged. So, no "want" here for gIncDec++.
}

// 6. Test variables that were previously on the isAllowed list for declaration
var normallyAllowedRegexp = regexp.MustCompile(".*")
var normallyAllowedError = errors.New("test err")
var versionString = "1.0.0" // Example of another previously 'allowed' declaration

func f6() {
	normallyAllowedRegexp = regexp.MustCompile("^abc$") // want "mutation of global variable 'normallyAllowedRegexp'"
	normallyAllowedError = errors.New("new err")      // want "mutation of global variable 'normallyAllowedError'"
	versionString = "1.0.1"                           // want "mutation of global variable 'versionString'"
}

// Test for global variable used as a receiver (should not be flagged as mutation)
type MyType int
var myTypeGlobal MyType
func (m MyType) MutateMethod() {
	// This is a method call, not a direct assignment to myTypeGlobal
	// No "want" expected here
}
func (m *MyType) MutatePointerMethod() {
	// This is a method call, not a direct assignment to myTypeGlobal
	// No "want" expected here
}
func callReceiverMethods() {
	myTypeGlobal.MutateMethod()
	(&myTypeGlobal).MutatePointerMethod()
}

// Test for mutations within init functions
var initGlobalVar = 10
func init() {
	initGlobalVar = 20 // want "mutation of global variable 'initGlobalVar'"
}

// Test for mutations of struct fields (indirect mutation, should not be flagged by current linter)
type Point struct { X, Y int }
var pGlobal = Point{X:1, Y:2}
func mutatePointField() {
    pGlobal.X = 100 // This is not a direct assignment to pGlobal
}

// Test for mutations of map elements (indirect mutation, should not be flagged by current linter)
var mapGlobal = make(map[string]string)
func mutateMapElement() {
	mapGlobal["key"] = "value" // This is not a direct assignment to mapGlobal
}

// Test for mutations of slice elements (indirect mutation, should not be flagged by current linter)
var sliceGlobal = []int{1,2,3}
func mutateSliceElement() {
	if len(sliceGlobal) > 0 {
		sliceGlobal[0] = 100 // This is not a direct assignment to sliceGlobal
	}
}

// Test for append to slice (reassignment of the global)
func appendToSlice() {
	sliceGlobal = append(sliceGlobal, 4) // want "mutation of global variable 'sliceGlobal'"
}
