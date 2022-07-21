package search

import "testing"

func TestBinaryWhichFindCorrectValue(t *testing.T) {
	arr := []float64 {
		0.01, 1, 3, 4.21, 14.54532, 20,
	}
	expectation := 3
	got, _ := BinarySearch(arr, 4.21)
	if expectation != got {
		t.Errorf("Output %q not equal to expected %q", got, expectation)
	}
}

func TestBinaryWhichDidNotFindValue(t *testing.T) {
	arr := []float64 {
		0.01, 1, 3, 4.21, 14.54532, 20,
	}
	expectation := -1
	got, _ := BinarySearch(arr, 0)
	if expectation != got {
		t.Errorf("Output %q not equal to expected %q", got, expectation)
	}
}

func BenchmarkBinary(b *testing.B) {
	arr := []float64 {
		0.01, 1, 3, 4.21, 14.54532, 20,
	}
	for _,v := range arr {
		BinarySearch(arr, v)
	}
}