package raytracing

type HittableList struct {
	objects []Hittable
}

func (hl *HittableList) hit(ray *Ray, tMin float32, tMax float32, hitRecord *HitRecord) bool {
	var (
		tempRec      HitRecord
		hitAnything  = false
		closestSoFar = tMax
	)

	for idx := range hl.objects {
		if hl.objects[idx].hit(ray, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*hitRecord = tempRec
		}
	}

	return hitAnything
}

func (hl *HittableList) Add(object Hittable) {
	hl.objects = append(hl.objects, object)
}
