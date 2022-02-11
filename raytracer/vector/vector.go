package vector

import "math"

type Vec struct {
	X, Y, Z float64
}

func Times(k float64, v Vec) Vec {
	return Vec{k * v.X, k * v.Y, k * v.Z}
}

func Minus(v1 Vec, v2 Vec) Vec {
	return Vec{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func Plus(v1 Vec, v2 Vec) Vec {
	return Vec{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func Dot(v1 Vec, v2 Vec) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Cross(v1 Vec, v2 Vec) Vec {
	return Vec{v1.Y*v2.Z - v1.Z*v2.Y, v1.Z*v2.X - v1.X*v2.Z, v1.X*v2.Y - v1.Y*v2.X}
}

func Mag(v Vec) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func Norm(v Vec) Vec {
	mag := Mag(v)
	div := math.Inf(1)
	if mag > 0 {
		div = 1.0 / mag
	}
	return Times(div, v)
}
