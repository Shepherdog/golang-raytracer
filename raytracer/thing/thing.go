package thing

import (
	"raytracer/surface"
	"raytracer/vector"
)

type Intersection struct {
	Thing Thing
	Ray   vector.Ray
	Dist  float64
}

type Thing interface {
	Intersect(ray vector.Ray) Intersection
	Normal(pos vector.Vec) vector.Vec
	GetSurface() surface.Surface
}
