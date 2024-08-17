package raytracing

type Ray struct {
	Origin *Point
	Dir    *Vec3
}

func (ray *Ray) At(t float32) *Vec3 {
	direction := ScalarMultiply(t, ray.Dir)
	return Add(ray.Origin, direction)
}

func RayColor(ray *Ray, world Hittable) *Color {
	var rec HitRecord

	if world.hit(ray, 0, float32(Infinity), &rec) {
		var normal *Vec3 = rec.Normal
		return NewColor(0.5*(normal.x()+1), 0.5*(normal.y()+1), 0.5*(normal.z()+1))
	}

	var unitDir *Vec3 = UnitVector(ray.Dir)
	// -1 < y < 1 but we need 0 <= a <= 1
	a := 0.5 * (unitDir.y() + 1.0)
	return Add(
		ScalarMultiply(1-a, NewColor(1.0, 1.0, 1.0)),
		ScalarMultiply(a, NewColor(0.5, 0.7, 1.0)),
	)
}
