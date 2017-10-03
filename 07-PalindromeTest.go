// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: Palindrome test
// Source: http://golangcookbook.com/chapters/strings/reverse/
// Objective: Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.
// 1. Convert the string to a mutable array of runes '[]rune'
// 2. Re-cast to a string in the reverse function string(chars)
// 3. Campare two strings, original and reversed, using EqualFold function

package main

import (
	"strings"
	"fmt"
)

func reverse(s string) string {
	chars := []rune(s)

	// multiple initialisation i is 0, j is len(chars)-1;
	// a consolidated bool expression i < j;
	// multiple incrementation i is i+1, j is j-1
	for i, j := 0, len(chars)-1; i < len(chars)/2; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {

	var dummyData string = "Hannnah"

	fmt.Printf("%v\n", reverse(dummyData))
	fmt.Println(strings.EqualFold(reverse(dummyData), dummyData))
}