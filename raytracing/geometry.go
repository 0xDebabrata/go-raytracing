package raytracing

import "math"

func hitSphere(center *Point, radius float32, ray *Ray) float32 {
	var oc *Vec3 = Add(center, ScalarMultiply(-1, ray.Origin))
	a := ray.Dir.LengthSquared()
	b := -2 * Dot(ray.Dir, oc)
	c := oc.LengthSquared() - radius*radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - float32(math.Sqrt(float64(discriminant)))) / 2 * a
	}
}
