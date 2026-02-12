package main


func GenerateRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min
	result := make([]int, size)

	for i := range result { // use range
		result[i] = min + i
	}

	return result
}
