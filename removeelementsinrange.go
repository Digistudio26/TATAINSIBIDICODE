package maingit 


func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	// Ensure correct order
	if from > to {
		from, to = to, from
	}

	
	if from < 0 {
		from = 0
	}
	if to > len(arr) {
		to = len(arr)
	}

	// If nothing to remove, return original array
	if from >= to {
		return arr
	}

	// Remove elements in range [from, to)
	return append(arr[:from], arr[to:]...)
}