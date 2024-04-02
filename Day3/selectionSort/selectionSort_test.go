package main

import (
	"reflect"
	"testing"
)

var (
	randomArray    = []int32{4, 2, 1, 7, 3}
	duplicateArray = []int32{4, 10, 3, 1, 5, 12, 10, 2, 2}
)

func TestSelectionSortRandom(t *testing.T) {
	if !reflect.DeepEqual(selectionSort(randomArray), []int32{1, 2, 3, 4, 7}) {
		t.Error("Test Failed: Should Return {1,2,3,4,7}")
	}
}

func TestSelectionSortDuplicate(t *testing.T) {
	if !reflect.DeepEqual(selectionSort(duplicateArray), []int32{1, 2, 2, 3, 4, 5, 10, 10, 12}) {
		t.Error("Test Failed: Should Return {1,2,2,3,4,5,10,10,12}")
	}
}
