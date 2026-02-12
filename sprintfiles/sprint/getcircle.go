package main

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	const pi float32 = 3.14

	return Circle{
		Radius:    r,
		Diameter: 2 * r,
		Area:      pi * r * r,
		Perimeter: 2 * pi * r,
	}
}