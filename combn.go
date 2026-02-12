package main




func CombN(n int) []string {
	var res []string

	if n < 1 || n > 10 {
		return res
	}

	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = i
	}

	for {
		b := make([]byte, n)
		for i, v := range indexes {
			b[i] = byte('0' + v)
		}
		res = append(res, string(b))

		i := n - 1
		for i >= 0 && indexes[i] == 10-n+i {
			i--
		}
		if i < 0 {
			break
		}

		indexes[i]++
		for j := i + 1; j < n; j++ {
			indexes[j] = indexes[j-1] + 1
		}
	}

	return res
}