package main

func SplitWhitespaces(s string) []string {
	var words []string
	start := -1

	for i, r := range s {
		if r != ' ' && r != '\t' && r != '\n' {
			if start == -1 {
				start = i // mark start of a word
			}
		} else {
			if start != -1 {
				words = append(words, s[start:i]) // slice word
				start = -1
			}
		}
	}

	// Append last word if string doesn't end with whitespace
	if start != -1 {
		words = append(words, s[start:])
	}

	return words
}