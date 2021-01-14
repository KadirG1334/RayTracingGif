package main

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float64
}

var UnitVector = Vec3{1, 1, 1}

func VectorInUnitSphere(rnd *rand.Rand) Vec3 {
	for {
		r := Vec3{rnd.Float64(), rnd.Float64(), rnd.Float64()}
		p := r.MultiplyScalar(2.0).Subtract(UnitVector)
		if p.SquaredLength() >= 1.0 {
			return p
		}
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Normalize() Vec3 {
	return v.DivideScalar(v.Length())
}

func (v Vec3) Dot(ov Vec3) float64 {
	return v.X*ov.X + v.Y*ov.Y + v.Z*ov.Z
}

func (v Vec3) Cross(ov Vec3) Vec3{
	return Vec3{
		v.Y*ov.Z - v.Z*ov.Y,
		v.Z*ov.X - v.X*ov.Z,
		v.X*ov.Y - v.Y*ov.X,
	}
}

func (v Vec3) Add(ov Vec3) Vec3 {
	return Vec3{v.X + ov.X, v.Y + ov.Y, v.Z + ov.Z}
}

func (v Vec3) Subtract(ov Vec3) Vec3 {
	return Vec3{v.X - ov.X, v.Y - ov.Y, v.Z - ov.Z}
}

func (v Vec3) Multiply(ov Vec3) Vec3 {
	return Vec3{v.X * ov.X, v.Y * ov.Y, v.Z * ov.Z}
}

func (v Vec3) Divide(ov Vec3) Vec3 {
	return Vec3{v.X / ov.X, v.Y / ov.Y, v.Z / ov.Z}
}

func (v Vec3) AddScalar(t float64) Vec3{
	return Vec3{v.X + t, v.Y + t, v.Z + t}
}

func (v Vec3) SubtractScalar(t float64) Vec3 {
	return Vec3{v.X - t, v.Y - t, v.Z - t}
}

func (v Vec3) MultiplyScalar(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) DivideScalar(t float64) Vec3 {
	return Vec3{v.X / t, v.Y / t, v.Z / t}
}

