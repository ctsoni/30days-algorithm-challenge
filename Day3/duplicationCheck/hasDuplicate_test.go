package main

import (
	"testing"
)

var (
	arrayContainDuplicate    []int32 = []int32{1, 5, 3, 4, 1}
	arrayNotContainDuplicate []int32 = []int32{5, 6, 69, 81, 100}
)

func TestHasDuplicateTrue(t *testing.T) {
	if hasDuplicate(arrayContainDuplicate) != true {
		t.Error("Test Failed: Should be true!")
	}
}

func TestHasDuplicateFalse(t *testing.T) {
	if hasDuplicate(arrayNotContainDuplicate) != false {
		t.Error("Test Failed: Should be false!")
	}
}

func TestHasDuplicateOptimiedTrue(t *testing.T) {
	if hasDuplicateOptimized(arrayContainDuplicate) != true {
		t.Error("Test Failed: Should be true!")
	}
}

func TestHasDuplicateoptimizedFalse(t *testing.T) {
	if hasDuplicateOptimized(arrayNotContainDuplicate) != false {
		t.Error("Test Failed: Should be false!")
	}
}
