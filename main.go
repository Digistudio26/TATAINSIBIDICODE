package main

import "fmt"

func main() {
    s := "Hello, World!"
    last := GetLastRune(s)
    fmt.Printf("The last rune of \"%s\" is \"%c\"\n", s, last)
}
