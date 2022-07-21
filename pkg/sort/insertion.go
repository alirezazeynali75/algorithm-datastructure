package sort

import "errors"

func InsertionSort (arr []float64, n int) error {
	if (len(arr) < 2) {
		return errors.New("array length should be grater than 2")
	}
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && arr[j - 1] > arr[j] {
			key := arr[j]
			arr[j] = arr[j - 1];
 			arr[j - 1] = key;
			j--
		}
	}
	return nil
}