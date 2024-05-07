package core

import (
	"fmt"
)

type Rectangle struct {
	Pos    Vec
	Bounds Vec
}

func NewRect(pos Vec, bounds Vec) Rectangle {
	return Rectangle{
		Pos:    pos,
		Bounds: bounds,
	}
}

func (a Rectangle) Min() Vec {
	return a.Pos.Sub(a.Bounds.Div(2))
}

func (a Rectangle) Max() Vec {
	return a.Pos.Add(a.Bounds.Div(2))
}

func (a Rectangle) CollideRect(b Rectangle) bool {
	amin := a.Min()
	amax := a.Max()

	bmin := b.Min()
	bmax := b.Max()

	if amax.X < bmin.X || amin.X > bmax.X {
		return false
	}

	if amax.Y < bmin.Y || amin.Y > bmax.Y {
		return false
	}

	return true
}

func (a Rectangle) BottomLeft() Vec {
	return NewVec(a.Pos.X+a.Bounds.LeftFromCenter(), a.Pos.Y+a.Bounds.DownFromCenter())
}

func (a Rectangle) BottomRight() Vec {
	return NewVec(a.Pos.X+a.Bounds.RightFromCenter(), a.Pos.Y+a.Bounds.DownFromCenter())
}

func (a Rectangle) Top() Vec {
	return NewVec(a.Pos.X, a.Pos.Y-(a.Bounds.Y/2.0))
}

func (a Rectangle) Bottom() Vec {
	return NewVec(a.Pos.X, a.Pos.Y+(a.Bounds.Y/2.0))
}

func (a Rectangle) String() string {
	return fmt.Sprintf("{min: %s, max: %s}", a.Pos, a.Bounds)
}
