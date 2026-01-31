package main


func SubstrIndex(s string, toFind string) int {
	// Edge case: empty toFind
	if len(toFind) == 0 {
		return 0
	}

	// If toFind is longer than s, it can't be found
	if len(toFind) > len(s) {
		return -1
	}

	// Loop through s and try to match toFind
	for i := 0; i <= len(s)-len(toFind); i++ {
		match := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}