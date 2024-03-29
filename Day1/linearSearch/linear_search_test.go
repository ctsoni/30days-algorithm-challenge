package main

import "testing"

var (
	arr            []int32 = []int32{1, 27, 69, 100, 101}
	searchFound    int32   = 69
	searchNotFound int32   = 89
)

func TestElementFound(t *testing.T) {
	if linearSearch(arr, searchFound) != 2 {
		t.Error("Test Failed: Should be 2")
	}
}

func TestElementNotFound(t *testing.T) {
	if linearSearch(arr, searchNotFound) != -1 {
		t.Error("Test Failed: Should be -1")
	}
}
