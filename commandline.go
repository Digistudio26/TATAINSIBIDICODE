package main


import (
	"bufio"
    "fmt"
    "os"
	 "strings"
)
func main ()  {
		scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Select operation (1/2):\n 1. Encrypt.\n 2. Decrypt")
	

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	for input != "1" || input != "2"{
		scanner.Scan()
		fmt.Println("Select operation (1/2):\n 1. Encrypt.\n 2. Decrypt.")
	}
	input = strings.TrimSpace(scanner.Text())
	fmt.Printf("%s\n", input)
	
	
	

} 