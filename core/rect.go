package core

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
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

func (a Rectangle) String() string {
	return fmt.Sprintf("{min: %s, max: %s}", a.Pos, a.Bounds)
}

func DrawRect(rect Rectangle, target *ebiten.Image, col color.Color) {
	vector.DrawFilledRect(target, float32(rect.Pos.X-(rect.Bounds.X/2.0)), float32(rect.Pos.Y-(rect.Bounds.Y/2.0)), float32(rect.Bounds.X), float32(rect.Bounds.Y), col, false)
}
