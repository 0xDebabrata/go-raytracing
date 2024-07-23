package main

import (
	"fmt"
	"log"
	"raytracing/raytracing"
)

func main() {
    imgWidth := 256
    imgHeight := 256

    fmt.Printf("P3\n%d %d\n255\n", imgWidth, imgHeight)

    for i := 0; i < imgHeight; i++ {
        log.Printf("Scanlines remaining: %d\n", imgHeight - i)
        for j := 0; j < imgWidth; j++ {
            r := float32(i) / float32(imgWidth - 1)
            g := float32(j) / float32(imgWidth - 1)
            b := float32(0.0)

            var pixel raytracing.Color = raytracing.New(r, g, b)
            pixel.WriteColor()
        }
    }
    log.Printf("Done\n")
}
