package main

import (
	"reflect"
	"testing"
)

func TestSortedArrayCaseOne(t *testing.T) {
	unsortedArray := []int32{100, 2, 15, 60, 80}
	sortedArray := []int32{2, 15, 60, 80, 100}

	if !reflect.DeepEqual(bubbleSort(unsortedArray), sortedArray) {
		t.Error("Test Failed: Should be [2, 15, 60, 80, 100]")
	}
}

func TestSortedArrayCaseTwo(t *testing.T) {
	unsortedArray := []int32{4, 3, 1, 7, 5}
	sortedArray := []int32{1, 3, 4, 5, 7}

	if !reflect.DeepEqual(bubbleSort(unsortedArray), sortedArray) {
		t.Error("Test Failed: Should be [1, 3, 4, 5, 7]")
	}
}
