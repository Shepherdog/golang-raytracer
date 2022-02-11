package surface

import (
	"raytracer/color"
	"raytracer/vector"
)

type Shiny struct {
	Roughness float64
}

func (Shiny) Diffuse(pos vector.Vec) color.Color {
	return color.White
}

func (Shiny) Specular(pos vector.Vec) color.Color {
	return color.Grey
}

func (Shiny) Reflect(pos vector.Vec) float64 {
	return 0.7
}

func (s Shiny) GetRoughness() float64 {
	return s.Roughness
}
