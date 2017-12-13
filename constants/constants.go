// In this we will see constants in Golang

package main

import (
	"fmt"
	"math"
)

func main() {

	/*********************
		Constants
	*********************/
	// Constants are used to denote fixed values
	// Cannot be re-assigned again to any other value
	const a = 55	// allowed
	/*
		a = 89			// re-assigned not allowed, throw error
	*/

	// Constant value should be known at compile time. 
	// Hence, cannot be assigned to value returned by function call
	// as function call takes place at runtime.
	var b = math.Sqrt(4)	// allowed
	fmt.Println("Value of b", b)
	/*
		const c = math.Sqrt(4)	// not allowed, throw error
	*/

	/**************************
		String Constants
	**************************/
	// The type of string constants belong to "untyped" or constant strings doesn't have any type.
	const hello = "Hello World"
	// constant hello has no type

	// Creating typed constant, like as follows: -
	const typedhello string = "Hello World"

	// Go is strictly typed language.
	// Mixing types during assignment is not allowed
	var defaultName = "Ashwani"	// allowed
	fmt.Println("Value of defaultName is", defaultName)
	type myString string
	var customName myString = "Again, Ashwani"	// allowed
	fmt.Println("Value of customName is", customName)
	/*
		customName = defaultName	// not allowed
	*/

	/******************************
		Boolean Constants
	******************************/
	// There are two untyped constants "true" and "false"
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst	// allowed
	fmt.Println("Value of defaultBool is", defaultBool)
	var customBool myBool = trueConst	// allowed
	fmt.Println("customBool:", customBool)
	/*
		defaultBool = customBool	// not allowed
	*/

	/*******************************
		Numeric Constants
	*******************************/
	// Includes integers, floats and complex constants
	const x = 5
	var intVar int = x
	var int32Var int32 = x
	var float64Var float64 = x
	var complex64Var complex64 = x
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	/*******************************
		Numeric Expressions
	*******************************/
	// Numeric constants are free to be mixed and matched in expressions
	var y = 5.9/8
	fmt.Printf("y's type %T value %v\n", y, y)

}