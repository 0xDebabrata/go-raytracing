package raytracing

import "math"

type HitRecord struct {
	P      *Point  // Point at which ray hits the object
	Normal *Vec3   // Normal to the surface of the object at the point of intersection
	T      float32 // Value of t in the equation P = origin + dir.t
}

type Hittable interface {
	hit(*Ray, float32, float32, *HitRecord) bool
}

type Sphere struct {
	center *Point
	radius float32
}

func (s *Sphere) hit(ray *Ray, tMin float32, tMax float32, rec *HitRecord) bool {
	var oc *Vec3 = Add(s.center, ScalarMultiply(-1, ray.Origin))
	a := ray.Dir.LengthSquared()
	b := -2 * Dot(ray.Dir, oc)
	c := oc.LengthSquared() - s.radius*s.radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return false
	}

	t := (-b - float32(math.Sqrt(float64(discriminant)))) / 2 * a
	if t <= tMin || tMax <= t {
		t = (-b + float32(math.Sqrt(float64(discriminant)))) / 2 * a
		if t <= tMin || tMax <= t {
			return false
		}
	}

	rec.T = t
	rec.P = ray.At(rec.T)
	rec.Normal = Add(rec.P, ScalarMultiply(-1, s.center))
	// Normalise normal vector
	radiusInverse := 1 / s.radius
	rec.Normal = ScalarMultiply(radiusInverse, rec.Normal)

	return true
}
