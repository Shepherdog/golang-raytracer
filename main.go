package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	. "raytracer"
	"raytracer/camera"
	"raytracer/color"
	"raytracer/light"
	"raytracer/scene"
	"raytracer/surface"
	"raytracer/thing"
	"raytracer/vector"
	"time"
)

func main() {
	scene := scene.Scene{
		Things: []thing.Thing{
			&thing.Plane{Offset: 0.0, Norm: vector.Vec{0.0, 1.0, 0.0}, Surface: surface.Checkerboard{Roughness: 150}},
			&thing.Sphere{Radius: 1.0, Center: vector.Vec{0.0, 1.0, -0.25}, Surface: surface.Shiny{Roughness: 250}},
			&thing.Sphere{Radius: 0.5, Center: vector.Vec{-1.0, 0.5, 1.5}, Surface: surface.Shiny{Roughness: 250}},
		},
		Lights: []light.Light{
			{Pos: vector.Vec{-2.0, 2.5, 0.0}, Color: color.Color{0.49, 0.07, 0.07}},
			{Pos: vector.Vec{1.5, 2.5, 1.5}, Color: color.Color{0.07, 0.07, 0.49}},
			{Pos: vector.Vec{1.5, 2.5, -1.5}, Color: color.Color{0.07, 0.49, 0.071}},
			{Pos: vector.Vec{0.0, 3.5, 0.0}, Color: color.Color{0.21, 0.21, 0.35}},
		},
		Camera: camera.New(vector.Vec{3.0, 2.0, 4.0}, vector.Vec{-1.0, 0.5, 0.0}),
	}
	rayTracer := &RayTracer{}

	t1 := time.Now()
	fmt.Println("rendering...")
	var screenWidth int = 256
	var screenHeight int = 256
	img := rayTracer.Render(scene, screenWidth, screenHeight)
	duration := time.Since(t1).Milliseconds()
	fmt.Printf("done, used %d ms\n", duration)

	outputFile, err := os.Create("output.png")
	if err == nil {
		defer outputFile.Close()
	} else {
		log.Fatal("failed to create file")
	}
	png.Encode(outputFile, img)
	fmt.Println("saved to output.png")
}
