package main


 func StrLength(s string) []int {
	lenrune := len([]rune(s))
	lenBytes := len([] byte(s))
	return [] int {lenrune, lenBytes}
 }