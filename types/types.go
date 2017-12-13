// In this tutorial, we will see the types which are available in GoLang

package main

import (
	"fmt"
	"unsafe"	// import package in order to access "Sizeof" function
)

func main() {

	/**************
		Bool Type
	**************/
	// Represets a boolean and is either true or false
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	/***********************
		Signed Integers
	***********************/
	// int8 - 8 bit signed integer with size 8 bits
	// int16 - 16 bit signed integer with size 16 bits
	// int32 - 32 bit signed integer with size of 32 bits
	// int64 - 64 bit signed integer with size of 64 bits
	// int - may be 32 or 64 bit depending on underlying platform. (like 32 bits in 32-bit systems and 64 bit in 64 bit systems)
	var p int = 89	// int type
	q := 95	// inferred type based on value
	fmt.Println("Value of p is", p, "and q is", q)

	// %T - used to print variable type in Printf statement
	// %d - used to print the size of variable
	// The "unsafe" package in Go has "Sizeof" function which returns in bytes the size of variable
	// "unsafe" package should be used with care as it might have portability issues
	var r int = 89
	s := 95
	fmt.Println("Value of r is", r, "and s is", s)
	fmt.Printf("Type of r is %T, size of r is %d", r, unsafe.Sizeof(r))
	fmt.Printf("\nType of s is %T, size of s is %d\n", s, unsafe.Sizeof(s))

	/*************************
		Unsigned Integers
	*************************/
	// uint8 - 8 bit unsigned integers with size of 8 bits
	// uint16 - 16 bit unsigned integers with size of 16 bits
	// uint32 - 32 bit unsigned integers with size of 32 bits
	// uint64 - 64 bit unsigned integers with size of 64 bits
	// uint - 32 or 64 bit unsigned integers depending on underlying platform

	/*******************************
		Floating Point Types
	*******************************/
	// float32 - 32 bit floating point numbers
	// float64 - 64 bit floating point numbers
	w, u := 5.67, 8.97	// inferred type of float64 bcoz it's default floating type
	fmt.Println("Type of w %T and u is %T\n", w, u)
	sum := w + u
	diff := w - u
	fmt.Println("Sum is", sum, "and difference is", diff)

	no1, no2 := 56, 89
	fmt.Println("Sum is", no1 + no2, "and difference is", no1 - no2)

	/***********************************
		Complex Types
	***********************************/
	// complex64  - complex numbers with float32 real and imaginary parts
	// complex128 - complex numbers with float64 real and imaginary parts
	// builtin "complex" function is used to construct a complex number with real and imaginary parts
	
	// Syntax - func complex(r, i FloatType) ComplexType
	// Takes real and imaginary part as parameter and returns ComplexType
	// Both real and imaginary parts should be of same type: either float32 or float64
	// If both real and imaginary are float32, then function returns complex value of complex64.
	// If both real and imaginary are float64, then function returns complex value of complex128.
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("Sum is", cadd)
	cmul := c1 * c2
	fmt.Println("Multiplication is", cmul)

	/***************************************
		Other Numeric Types
	***************************************/
	// byte - alias of uint8
	// rune - alias of int32

	/*****************************
		String Type
	*****************************/
	// Strings are collection of bytes in golang
	first := "Ashwani"
	last := "Luhaniwal"
	name := first + " " + last
	fmt.Println("My name is", name)

	/********************************
		Type Conversion
	********************************/
	// Go is very strict about explicit typing
	// There is no automatic type promotion or conversion

	/*
		i := 55	// int
		j := 67.8	// float64
		sum1 := i + j	// int + float64 not allowed
	*/

	// T(v) - syntax to convert a value v to type T
	i := 55		// int
	j := 67.8	// float64
	sum1 := i + int(j)	// j is converted to int
	fmt.Println("Sum of int and float after type conversion is", sum1)

	// Explicit type conversion is required to assign a variable of one type to another
	/*
		k := 10
		var l float64 = float64(k)	// this statement won't work without explicit conversion
		fmt.Println("l", l)
	*/
}