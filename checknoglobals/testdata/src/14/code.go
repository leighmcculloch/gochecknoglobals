package code

// Test direct updating of a global variable
var directUpdateTarget = 10

func directUpdate() {
	directUpdateTarget = 20 // want "global variable directUpdateTarget is being mutated"
}

// Test updating a global through a pointer
var pointerUpdateTarget = 30
var pointerToGlobal = &pointerUpdateTarget

func pointerUpdate() {
	*pointerToGlobal = 40 // want "global variable pointerUpdateTarget is being mutated"
}

// Test updating a global struct field
type Config struct {
	MaxItems  int
	EnableLog bool
}

var globalConfig = Config{
	MaxItems:  100,
	EnableLog: false,
}

func updateGlobalConfig() {
	globalConfig.MaxItems = 200  // want "global variable globalConfig is being mutated"
	globalConfig.EnableLog = true // want "global variable globalConfig is being mutated"
}

// Test local shadows (should not report)
func localShadow() {
	directUpdateTarget := 50   // Local variable, not global
	directUpdateTarget = 60    // Not a mutation of global
	
	pointerUpdateTarget := 70  // Local variable, not global
	pointerUpdateTarget = 80   // Not a mutation of global
}