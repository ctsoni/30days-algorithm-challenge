// Based on Common Sense Guide to Data Structures and Algorithm Book

package main

import "fmt"

func main() {
	var arr []int32 = []int32{1, 27, 69, 100, 101}

	var resultFound int32 = linearSearch(arr, 69)
	var resultNotFound int32 = linearSearch(arr, 89)

	fmt.Println(resultFound)
	fmt.Println(resultNotFound)
}

/*
LinearSearch return the index of searchValue on an ordered array.
If the searchValue not exists it will return -1.
*/
func linearSearch(array []int32, searchValue int32) int32 {
	for i := 0; i <= len(array); i++ {
		if array[i] == searchValue {
			return int32(i)
		}

		// in ordered array we can break if the searchValue already < element
		if searchValue < array[i] {
			break
		}
	}

	return -1
}
