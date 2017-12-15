// Map is builtin type in Go which associates a value to key
// The value can be retrieved using corresponding key.

/********************************
	Create a map
********************************/
// Map can be created by passing type of key and value to make function
/* Syntax: -
	make(map[type of key]type of value)

	personSalary := make(map[string]int)
	// personSalary has "string" keys and "int" values
*/

// zero value of map is nil. So, map has to initialize using make function

package main

import (
	"fmt"
)

func main() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	/***********************************
		Adding items to a map
	***********************************/
	// Syntax for adding new items to a map is same as that of arrays.
	salary := make(map[string]int)
	salary["steve"] = 12000
	salary["jamie"] = 15000
	salary["mike"] = 9000
	fmt.Println("salary map contents:", salary)

	// It's also possible to initialize a map during declaration
	perSalary := map[string]int {
		"steve": 12000,
		"jamie": 15000,
	}
	perSalary["mike"] = 9000
	fmt.Println("perSalary map contents:", perSalary)

	// It's not necessary only string types should be keys.
	// All comparable types such as boolean, integer, float, complex, string can also be keys.

	/***********************************
		Acessing items of map
	***********************************/
	// map[key] - used to retrieve elements of map
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", perSalary[employee])

	// If element is not present, then map will return zero value of the type of that element.
	employee1 := "jamie"
	fmt.Println("Salary of", employee1, "is", perSalary[employee])
	fmt.Println("Salary of joe is", perSalary["joe"])

	// We want to know whether a key is present in a map or not
	/* Syntax: -
		value, ok := map[key]
		// If ok is true, then key is present and its value is present in the variable value
	*/
	newEmp := "joe"
	value, ok := perSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	// "range" form of for loop is used to iterate over all elements of map
	fmt.Println("All items of a map")
	for key, value := range perSalary {
		fmt.Printf("perSalary[%s] = %d\n", key, value)
	}

	// Important :- the order of retrieval of values from a map when using "for range" is not
	// guaranteed to be same for each execution program

	/****************************
		Deleting items from map
	****************************/
	// delete(map, key) - used to delete key from a map.
	// the delete function doesn't return any value
	fmt.Println("map before deletion", perSalary)
	delete(perSalary, "steve")
	fmt.Println("map after deletion", perSalary)

	/*******************************
		Length of map
	*******************************/
	// Length can be determined using "len" function
	fmt.Println("length is", len(perSalary))

	/******************************
		Maps are reference types
	*******************************/
	// when map is assigned to new variable, they both point to same internal data structure.
	// So, changes made in one will reflect in other
	fmt.Println("Original person salary", perSalary)
	newPersonSalary := perSalary
	newPersonSalary["ashwani"] = 18000
	fmt.Println("Person salary changed", perSalary)
	
	// when maps are passed as parameters to functions
	// When any change is made to map inside function, it will be visible to caller.

	/*******************************
		Maps equality
	*******************************/
	// Maps can't be compared using == operator.
	// The == can be used to check if a map is nil
	map1 := map[string]int {
		"one": 1,
		"two": 2,
	}
	map2 := map1
	fmt.Println(map2)
	/*
		if map1 == map2 {

		}
	*/
}