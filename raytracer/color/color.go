package color

import (
	"math"
)

type Color struct {
	R, G, B float64
}

var (
	White Color = Color{1.0, 1.0, 1.0}
	Grey  Color = Color{0.5, 0.5, 0.5}
	Black Color = Color{0.0, 0.0, 0.0}
)

var (
	Background   Color = Black
	DefaultColor Color = Black
)

func Scale(k float64, v Color) Color {
	return Color{k * v.R, k * v.G, k * v.B}
}

func Plus(v1 Color, v2 Color) Color {
	return Color{v1.R + v2.R, v1.G + v2.G, v1.B + v2.B}
}

func Times(v1 Color, v2 Color) Color {
	return Color{v1.R * v2.R, v1.G * v2.G, v1.B * v2.B}
}

func ToDrawingColor(c Color) Color {
	legalize := func(d float64) float64 {
		if d > 1 {
			return 1
		} else {
			return d
		}
	}

	return Color{
		math.Floor(legalize(c.R) * 255),
		math.Floor(legalize(c.G) * 255),
		math.Floor(legalize(c.B) * 255),
	}
}
