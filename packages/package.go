// In this, we will see how to use packages in Golang

// Packages are used to organize go source code for better reusability and readability
package main

import (
	"fmt"
	"golang-tutorial/packages/CustomPackages/rectangle"	// importing custom package
	"log"
)

// package variables
var rectLen, rectWidth float64 = 6, 7

// init function to check if length and width are greater then zero
func init() {
	fmt.Println("Main package initialized")
	if rectLen < 0 {
		log.Fatal("Length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("Width is less than zero")
	}
}

/*****************************
	Use of blank identifier 
******************************/
// Illegal in Golang to import package and not to use it anywhere
// This is to avoid bloating of unused packages in order to increase the compilation time
// During active development, you can use blank identifer

func main() {
	fmt.Println("Geometrical shape properties")
	// Area function of rectangle package
	// %.2f format specifier is used to truncate floating point to two decimal places
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	// Diagonal function of rectangle package
	fmt.Printf("Diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}