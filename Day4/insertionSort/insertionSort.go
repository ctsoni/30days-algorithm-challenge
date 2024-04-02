package main

import "fmt"

func main() {
	var randomArray []int32 = []int32{4, 2, 7, 1, 3}
	fmt.Println(insertionSort(randomArray))
}

func insertionSort(unorderedArray []int32) []int32 {
	for i := 1; i < len(unorderedArray); i++ {
		temp := unorderedArray[i]
		position := i
		for j := i - 1; j >= 0; j-- {
			if unorderedArray[j] > temp {
				unorderedArray[j+1] = unorderedArray[j]
				position = position - 1
			}
		}

		unorderedArray[position] = temp
	}

	return unorderedArray
}
