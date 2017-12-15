// Variadic function which accepts variable number of arguments
// If the last parameter of function is denoted by ...T, 
// then function can accept any number of arguments of type T for last parameter
// Only last parameter of function is allowed to be variadic.

/* Example: -
	func append(slice []Type, elems ...Type) []Type
	// elems is the variadic parameter
*/

package main

import (
	"fmt"
)

// variadic function
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func findNew(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
}

func changeAnother(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	// variadic function work is by converting variable number of arguments passed to new slice of type variadic parameter
	// find function expects variadic int argument, so compiler convert it to slice of type int []int{89, 90, 95}
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	// nums is nil slice with length and capacity 0
	find(87)

	/*******************************************
		Passing a slice to variadic function
	*******************************************/
	nums := []int{89, 90, 95}
	/*
		// According to definition, num ...int means it will accept variable numer of arguments of type int
		// The variadic arguments will be converted into slice of type int
		// here nums is already a int slice and new slice []int is tried to created using nums which results in error
		
		findNew(89, nums)	// throws compiler error
	*/

	// There is syntactix sugar which can be used to pass a slice to variadic function.
	// You have to suffix the slice with ... If that is done, the slice is directly passed to function
	// without a new slice being created.
	findNew(89, nums...)

	// Other examples of passing slice in variadic function
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)

	welcomenew := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcomenew)
}