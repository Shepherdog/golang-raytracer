package surface

import (
	"math"
	"raytracer/color"
	"raytracer/vector"
)

type Checkerboard struct {
	Roughness float64
}

func (Checkerboard) Diffuse(pos vector.Vec) color.Color {
	if (int64)(math.Floor(pos.Z)+math.Floor(pos.X))%2 != 0 {
		return color.White
	} else {
		return color.Black
	}
}

func (Checkerboard) Specular(pos vector.Vec) color.Color {
	return color.White
}

func (Checkerboard) Reflect(pos vector.Vec) float64 {
	if (int64)(math.Floor(pos.Z)+math.Floor(pos.X))%2 != 0 {
		return 0.1
	} else {
		return 0.7
	}
}

func (c Checkerboard) GetRoughness() float64 {
	return c.Roughness
}
