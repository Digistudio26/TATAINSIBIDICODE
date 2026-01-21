package main
import "fmt"
func Mean(a, b, c float32) float32 {
      return (( a + b + c ) /3)
}
func main() {
	fmt.Println(Mean(2, 0.25, 9.75))
}
