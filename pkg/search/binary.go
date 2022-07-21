package search

func BinarySearch(arr []float64, value float64) (int, int) {
	max := len(arr) - 1
	min := 0
	var guess int
	step := 0
	for max >= min {
		guess = (max + min) / 2
		step++
		if arr[guess] == value {
			return guess, step
		} else if arr[guess] > value {
			max = guess - 1
		} else {
			min = guess + 1
		}
	}
	return -1, 0
}