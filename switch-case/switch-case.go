// Conditional statement which evaluates an expression and 
// compares it against list of possible matches.

// Duplicate cases with same constant value are not allowed

package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {
	finger := 4
	switch finger {
		case 1:
			fmt.Println("Thumb")
		case 2:
			fmt.Println("Index")
		case 3:
			fmt.Println("Middle")
		case 4:
			fmt.Println("Ring")
		case 5:
			fmt.Println("Pinky")
	}

	/******************************
		Default case
	******************************/
	switch finger := 8; finger {
		case 1:
			fmt.Println("Thumb")
		case 2:
			fmt.Println("Index")
		case 3:
			fmt.Println("Middle")
		case 4:
			fmt.Println("Ring")
		case 5:
			fmt.Println("Pinky")
		default:	// default case
			fmt.Println("incorrect finger number")
	}

	/************************************
		Multiple expressions in case
	************************************/
	// Include multiple expressions in case by separating them with comma
	letter := "i"
	switch letter {
		case "a", "e", "i", "o", "u":	// multiple expressions in case
			fmt.Println("Vowel")
		default:
			fmt.Println("Not a vowel")
	}

	/******************************
		Expressionless switch
	******************************/
	// Expression in switch is optional and can be omitted
	// If expression is omitted, then switch is considered to be "switch true"
	// and each "case" expression is evaluated for truth
	num := 75
	switch {	// expression is omitted
		case num >= 0 && num <= 50:
			fmt.Println("num is greater then 0 and less than 50")
		case num >= 51 && num <= 100:
			fmt.Println("num is greater then 51 and less than 100")
		case num >= 101:
			fmt.Println("num is greater than 100")
	}

	/***************************
		Fallthrough
	**************************/
	// "fallthrough" statement is used to transfer control to first statement of case
	// which is present immediately after the case which has been executed.
	// It should be last statement in a case. In it presets somewhere in middle
	// then compiler throws an error
	switch num1 := number(); {	// num1 is not constant
		case num1 < 50:
			fmt.Printf("%d is lesser than 50\n", num1)
			fallthrough
		case num1 < 100:
			fmt.Printf("%d is lesser than 100\n", num1)
			fallthrough
		case num1 < 200:
			fmt.Printf("%d is lesser than 200", num1)
	}
}