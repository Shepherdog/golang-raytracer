package raytracer

import (
	"math"
	"raytracer/scene"
	"raytracer/surface"
	"raytracer/thing"
	"raytracer/vector"
	"testing"
)

func TestRayTracer_testRay(t *testing.T) {
	s := scene.Scene{Things: []thing.Thing{
		&thing.Sphere{Radius: 3.0, Center: vector.Vec{5.0, 0.0, 0.0}, Surface: surface.Shiny{Roughness: 250}},
	}}
	type args struct {
		ray   vector.Ray
		scene scene.Scene
	}
	tests := []struct {
		name string
		r    *RayTracer
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"tangent point", &RayTracer{}, args{ray: vector.Ray{vector.Vec{0.0, 0.0, 0.0}, vector.Vec{0.8, 0.6, 0.0}}, scene: s}, 4.0},
		{"no intersection", &RayTracer{}, args{ray: vector.Ray{vector.Vec{0.0, 0.0, 0.0}, vector.Vec{0.79, 0.6, 0.0}}, scene: s}, math.Inf(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.testRay(tt.args.ray, tt.args.scene); got != tt.want {
				t.Errorf("RayTracer.testRay() = %v, want %v", got, tt.want)
			}
		})
	}
}
