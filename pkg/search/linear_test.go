package search

import "testing"

func TestLinearWhichFindCorrectValue(t *testing.T) {
	arr := []float64 {
		0.01, 1, 3, 20, 4.21, 14.54532,
	}
	expectation := 4
	got := LinearSearch(arr, 4.21)
	if expectation != got {
		t.Errorf("Output %q not equal to expected %q", got, expectation)
	}
}

func TestLinearWhichDidNotFindValue(t *testing.T) {
	arr := []float64 {
		0.01, 1, 3, 20, 4.21, 14.54532,
	}
	expectation := -1
	got := LinearSearch(arr, 0)
	if expectation != got {
		t.Errorf("Output %q not equal to expected %q", got, expectation)
	}
}

func BenchmarkLinear(b *testing.B) {
	arr := []float64 {
		0.01, 1, 3, 20, 4.21, 14.54532,
	}
	for _,v := range arr {
		LinearSearch(arr, v)
	}
}