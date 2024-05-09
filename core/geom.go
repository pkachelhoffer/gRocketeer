package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type TransOpt struct {
	Scale Vec
	Rot   float64
}

func Trans(g *ebiten.GeoM, size Vec, pos Vec, o TransOpt) {
	if o.Scale.X != 1 || o.Scale.Y != 1 {
		g.Scale(o.Scale.X, o.Scale.Y)
	}

	g.Translate(-size.X/2, -size.Y/2)

	if o.Rot != 0 {
		g.Rotate(o.Rot)
		//g.Rotate(o.Rot * math.Pi / 180)
	}

	g.Translate(pos.X, pos.Y)
}
