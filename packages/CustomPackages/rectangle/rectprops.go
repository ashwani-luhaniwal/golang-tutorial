
package rectangle

import (
	"math"
	"fmt"
)

/****************************
	Exported Names
****************************/
// Any variable or function which starts with capital letter are exported names in Golang

/**********************
	Init function
**********************/
// Should not have any return type and should not have any parameters
// It cannot be called explicitly in our source code
/* Syntax: -
	func init() {

	}
*/
// Used to perform initialisation tasks
// Used to verify the correctness of program before execution starts

// Package level variables are initialized first
// init function is called next.
// Package can have multiple init functions and called in sequential order
// If package imports other packages, then imported packages are initialised first.
// Package is initialised only once even if it is imported from multiple packages.

// init function added
func init() {
	fmt.Println("Rectangle package initialized")
}

// calculate area of rectangle
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// calculate diagonal of a rectangle
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}