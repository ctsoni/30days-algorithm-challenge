package main

import "fmt"

func main() {
	var unsortedArray []int32 = []int32{4, 3, 1, 7, 5}

	sortedArray := bubbleSort(unsortedArray)

	fmt.Println(sortedArray)
}

/*
BubbleSort return ordered array from given unordered array.
*/
func bubbleSort(unsortedArray []int32) []int32 {
	sortedArray := make([]int32, len(unsortedArray))

	_ = copy(sortedArray, unsortedArray)

	// iterate for all element
	for j := len(sortedArray) - 1; j >= 0; j-- {
		// take pointer to first element
		// compare it to the next element
		for i := 0; i < len(sortedArray)-1; i++ {
			// if the first element greater than next element, then swap it
			if sortedArray[i] > sortedArray[i+1] {
				sortedArray[i], sortedArray[i+1] = sortedArray[i+1], sortedArray[i]
			}
			// if the first element less than next element, then do nothing
		}
	}

	return sortedArray
}
