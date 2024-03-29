package main

import "fmt"

func main() {
	var array []int32 = []int32{1, 3, 5, 1, 7}

	fmt.Println(hasDuplicate(array))
	fmt.Println(hasDuplicateOptimized(array))
}

/*
hasDuplicate return true if an array has a duplicate value.
The time complexity is N^2.
*/
func hasDuplicate(array []int32) bool {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i] == array[j] && i != j {
				return true
			}
		}
	}
	return false
}

/*
hasDuplicateOptimized return true if an array has a duplicate value.
The time complexity is N.
Optimized using map data structure to track element exist in array.
*/
func hasDuplicateOptimized(array []int32) bool {
	var existingElement = make(map[int32]int32)
	for i := 0; i < len(array); i++ {
		if existingElement[array[i]] == 1 {
			return true
		} else {
			existingElement[array[i]] = 1
		}
	}

	return false
}
