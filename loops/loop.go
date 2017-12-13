// Here we will see loops in Golang

package main

import "fmt"

func main() {

	// for is the only loop available in Go
	/* Syntax: -
		for initialisation; condition; post {

		}
	*/
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}

	/*******************
		break statement
	*********************/
	// terminate for loop abruptly before its normal execution
	for j := 1; j <= 10; j++ {
		if j > 5 {
			break	// loop is terminated if i > 5
		}
		fmt.Printf(" %d", j)
	}
	fmt.Printf("\nline after for loop")

	/****************************
		continue statement
	****************************/
	// Used to skip the current iteration of the for loop
	for k := 1; k <= 10; k++ {
		if k % 2 == 0 {
			continue
		}
		fmt.Printf("\n %d ", k)
	}

	l := 0
	for ; l <= 10; {	// initialisation and post are omitted
		fmt.Printf("%d ", l)
		l += 2
	}

	m := 0
	for m <= 10 {
		fmt.Printf("%d ", m)
		m += 2
	}

	// Multiple variable declaration in for loop
	for no, n := 10, 1; n <= 10 && no <= 19; n, no = n + 1, no + 1 {
		fmt.Printf("%d * %d = %d\n", no, n, no*n)
	}

	/****************************
		infinite loop
	****************************/
	/* Syntax: -
		for {

		}
	*/
}