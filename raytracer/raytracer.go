package raytracer

import (
	"image"
	"math"
	"raytracer/camera"
	"raytracer/color"
	"raytracer/light"
	"raytracer/scene"
	"raytracer/thing"
	"raytracer/util"
	"raytracer/vector"
)

const maxDepth uint8 = 5

type RayTracer struct{}

func (*RayTracer) intersections(ray vector.Ray, scene scene.Scene) (closestInter thing.Intersection) {
	closest := math.Inf(1)
	for _, thing := range scene.Things {
		if inter := thing.Intersect(ray); inter.Thing != nil && inter.Dist < closest {
			closestInter = inter
			closest = inter.Dist
		}
	}
	return
}

func (r *RayTracer) testRay(ray vector.Ray, scene scene.Scene) float64 {
	if isect := r.intersections(ray, scene); isect.Thing != nil {
		return isect.Dist
	} else {
		return math.Inf(1)
	}
}

func (r *RayTracer) traceRay(ray vector.Ray, scene scene.Scene, depth uint8) color.Color {
	if isect := r.intersections(ray, scene); isect.Thing != nil {
		return r.shade(isect, scene, depth)
	} else {
		return color.Background
	}
}

func (r *RayTracer) shade(isect thing.Intersection, scene scene.Scene, depth uint8) color.Color {
	d := isect.Ray.Dir
	pos := vector.Plus(vector.Times(isect.Dist, d), isect.Ray.Start)
	normal := isect.Thing.Normal(pos)
	reflectDir := vector.Minus(d, vector.Times(2, vector.Times(vector.Dot(normal, d), normal)))
	naturalColor := color.Plus(color.Background, r.getNaturalColor(isect.Thing, pos, normal, reflectDir, scene))
	reflectedColor := color.Grey
	if depth < maxDepth {
		reflectedColor = r.getReflectionColor(isect.Thing, pos, normal, reflectDir, scene, depth)
	}
	return color.Plus(naturalColor, reflectedColor)
}

func (r *RayTracer) getReflectionColor(thing thing.Thing, pos vector.Vec, normal vector.Vec, rd vector.Vec, scene scene.Scene, depth uint8) color.Color {
	return color.Scale(thing.GetSurface().Reflect(pos), r.traceRay(vector.Ray{Start: pos, Dir: rd}, scene, depth+1))
}

func (r *RayTracer) getNaturalColor(thing thing.Thing, pos vector.Vec, norm vector.Vec, rd vector.Vec, scene scene.Scene) color.Color {
	addLight := func(col color.Color, light light.Light, idx int) color.Color {
		ldis := vector.Minus(light.Pos, pos)
		livec := vector.Norm(ldis)
		isInShadow := false
		if neatIsect := r.testRay(vector.Ray{Start: pos, Dir: livec}, scene); neatIsect <= vector.Mag(ldis) {
			isInShadow = true
		}
		if isInShadow {
			return col
		} else {
			illum := vector.Dot(livec, norm)
			lcolor := color.DefaultColor
			if illum > 0 {
				lcolor = color.Scale(illum, light.Color)
			}
			specular := vector.Dot(livec, vector.Norm(rd))
			scolor := color.DefaultColor
			if specular > 0 {
				scolor = color.Scale(math.Pow(specular, thing.GetSurface().GetRoughness()), light.Color)
			}
			return color.Plus(col, color.Plus(color.Times(thing.GetSurface().Diffuse(pos), lcolor), color.Times(thing.GetSurface().Specular(pos), scolor)))
		}
	}
	lightColor, _ := util.Reduce(scene.Lights, color.DefaultColor, addLight)
	return lightColor.(color.Color)
}

func (r *RayTracer) Render(scene scene.Scene, screenWidth int, screenHeight int) *image.RGBA {
	getPoint := func(x int, y int, camera *camera.Camera) vector.Vec {
		recenterX := func(x int) float64 {
			return (float64(x) - float64(screenWidth)/2.0) / 2.0 / float64(screenWidth)
		}
		recenterY := func(y int) float64 {
			return -(float64(y) - float64(screenHeight)/2.0) / 2.0 / float64(screenHeight)
		}
		return vector.Norm(vector.Plus(camera.Forward, vector.Plus(vector.Times(recenterX(x), camera.Right), vector.Times(recenterY(y), camera.Up))))
	}
	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	idx := 0
	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			c := r.traceRay(vector.Ray{Start: scene.Camera.Pos, Dir: getPoint(x, y, scene.Camera)}, scene, 0)
			rgb := color.ToDrawingColor(c)
			img.Pix[idx] = byte(rgb.R)
			img.Pix[idx+1] = byte(rgb.G)
			img.Pix[idx+2] = byte(rgb.B)
			img.Pix[idx+3] = 255
			idx += 4
		}
	}
	return img
}
