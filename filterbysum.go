package main


func FilterBySum(arr [][]int, limit int) [][]int {
	result := [][]int{}

	for _, sub := range arr {
		sum := 0
		for _, val := range sub {
			sum += val
		}

		if sum >= limit {
			result = append(result, sub)
		}
	}

	return result
}