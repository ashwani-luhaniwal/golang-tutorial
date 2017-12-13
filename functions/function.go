// In this, we will discuss how to use functions in Golang

package main

import "fmt"

// A block of code which performs specific tasks
/**************************
	Function Declaration
***************************/
/* Syntax: -
	func functionname(parametername type) returntype {
		// function body
	}
*/
func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// If consecutive parameters are of same type, then we can avoid writing type eachtime
func calculateBillNew(price, no int) int {
	var totalPriceNew = price * no
	return totalPriceNew
}

/**********************************
	Multiple Return Values
**********************************/
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

/*********************************
	Named Return Values
********************************/
// Return named values from a function
func areaPerimeter(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return 	// no explicit return value
}

/**********************************
	Blank Identifier
**********************************/
// _ is known as blank identifier in Golang
// Used in place of any value of any type
func blankIdentifier(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {

	price, no := 90, 6
	// Syntax to call a function as follows: -
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
	
	totalPriceNew := calculateBillNew(price, no)
	fmt.Println("New total price is", totalPriceNew)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	newArea, newPerimeter := areaPerimeter(10.8, 5.6)
	fmt.Printf("New area %f and New perimeter %f\n", newArea, newPerimeter)

	// Skip specific return value using blank identifier
	calArea, _ := blankIdentifier(10.8, 5.6)	// perimeter is discarded
	fmt.Printf("Area %f\n", calArea)
}