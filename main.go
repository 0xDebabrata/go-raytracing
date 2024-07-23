package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
    imgWidth := 256
    imgHeight := 256

    var builder strings.Builder
    builder.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imgWidth, imgHeight))

    for i := 0; i < imgHeight; i++ {
        log.Printf("Scanlines remaining: %d\n", imgHeight - i)
        for j := 0; j < imgWidth; j++ {
            r := float64(i) / float64(imgWidth - 1)
            g := float64(j) / float64(imgWidth - 1)
            b := 0.0
            builder.WriteString(fmt.Sprintf("%d %d %d\n", int(r * 255), int(g * 255), int(b)))
        }
    }
    log.Printf("Done\n")

    fmt.Print(builder.String())
}
