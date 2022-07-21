package search

func LinearSearch (arr []float64, element float64) (int) {
	for i := range arr {
		if arr[i] == element {
			return i
		}
	}
	return -1
}