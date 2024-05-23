package core

import (
	"fmt"
	"math"
)

type Vec struct {
	X float64
	Y float64
}

func NewVec(x float64, y float64) Vec {
	return Vec{
		X: x,
		Y: y,
	}
}

func (v Vec) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vec) IsNegative() bool {
	return v.X < 0 || v.Y < 0
}

func (v Vec) Sub(v2 Vec) Vec {
	return NewVec(v.X-v2.X, v.Y-v2.Y)
}

func (v Vec) Add(v2 Vec) Vec {
	return NewVec(v.X+v2.X, v.Y+v2.Y)
}

func (v Vec) Div(v2 float64) Vec {
	return NewVec(v.X/v2, v.Y/v2)
}

func (v Vec) Mul(v2 Vec) Vec {
	return NewVec(v.X*v2.X, v.Y*v2.Y)
}

func (v Vec) String() string {
	return fmt.Sprintf("{posX: %f, posY: %f}", v.X, v.Y)
}

func (v Vec) Distance(b Vec) float64 {
	x := b.X - v.X
	y := b.Y - v.Y

	zi := math.Sqrt(Sqr(x) + Sqr(y))
	return zi
}

func (v Vec) Normalise() Vec {
	zi := math.Sqrt(Sqr(v.X) + Sqr(v.Y))
	if zi == 0 {
		return NewVec(0, 0)
	}
	return NewVec(v.X/zi, v.Y/zi)
}

func (v Vec) LeftFromCenter() float64 {
	return -(v.X / 2)
}

func (v Vec) RightFromCenter() float64 {
	return v.X / 2
}

func (v Vec) UpFromCenter() float64 {
	return -(v.Y / 2)
}

func (v Vec) DownFromCenter() float64 {
	return v.Y / 2
}

func Sqr(x float64) float64 {
	return x * x
}
