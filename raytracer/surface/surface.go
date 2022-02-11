package surface

import (
	"raytracer/color"
	"raytracer/vector"
)

type Surface interface {
	Diffuse(pos vector.Vec) color.Color
	Specular(pos vector.Vec) color.Color
	Reflect(pos vector.Vec) float64
	GetRoughness() float64
}
