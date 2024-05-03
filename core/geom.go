package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type TransOpt struct {
	Scale float64
	Rot   float64
}

func Trans(g *ebiten.GeoM, size Vec, pos Vec, o TransOpt) {
	g.Translate(-size.X/2, -size.Y/2)

	if o.Rot != 0 {
		g.Rotate(o.Rot)
		//g.Rotate(o.Rot * math.Pi / 180)
	}

	if o.Scale != 1 {
		g.Scale(o.Scale, o.Scale)
	}

	g.Translate(pos.X, pos.Y)
}
