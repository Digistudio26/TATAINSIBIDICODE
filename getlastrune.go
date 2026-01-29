package main



func GetLastRune (s string) rune {

	new_array := []rune{s}
	return new_array[len(new_array) - 1]
}