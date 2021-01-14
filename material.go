package main
import (
	
	"math/rand"
)

type Material interface {
	Bounce(input Ray, hit Hit, rnd *rand.Rand) (bool, Ray)
	Color() Color
}

type Lambertian struct {
	Attenuation Color
}

func (l Lambertian) Color() Color {
	return l.Attenuation
}

func (l Lambertian) Bounce(input Ray, hit Hit, rnd *rand.Rand) (bool, Ray) {
	direction := hit.Normal.Add(VectorInUnitSphere(rnd))
	return true, Ray{hit.Point, direction}
}

type Metal struct {
	Attenuation Color
	Fuzz        float64
}

func (m Metal) Color() Color {
	return m.Attenuation
}

func (m Metal) Bounce(input Ray, hit Hit, rnd *rand.Rand) (bool, Ray) {
	direction := input.Direction.Reflect(hit.Normal)
	bouncedRay := Ray{hit.Point, direction.Add(VectorInUnitSphere(rnd).MultiplyScalar(m.Fuzz))}
	bounced := direction.Dot(hit.Normal) > 0
	return bounced, bouncedRay
}
