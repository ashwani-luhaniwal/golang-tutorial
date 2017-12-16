// Pointer is a variable which stores memory address of another variable

/*******************************
	Declaring pointers
*******************************/
// *T is the type of pointer variable which points to a value of type T

package main

import (
	"fmt"
)

// change the value of pointer
func change(val *int) {
	*val = 55
}

// modify the value in array using pointer
func modify(arr *[3]int) {
	// a[x] is shorthand for (*a)[x]
	(*arr)[0] = 90	// or arr[0] = 90
}

// modifying the value in array using pointer
// recommended way is using the slice
func modifyArray(sls []int) {
	sls[0] = 90
}

func main() {
	b := 225
	// & operator is used to get the address of variable
	var a *int = &b
	fmt.Printf("Type of a in %T\n", a)
	fmt.Println("address of b is", a)	// address may be different since location can by anywhere in memory

	/********************************
		zero value of pointer
	********************************/
	// zero value of pointer is nil
	c := 25
	var d *int
	if d == nil {
		fmt.Println("d is", d)
		d = &c
		fmt.Println("d after initialization is", d)
	}

	/*****************************
		Dereferencing a pointer
	*****************************/
	// Dereferencing a pointer means accessing the value of the variable which the pointer points to.
	// *a is the syntax to dereference a.
	e := 255
	f := &e
	fmt.Println("address of e is", f)
	fmt.Println("value of e is", *f)
	*f++	// change the value of f using pointer
	fmt.Println("new value of e is", e)

	/****************************************
		Passing pointer to a function
	****************************************/
	g := 58
	fmt.Println("value of g before function call is", g)
	h := &g
	change(h)
	fmt.Println("value of g after function call is", g)

	/*********************************************************
		Do not pass a pointer to an array as an argument
		to a function. Use slice instead
	*********************************************************/
	k := [3]int{89, 90, 91}
	modify(&k)
	fmt.Println(k)

	l := [3]int{89, 90, 91}
	modifyArray(l[:])
	fmt.Println(l)

	/**************************************************
		Go doesn't support pointer arithmetic
	**************************************************/
	// Go doesn't support pointer arithmetic like C language
	m := [...]int{109, 110, 111}
	n := &m
	fmt.Println("value of n is", n)
	/*
		n++	// throws an error
	*/
}