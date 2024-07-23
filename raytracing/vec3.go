package raytracing

import "fmt"

type vec3 struct {
    e [3]float32
}

type Color = vec3
type Point = vec3

func New(x float32, y float32, z float32) vec3 {
    return vec3{[3]float32{x, y, z}}
}

func (v *vec3) x() float32 {
    return v.e[0]
}

func (v *vec3) y() float32 {
    return v.e[1]
}

func (v *vec3) z() float32 {
    return v.e[2]
}

func (c *Color) WriteColor() {
    rByte := int(255 * c.x())
    gByte := int(255 * c.y())
    bByte := int(255 * c.z())
    fmt.Printf("%d %d %d\n", rByte, gByte, bByte)
}
