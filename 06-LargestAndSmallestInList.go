// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: Largest and smallest in list
// Objective: Write a function that returns the largest and smallest elements in a list.

package main

import (
	"container/list"
	"strings"
	"strconv"
	"fmt"
)

func main() {

	var userInput string
	var highNum int64
	var lowNum int64
	var iteration int
	
	// create a new list and put some numbers in it.
	l := list.New()
	iteration = 1

	fmt.Print("Please enter a number you want to put into a list, quit for stop it: ")
	fmt.Scanln(&userInput)

	for strings.EqualFold(userInput, "quit") == false {		
		// parse a string value to an integer number
		pInt, _ := strconv.ParseInt(userInput, 0, 64)
		
		// store a value in a list
		l.PushBack(pInt)
		
		// initialise the first input as highest number and lowest number
		if iteration == 1 {
			highNum = pInt
			lowNum = pInt
		}

		// update new highest number and lowest number
		if pInt > highNum {
			highNum = pInt
		} else if pInt < lowNum {
			lowNum = pInt
		}

		fmt.Print("Would you like to put more number into a list? quit fot stop it: ")
		fmt.Scanln(&userInput)
		
		// update the iteration
		iteration++
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	
	fmt.Printf("Largest number is %d and smallest number is %d\n", highNum, lowNum)

}