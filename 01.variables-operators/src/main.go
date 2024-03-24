// PART 6: VARIABLE SCOPE
// **********************
// To run this file, rename it to main.go and run as the listed call at the end of the file

// All the identifiers in Go live in a scope
// Go is a block-based scoped: { ... }
// The parent-child scopes relationship is defined when the code compiles, not when the code runs
// - package - Top-level scope

// Scope Resolution
// - When accessing a variable, Go looks at the scope the code was defined in
// - If found, stop and use the variable
// - If not found, look in the parent scope
// - If not found, then the parent's parent scope
// - Do the same until reaching the package scope
// - If still not found, raise an error
// This searching is all done based on a variable name
// If a variable with the name is found but is of the wrong type, Go raises an error
// NOTE: A function cannot access a variable defined in a child scope

// Package
// *******
package main

// Imports
// *******
import "fmt"

// Global Variables and Constants
// ******************************

var level = "pkg"

// main
// ****

// The main entry point of the application.
func main() {
    level = "main"
	fmt.Println("Main start  :", level)

	if true {
        level = "block"
		fmt.Println("Block start :", level)
		funcScope()
	}
}

// Functions
// *********

// Functions creating another scope.
func funcScope() {
	fmt.Println("funcScope start :", level)
}

// FOR WINDOWS:
//  To run:                 go run 01.variables-operators\src\main.go
//  To compile:             go build -o 01.variables-operators\bin\variables-operators.exe 01.variables-operators\src\main.go
//  To run after compile:   .\01.variables-operators\bin\variables-operators.exe
//  Compile + Run:          go build -o 01.variables-operators\bin\variables-operators.exe 01.variables-operators\src\main.go && .\01.variables-operators\bin\variables-operators.exe

// FOR LINUX:
//  To run:                 go run 01.variables-operators/src/main.go
//  To compile:             go build -o 01.variables-operators/bin/variables-operators 01.variables-operators/src/main.go
//  To run after compile:   ./01.variables-operators/bin/variables-operators
//  Compile + Run:          go build -o 01.variables-operators/bin/variables-operators 01.variables-operators/src/main.go && ./01.variables-operators/bin/variables-operators
