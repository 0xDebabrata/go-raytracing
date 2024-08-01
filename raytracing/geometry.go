package raytracing

func hitSphere(center *Point, radius float32, ray *Ray) bool {
	var oc *Vec3 = Add(center, ScalarMultiply(-1, ray.Origin))
	a := Dot(ray.Dir, ray.Dir)
	b := -2 * Dot(ray.Dir, oc)
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c

	return discriminant >= 0
}
