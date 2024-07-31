package main

import (
	"fmt"
	"log"
	"raytracing/raytracing"
)

func main() {
	var (
		aspectRatio float32 = 16.0 / 9.0
		imgWidth            = 400
		imgHeight           = int(float32(imgWidth) / aspectRatio)
	)
	// Ensure height is at least 1
	if imgHeight < 1 {
		imgHeight = 1
	}

	var (
		// Camera
		focalLength    float32           = 1.0
		viewportHeight float32           = 2.0
		viewportWidth                    = viewportHeight * (float32(imgWidth) / float32(imgHeight))
		cameraCenter   *raytracing.Point = raytracing.New(0, 0, 0)

		// Viewport
		// Right-handed coordinate system
		viewportU   = raytracing.New(viewportWidth, 0, 0)
		viewportV   = raytracing.New(0, -viewportHeight, 0)
		pixelDeltaU = raytracing.ScalarMultiply(1.0/float32(imgWidth), viewportU)
		pixelDeltaV = raytracing.ScalarMultiply(1.0/float32(imgHeight), viewportV)

		// Upper left pixel
		viewportUpperLeft = raytracing.Add(
			cameraCenter,                               // Camera position
			raytracing.New(0, 0, -focalLength),         // Point orthogonal to camera on viewport
			raytracing.ScalarMultiply(-0.5, viewportU), // Shift left
			raytracing.ScalarMultiply(-0.5, viewportV), // Shift up
		)
		// Pixels are inset into the viewport by half the inter pixel distance
		pixel00Loc *raytracing.Point = raytracing.Add(
			viewportUpperLeft,
			raytracing.ScalarMultiply(0.5, pixelDeltaU),
			raytracing.ScalarMultiply(0.5, pixelDeltaV),
		)
	)

	fmt.Printf("P3\n%d %d\n255\n", imgWidth, imgHeight)

	for i := 0; i < imgHeight; i++ {
		log.Printf("Scanlines remaining: %d\n", imgHeight-i)
		for j := 0; j < imgWidth; j++ {
			var pixelCenter *raytracing.Point = raytracing.Add(
				pixel00Loc,
				raytracing.ScalarMultiply(float32(j), pixelDeltaU),
				raytracing.ScalarMultiply(float32(i), pixelDeltaV),
			)
			var rayDir *raytracing.Vec3 = raytracing.Add(
				pixelCenter,
				raytracing.ScalarMultiply(-1, cameraCenter),
			)
			ray := raytracing.Ray{
				Origin: cameraCenter,
				Dir:    rayDir,
			}

			var pixelColor *raytracing.Color = raytracing.RayColor(&ray)
			pixelColor.WriteColor()
		}
	}
	log.Printf("Done\n")
}
