// You can edit this code!
// Click here and start typing.
package main

func main() {
	println("feb", "september")
}
func Season(month string) string {
	switch month {
	case "jan", "feb", "dec":
		return "winter"
	case "mar", "apr", "may":
		return "spring"
	case "jun", "jul", "aug":
		return "summer"
	case "sep", "oct", "nov":
		return "autumn"
	default:
		return "invalid input: " + month
	}
}
