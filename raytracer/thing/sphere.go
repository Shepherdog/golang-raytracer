package thing

import (
	"math"
	"raytracer/surface"
	"raytracer/vector"
)

type Sphere struct {
	Radius  float64
	Center  vector.Vec
	Surface surface.Surface
}

func (s *Sphere) Normal(pos vector.Vec) vector.Vec {
	return vector.Norm(vector.Minus(pos, s.Center))
}

func (s *Sphere) Intersect(ray vector.Ray) Intersection {
	eo := vector.Minus(s.Center, ray.Start)
	v := vector.Dot(eo, ray.Dir)
	dist := .0
	if v >= 0 {
		disc := s.Radius*s.Radius - (vector.Dot(eo, eo) - v*v)
		if disc >= 0 {
			dist = v - math.Sqrt(disc)
		}
	}
	if dist == 0 {
		return Intersection{}
	} else {
		return Intersection{Thing: s, Ray: ray, Dist: dist}
	}
}

func (s *Sphere) GetSurface() surface.Surface {
	return s.Surface
}
