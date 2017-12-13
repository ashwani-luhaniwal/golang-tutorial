
/* Syntax: -
	if condition {

	} else if condition {

	} else {
		
	}
*/

package main

import (
	"fmt"
)

func main() {
	num := 10
	if num % 2 == 0 {
		// checks if number is even
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// optional statement component in if which is executed before the condition is evaluated
	/* 
		if statement; condition {

		}
	*/
	if num1 := 10; num1 % 2 == 0 {
		fmt.Println(num1, "is even")
	} else {
		fmt.Println(num1, "is odd")
	}

	num3 := 99
	if num3 <= 50 {
		fmt.Println("Number is less than or equal to 50")
	} else if num3 >= 51 && num3 <= 100 {
		fmt.Println("Number is between 51 and 100")
	} else {
		// else statement should start in the same line else compiler raise an error
		// bcoz Go inserts semicolons automatically
		fmt.Println("Number is greater than 100")
	}
}