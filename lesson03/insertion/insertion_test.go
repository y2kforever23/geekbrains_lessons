package insertion

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	example := []int{1, 2, 3, 4, 5}
	notSorted := []int{5, 4, 3, 2, 1}
	sorted := InsertionSort(notSorted)

	for i, a := range sorted {
		if a != example[i] {
			t.Error("Expected ", example[i], "got ", a)
		}
	}
}
