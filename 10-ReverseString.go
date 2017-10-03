// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: Reverse string
// Objective: Write a function to reverse a string
// 1. 
// 2. 

package main

import "fmt"

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

	var userInput string

	fmt.Println("Please enter a string to be reversed")
	fmt.Scanf("%s", &userInput)
	fmt.Printf("%v\n", reverse(userInput))

}