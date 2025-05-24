package code

// Test struct type for nested fields
type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street  string
	City    string
	Country Country
}

type Country struct {
	Name string
	Code string
}

// Global variables that will have their addresses taken or fields modified
var globalPerson = Person{
	Name: "John",
	Age:  30,
	Address: Address{
		Street: "123 Main St",
		City:   "Anytown",
		Country: Country{
			Name: "USA",
			Code: "US",
		},
	},
}

var globalValue = 42

// Test taking the address of a global and modifying it
func modifyViaPointer() {
	ptr := &globalValue // Taking address is ok
	*ptr = 100          // want "global variable globalValue is being mutated"
}

// Test modifying a field of a global struct
func modifyStructField() {
	globalPerson.Name = "Jane" // want "global variable globalPerson is being mutated"
	globalPerson.Age = 31      // want "global variable globalPerson is being mutated"
}

// Test taking the address of a struct field and modifying it
func modifyViaFieldPointer() {
	namePtr := &globalPerson.Name // Taking address is ok
	*namePtr = "Bob"              // want "global variable globalPerson is being mutated"
}

// Test modifying a nested field
func modifyNestedField() {
	globalPerson.Address.Street = "456 Oak Ave"       // want "global variable globalPerson is being mutated"
	globalPerson.Address.City = "Othertown"           // want "global variable globalPerson is being mutated"
	globalPerson.Address.Country.Name = "Canada"      // want "global variable globalPerson is being mutated"
	globalPerson.Address.Country.Code = "CA"          // want "global variable globalPerson is being mutated"
}

// Test taking the address of a deeply nested field and modifying it
func modifyViaNestedFieldPointer() {
	countryNamePtr := &globalPerson.Address.Country.Name // Taking address is ok
	*countryNamePtr = "Mexico"                          // want "global variable globalPerson is being mutated"
}

// Local variable shadowing shouldn't be reported
func localShadowing() {
	// Create a local variable with the same name
	globalPerson := Person{Name: "Local"}
	
	// These should not be reported as they modify the local variable
	globalPerson.Name = "Local Modified"
	globalPerson.Age = 25
}