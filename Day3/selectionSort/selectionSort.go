package main

import "fmt"

func main() {
	array := []int32{4, 2, 7, 1, 3}
	fmt.Println(selectionSort(array))
}

/*
selectionSort return an ordered array from undordered one.
The time complexity of this algorithm is N^2.
*/
func selectionSort(unorderedArray []int32) []int32 {
	for i := 0; i < len(unorderedArray)-1; i++ {
		lowestIndex := i
		for j := i + 1; j < len(unorderedArray); j++ {
			if unorderedArray[j] < unorderedArray[lowestIndex] {
				lowestIndex = j
			}
		}

		if lowestIndex != i {
			unorderedArray[i], unorderedArray[lowestIndex] = unorderedArray[lowestIndex], unorderedArray[i]
		}

	}

	return unorderedArray
}
