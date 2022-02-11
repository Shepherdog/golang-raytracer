package thing

import (
	"raytracer/surface"
	"raytracer/vector"
)

type Plane struct {
	Offset  float64
	Norm    vector.Vec
	Surface surface.Surface
}

func (p *Plane) Normal(pos vector.Vec) vector.Vec {
	return p.Norm
}

func (p *Plane) Intersect(ray vector.Ray) Intersection {
	denom := vector.Dot(p.Norm, ray.Dir)
	if denom > 0 {
		return Intersection{}
	} else {
		dist := (vector.Dot(p.Norm, ray.Start) + p.Offset) / -denom
		return Intersection{Thing: p, Ray: ray, Dist: dist}
	}
}

func (p *Plane) GetSurface() surface.Surface {
	return p.Surface
}
