package sort

import "testing"

func TestInsertion (t *testing.T) {
	arr := []float64 {
		1, 2, 0.01,
	}
	InsertionSort(arr, len(arr))
	for i, v := range arr {
		if (i == 0 && v != 0.01) {
			t.Errorf("element 0 is not 0.01")
		}
		if (i == 1 && v != 1) {
			t.Errorf("element 1 is not 1")
		}
		if (i == 2 && v != 2) {
			t.Errorf("element 2 is not 2")
		}
	}
}

func BenchmarkInsertion(b *testing.B) {
	arr := []float64 {
		1, 2, 0.01,
	}
	InsertionSort(arr, len(arr))
}

