package sort

import "errors"

func findMinimumOfArray (arr []float64, startingIndex int, n int) (int, error) {
	if len(arr) < 2 {
		return 0, errors.New("arr length should be grater than 2")
	}
	var minValue float64 = arr[startingIndex]
	var minIndex int = startingIndex
	for i := startingIndex; i < n; i++ {
		if arr[i] < minValue {
			minValue = arr[i]
			minIndex = i
		}
	}
	return minIndex, nil
}

func swap(arr []float64, firstIndex int, secondIndex int) error {
	if len(arr) < 2 {
		return errors.New("arr length should be grater than 2")
	}
	if firstIndex == secondIndex {
		return errors.New("firstIndex and secondIndex should be different")
	}
	temp := arr[firstIndex]
	arr[firstIndex] = arr[secondIndex]
	arr[secondIndex] = temp
	return nil
}

func SelectionSort(arr []float64, n int) error {
	if len(arr) < 2 {
		return errors.New("arr length should be grater than 2")
	}
	for i:=0; i < n; i++ {
		index, err := findMinimumOfArray(arr, i, n)
		if err != nil {
			return err
		}
		swap(arr, i, index)
	}
	return nil
}