package scene

import (
	"raytracer/camera"
	"raytracer/light"
	"raytracer/thing"
)

type Scene struct {
	Things []thing.Thing
	Lights []light.Light
	Camera *camera.Camera
}
