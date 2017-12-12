package main

import (
	"fmt"
	"math"
)

func main() {
	/******************************************
		Declaring single variable
	******************************************/
	// Syntax: var name type (where name is the variable name)

	var age int	// variable declaration
	// Note: If variable is not assigned any value, go automatically initializes it with zero value of variable's type
	fmt.Println("My age is", age)

	age = 29	// Assignment
	fmt.Println("My age is", age)

	/*********************************************
		Declaring variable with initial value
	*********************************************/
	// Syntax: var name type = initialValue

	var newAge int = 30	// variable declaration with initial value
	fmt.Println("My age is", newAge)

	/**************************************
		Type Inference
	**************************************/
	// Syntax: var name = initialValue
	// Note: If variable is declared without type, then Go automatically infer the type of that variable

	var typeAge = 35	// type will be inferred
	fmt.Println("My age is", typeAge)

	/**************************************
		Multiple variable declaration
	**************************************/
	// Syntax: var name1, name2 type = initialValue1, initialValue2

	var width, height = 100, 50	// declaring multiple variables using type inference
	fmt.Println("Width is", width, "Height is", height)
	
	var length, bredth int
	fmt.Println("Length is", length, "bredth is", bredth)
	length = 100
	bredth = 50
	fmt.Println("New length is", length, "new bredth is", bredth)

	// Declare variables belong to different types in single statement
	/* 
		var (
			name1 = initialValue1
			name2 = initialValue2
		)
	*/
	var (
		name = "Ashwani"
		myAge = 29
		myHeight int
	)
	fmt.Println("My name is", name, ", age is", myAge, "and height is", myHeight)

	/********************************
		Short hand declaration
	*********************************/
	// Syntax: name := initialValue
	// Short hand declaration requires intial values for all variables in left hand side of assignment

	yourName, yourAge := "Ashwani", 29	// Short hand declaration
	fmt.Println("Your name is", yourName, "and your age is", yourAge)

	// Error test
	// testName, testAge := "Ashwani"	/* Error as testAge has no initial value */

	// Short hand syntax only be used when at least one of the variables in the left hand side of := is newly declared.
	a, b := 20, 30	// declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50	// b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90	// assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)

	// Variables can also be assigned values which are computed during runtime
	x, y := 145.8, 543.8
	z := math.Min(x, y)
	fmt.Println("Minimum value is ", z)

	// Go is strongly typed, variables declared as belonging to one type cannot be assigned a value of another type
	/*
		height := 172	// age is int
		height = "Ashwani"	// error since we are trying to assign a string to variable of type int
	*/
}