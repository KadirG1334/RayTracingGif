package main

import (
	"image"
	"math"
	"math/rand"
	"sync"
	"time"

	
)

const (
	maxDepth = 50
	tMin     = 0.001
)

func color(ray Ray, hitable Hitable, rnd *rand.Rand, depth int) Color {
	hit, record := hitable.Hit(ray, tMin, math.MaxFloat64)

	if hit {
		if depth < maxDepth {
			bounced, bouncedRay := record.Bounce(ray, record, rnd)
			if bounced {
				newColor := color(bouncedRay, hitable, rnd, depth+1)
				return record.Material.Color().Multiply(newColor)
			}
		}
		return Black
	}

	return Gradient(White, Blue, ray.Direction.Normalize().Y) //for background
}

// sample samples rays for anti-aliasing
func sample(hitable Hitable, camera *Camera, rnd *rand.Rand, samples, width, height, i, j int) Color {
	rgb := Color{}

	for s := 0; s < samples; s++ {
		u := (float64(i) + rnd.Float64()) / float64(width)
		v := (float64(j) + rnd.Float64()) / float64(height)

		ray := camera.RayAt(u, v, rnd)
		rgb = rgb.Add(color(ray, hitable, rnd, 0))
	}

	// average
	return rgb.DivideScalar(float64(samples))
}

// Do performs the render, sampling each pixel the provided number of times
func Do(hitable Hitable, camera *Camera, cpus, samples, width, height int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup

	for i := 0; i < cpus; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

			for row := i; row < height; row += cpus {
				for col := 0; col < width; col++ {
					rgb := sample(hitable, camera, rnd, samples, width, height, col, row)
					img.Set(col, height-row-1, rgb)
				}
				
			}
		}(i)
	}

	wg.Wait()
	return img
}