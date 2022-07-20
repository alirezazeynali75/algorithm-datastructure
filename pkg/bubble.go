package pkg

import "errors"


func BubbleSort (arr []float64, n int) error {
	if len(arr) == 0 {
		return errors.New("array length should be grater than 0")
	}
	flag := false
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nil
}