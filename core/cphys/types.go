package cphys

import (
	"playground/core"
)

type Circle struct {
	Radius   float64
	Position core.Vec
}

func NewCircle(pos core.Vec, rad float64) Circle {
	return Circle{
		Radius:   rad,
		Position: pos,
	}
}

func (a Circle) CollideCirc(b Circle) bool {
	r := core.Sqr(a.Radius + b.Radius)
	z := core.Sqr(a.Position.X+b.Position.X) + core.Sqr(a.Position.Y+b.Position.Y)

	if z > r {
		return false
	}

	return true
}
