package main

import (
	"fmt"
	"math"
)

func main() {
	var arr []int32 = []int32{1, 13, 69, 88, 90}
	var resultFound int32 = binarySearch(arr, 13)
	var resultNotFound int32 = binarySearch(arr, 50)

	fmt.Println(resultFound)
	fmt.Println(resultNotFound)
}

/*
BinarySearch return the index of searchElement on an ordered array.
If the searchValue not exists on array it will return -1.
*/
func binarySearch(array []int32, searchValue int32) int32 {
	// takes the length of array and take the middle element
	// keep track the biggest value and lowest value
	var lastElement = len(array) - 1
	var firstElement = 0
	// compare it with searchValue
	for i := 0; i < len(array)/2; i++ {
		var middleElement = math.Round((float64(firstElement) + float64(lastElement)/2))

		if searchValue == array[int32(middleElement)] {
			return int32(middleElement)
		}

		// if the searchValue < middle element, then biggest value become the middle element
		if searchValue < array[int32(middleElement)] {
			lastElement = int(middleElement)
		}

		// if the searchValue > middle element, then the lowest value become the middle element
		if searchValue > array[int32(middleElement)] {
			firstElement = int(middleElement)
		}
	}

	return -1
}
