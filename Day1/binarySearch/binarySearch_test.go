package main

import "testing"

var (
	arr             []int32 = []int32{1, 13, 69, 88, 90}
	elementFound    int32   = 13
	elementNotFound int32   = 100
)

func TestElementFound(t *testing.T) {
	if binarySearch(arr, elementFound) != 1 {
		t.Error("Test Failed: Seharusnya 1")
	}
}

func TestElementNotFound(t *testing.T) {
	if binarySearch(arr, elementNotFound) != -1 {
		t.Error("Test Failed: Seharusnya -1")
	}
}
