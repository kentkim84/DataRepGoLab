// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: Merge list and sort
// Source: http://www.geeksforgeeks.org/merge-sort/
// Objective: Write a function that merges two sorted lists into a new sorted list, e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].
// 1. Recursively divided in two halves till the size becomes 1
// 2. Once the size becomes 1, the merge processes comes into action and starts merging arrays back till the complete array is merged.

package main

import "fmt"

// Merges two subarrays of items.
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]
func merge(items []int, l int, m int, r int) []int {
    
    var i int
    var j int
    var k int
    
    n1 := m - l + 1
    n2 :=  r - m
    
    // create temp arrays
    var L = make([]int, n1)
    var R = make([]int, n2)

	// create new array
	var newItems = make([]int, len(L) + len(R))
 
    // copy data to temp arrays L[] and R[]
    for i = 0; i < n1; i++ {
		L[i] = items[l + i]
	}
    for j = 0; j < n2; j++ {
		R[j] = items[m + 1+ j]
	}
 
    // Merge the temp arrays back into arr[l..r]
    i = 0 // Initial index of first subarray
    j = 0 // Initial index of second subarray
    k = l // Initial index of merged subarray
    for (i < n1 && j < n2) {
        if (L[i] <= R[j]) {
            items[k] = L[i]
            i++
        } else {
            items[k] = R[j]
            j++
        }
        k++
    }
 
    // copy the remaining elements of L[], if there are any
    for (i < n1) {
        items[k] = L[i]
        i++
        k++
    }
 
    // copy the remaining elements of R[], if there are any
    for (j < n2) {
        items[k] = R[j]
        j++
        k++
	}
	return newItems
}
 
// l is for left index and r is right index of the
// sub-array of items to be sorted
func mergeSort(items []int, l int, r int) []int{
    var newItems = make([]int, len(items))

	// if there is nothing to be halved, return itself
	if (len(items) == 1) {
		return items
    }
    
	if (l < r) {
        // get the middle of the array
        m := l+(r-l)/2
 
        // sort first and second halves
        mergeSort(items, l, m)
        mergeSort(items, m+1, r)
 
        // merge left and right items
        newItems = merge(items, l, m, r)
    }
    return newItems
}

func main() {

    items := []int{12, 11, 13, 5, 6, 7}
    items_size := len(items)
 
    fmt.Println("Given array is", items)
 
    mergeSort(items, 0, items_size-1)
 
    fmt.Println("Sorted array is", items)
}