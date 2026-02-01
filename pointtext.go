package main



type Point struct {
	X    float64
	Y    float64
	Text string
}

func PointText(p Point) Point {
	// Convert floats to string with 6 decimal places
	xStr := strconv.FormatFloat(p.X, 'f', 6, 64)
	yStr := strconv.FormatFloat(p.Y, 'f', 6, 64)

	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: "Text at (" + xStr + ", " + yStr + ")",
	}
}