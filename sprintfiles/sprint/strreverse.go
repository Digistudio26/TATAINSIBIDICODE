package main

func StrReverse(s string) string {
	runes := []rune(s)        // Convert string to runes to handle UTF-8
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Swap runes
	}
	return string(runes)       // Convert runes back to string
}