package main
import (
	"math"
)

type Ray struct {
	Origin, Direction Vec3
}

func (r Ray) Point(t float64) Vec3 {
	b := r.Direction.MultiplyScalar(t)
	return r.Origin.Add(b)
}



func (v Vec3) Reflect(ov Vec3) Vec3 {
	b := 2 * v.Dot(ov)
	return v.Subtract(ov.MultiplyScalar(b))
}

func (v Vec3) Refract(ov Vec3, n float64) (bool, Vec3) {
	uv := v.Normalize()
	uo := ov.Normalize()
	dt := uv.Dot(uo)
	discriminant := 1.0 - (n * n * (1 - dt*dt))
	if discriminant > 0 {
		a := uv.Subtract(ov.MultiplyScalar(dt)).MultiplyScalar(n)
		b := ov.MultiplyScalar(math.Sqrt(discriminant))
		return true, a.Subtract(b)
	}
	return false, Vec3{}
}


type Hit struct {
	T             float64
	Point, Normal Vec3
	Material
}

type Hitable interface {
	Hit(r Ray, tMin float64, tMax float64) (bool, Hit)
}

type Sphere struct {
	Center Vec3
	Radius float64
	Material
}

func NewSphere(x, y, z, radius float64, m Material) *Sphere {
	return &Sphere{Vec3{x, y, z}, radius, m}
}

func (s *Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c

	hit := Hit{Material: s.Material}

	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			hit.T = temp
			hit.Point = r.Point(temp)
			hit.Normal = hit.Point.Subtract(s.Center).DivideScalar(s.Radius)
			return true, hit
		}
		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			hit.T = temp
			hit.Point = r.Point(temp)
			hit.Normal = hit.Point.Subtract(s.Center).DivideScalar(s.Radius)
			return true, hit
		}
	}
	return false, Hit{}
}

type World []Hitable

func (w *World) Add(h Hitable) {
	*w = append(*w, h)
}




func (w *World) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
	hitAnything := false
	closest := tMax
	record := Hit{}

	for _, element := range *w {
		if element != nil {
			hit, tempRecord := element.Hit(r, tMin, closest)

			if hit {
				hitAnything = true
				closest = tempRecord.T
				record = tempRecord
			}
		}
	}
	return hitAnything, record
}