// A collection of elements which belongs to same type
// Mixing values of different types in an array is not allowed in Golang

package main

import (
	"fmt"
)

// passing arrays by value
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("Inside function ", num)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries) - 2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)	// copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	/**********************
		Declaration
	**********************/
	// n[T] - n denotes number of elements and T represents type of each element
	var a [3]int	// int array with length 3
	a[0] = 12	// array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	// Array declaration using short hand
	b := [3]int{12, 78, 50}	// short hand declaration to create array
	fmt.Println(b)

	// Not necessary to assign value to all indexes
	c := [3]int{12}
	fmt.Println(c)

	// It's possible to ignore the length of arrau in declaration and
	// replace with ... and compiler will understand it
	d := [...]int{12, 78, 50}	// ... makes compiler determine the length of array
	fmt.Println(d)

	// The size of array is part of type and they cannot be resized.
	// "slices" can be used to resize the array
	/*
		f := [3]int{5, 78, 8}
		var g [5]int
		g = f	// not possible since [3]int and [5]int are distinct types
	*/

	/******************************
		Arrays are value types
	******************************/
	// Arrays in Go are value types and not reference types
	// When they are assigned to new variable, a copy of original array is 
	// assigned to the new variable. If changes are made to new variable, it
	// won't reflect in original array.
	f := [...]string{"USA", "China", "India", "Germany", "France"}
	g := f	// a copy of f is assigned to g
	g[0] = "Singapore"
	fmt.Println("f is ", f)
	fmt.Println("g is ", g)

	// When arrays are passed to functions as parameters, they are 
	// passed by value and the original array in unchanged
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("Before passing to function ", num)
	changeLocal(num)	// num is passed by value
	fmt.Println("After passing to function ", num)

	/****************************
		Length of array
	****************************/
	// Length of array is found by passing the array as parameter to "len" function
	h := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("Length of h is", len(h))

	/**************************************
		Iterating arrays using range
	**************************************/
	// "for" loop can be used to iterate over elements of array
	m := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(m); i++ {	// looping from 0 to length of array
		fmt.Printf("%d th element of m is %.2f\n", i, m[i])
	}

	// You can use "range" form of "for" loop
	// "range" returns both index and value at that index.
	p := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range p {	// range returns both index and value
		fmt.Printf("%d the element of p is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nSum of all elements of p", sum)

	// In case you only want value not index, then you can used _ blank identifier
	/* Syntax: -
		for _, v := range p {	// ignores index

		}
	*/

	/*****************************
		Multidimensional Arrays
	******************************/
	q := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},	// this comma is necessary bcoz compiler will throw as error
	}
	printarray(q)
	var r [3][2]string
	r[0][0] = "apple"
	r[0][1] = "samsung"
	r[1][0] = "microsoft"
	r[1][1] = "google"
	r[2][0] = "AT&T"
	r[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(r)

	/****************
		Slices
	****************/
	// slice is convenient, flexible and powerful wrapper on top of an array
	// Slices do not own any data on their own.
	// They are just references to existing arrays.
	
	/****************************
		Creating a slice
	****************************/
	// A slice with elements of type T is represented by []T
	s := [5]int{76, 77, 78, 79, 80}
	var t []int = s[1:4]	// creates a slice from a[1] to a[3]
	fmt.Println(t)

	// Another way to create slice
	u := []int{6, 7, 8}	// creates an array and returns a slice reference
	fmt.Println(u)

	/********************************
		Modifying a slice
	********************************/
	// Any modification done to slice will be reflected in underlying array
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	// When number of slices share the same underlying array, the changes
	// that each one makes will be reflected in the array
	numa := [3]int{78, 79, 80}
	nums1 := numa[:]	// creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	/************************************
		Length and capacity of a slice
	*************************************/
	// length of slice in the number of elements in slice
	// capacity of slice is the number of elements in the underlying array starting
	// from the index from which the slice is created.
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))	// length of is 2 and capacity is 6

	// A slice can be re-sliced upto its capacity. Anything beyond that will cause runtime error
	fruitarray1 := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice1 := fruitarray1[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice1), cap(fruitslice1))
	fruitslice1 = fruitslice1[:cap(fruitslice1)]	// re-slicing fruitslice1 till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice1), "and capacity is", cap(fruitslice1))

	/***************************************
		Creating a slice using make
	***************************************/
	// func make([]T, len, cap) - used to create a slice by passing the type, length and capacity.
	// The capacity parameter is optional and defaults to length
	// The make function creates an array and returns slice reference to it.
	i := make([]int, 5, 5)
	fmt.Println(i)

	/************************************
		Appending to slice
	************************************/
	// func append(s []T, x ...T) []T
	// x ...T means function accepts variable number of arguments for parameter x. These type of functions are called variadic functions
	// Arrays are of fixed length and their length cannot be increased
	// Slices are dynamic and new elements can be appended to slice using append function
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))	// capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))	// capacity of cars is doubled to 6


	// The zero value of slice type is nil.
	// A nil slice has length and capacity 0.
	// It is possible to append values to nil slice using append function
	var names []string	// zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	// It is also possible to append one slice to another using ... operator
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	/*********************************
		Passing a slice to a function
	**********************************/
	// Slices can be thought of as being represented internally by structure type
	/*
		type slice struct {
			Length			int
			Capacity		int
			ZerothElement	*byte
		}
	*/
	// When slice is passed to function, even though its passed by value, the pointer
	// variable will refer to the same underlying array.
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)	// function modifies the slice
	fmt.Println("slice after function call", nos)	// modifications are visible outside

	/************************************
		Multidimensional slices
	************************************/
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	}

	/********************************
		Memory Optimisation
	********************************/
	// If we have large array and we are interested in processing only a small part of it.
	// Henceforth we create a slice from that array and start processing the slice.
	// But the array will still be in memory since the slice references it.
	// One way to solve this problem is use "copy" function

	/* Syntax: -
		func copy(dst, src []T) int {	// to make a copy of that slice

		}
	*/
	// This way we can use new slice and the original array can be garbage collected.
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}