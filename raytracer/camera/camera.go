package camera

import (
	"raytracer/vector"
)

type Camera struct {
	Pos     vector.Vec
	LookAt  vector.Vec
	Forward vector.Vec
	Right   vector.Vec
	Up      vector.Vec
}

func New(pos vector.Vec, lookAt vector.Vec) *Camera {
	down := vector.Vec{0.0, -1.0, 0.0}
	forward := vector.Norm(vector.Minus(lookAt, pos))
	right := vector.Times(1.5, vector.Norm(vector.Cross(forward, down)))
	up := vector.Times(1.5, vector.Norm(vector.Cross(forward, right)))

	return &Camera{pos, lookAt, forward, right, up}
}
