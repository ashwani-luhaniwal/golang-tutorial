// string in Go is a slice of bytes.
// Can be created by enclosing their content inside " "
// strings in Go are unicode compliant and are UTF-8 encoded

package main

import (
	"fmt"
	"unicode/utf8"
)

// print Unicode UTF-8 encoded values of "Hello World"
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		// %x is format specifier for hexadecimal
		fmt.Printf("%x ", s[i])	
	}
}

// print character of string
func printChars(s string) {
	// string is converted to a slice of runes
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// %c format specifier used to print character of string
		fmt.Printf("%c ", runes[i])
	}
}

// print characters and bytes
func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

// length fron rune
func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

/*
func mutate(s string) string {
	// We are trying to change the string, but it's possible as strings are immutable in Golang
	s[0] = 'a'	// any valid unicode character within single quote is a rune
	return s
}
*/

// To make string mutable, first string needs to be converted to a slice of runes.
// Then that slice is mutable with whatever changes needed and converted back to new string
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	name := "Hello World"
	fmt.Println(name)

	/*************************************
		Accessing individual bytes of string
	*******************************************/
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)

	/************************
		rune
	************************/
	// It's the alias of int32
	// rune represents Unicode code point in Go.
	// It doesn't matter how many bytes the code point occupies, it can be represent by run.
	fmt.Printf("\n\n")
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)

	/*********************************************
		for range loop on a string
	********************************************/
	fmt.Printf("\n")
	printCharsAndBytes(name)

	/***************************************
		Constructing string from slice of bytes
	********************************************/
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)

	// decimal equivalent of hex values
	byteSlice1 := []byte{67, 97, 102, 195, 169}	// decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str1 := string(byteSlice1)
	fmt.Println(str1)

	/*********************************************
		Constructing a string from slice of runes
	*********************************************/
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str2 := string(runeSlice)
	fmt.Println(str2)

	/****************************
		Length of the string
	****************************/
	// func RuneCountInString(s string) (n int) - function of utf8 package is used to find length of string
	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)

	/********************************
		Strings are immutable
	********************************/
	h := "hello"
	/*
		fmt.Println(mutate(h)
	*/
	// h is converted to a slice of runes and passed to mutate function
	fmt.Println(mutate([]rune(h)))
}