package ch6

import "image/color"

type Point struct {
	x, y	float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}
