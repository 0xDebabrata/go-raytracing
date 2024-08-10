package raytracing

import (
	"fmt"
	"math"
)

type Vec3 struct {
	e [3]float32
}

type Color = Vec3
type Point = Vec3

func New(x float32, y float32, z float32) *Vec3 {
	return &Vec3{[3]float32{x, y, z}}
}

func NewColor(x float32, y float32, z float32) *Color {
	return &Color{[3]float32{x, y, z}}
}

func (v *Vec3) x() float32 {
	return v.e[0]
}

func (v *Vec3) y() float32 {
	return v.e[1]
}

func (v *Vec3) z() float32 {
	return v.e[2]
}

func (v *Vec3) LengthSquared() float32 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (c *Color) WriteColor() {
	rByte := int(255 * c.x())
	gByte := int(255 * c.y())
	bByte := int(255 * c.z())
	fmt.Printf("%d %d %d\n", rByte, gByte, bByte)
}

func Add(vectors ...*Vec3) *Vec3 {
	var (
		x float32
		y float32
		z float32
	)
	for idx := range vectors {
		x += vectors[idx].x()
		y += vectors[idx].y()
		z += vectors[idx].z()
	}
	return New(x, y, z)
}

func ScalarMultiply(a float32, v *Vec3) *Vec3 {
	return New(a*v.x(), a*v.y(), a*v.z())
}

func UnitVector(v *Vec3) *Vec3 {
	magnitude := float32(math.Sqrt(float64(v.x()*v.x() + v.y()*v.y() + v.z()*v.z())))
	return ScalarMultiply(1.0/magnitude, v)
}

func Dot(a *Vec3, b *Vec3) float32 {
	return a.x()*b.x() + a.y()*b.y() + a.z()*b.z()
}
