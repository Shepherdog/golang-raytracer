package light

import (
	"raytracer/color"
	"raytracer/vector"
)

type Light struct {
	Pos   vector.Vec
	Color color.Color
}
