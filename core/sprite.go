package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Sprite struct {
	Img *ebiten.Image

	Rect Rectangle

	Rot   float64
	Scale Vec

	ColorM colorm.ColorM

	Geom ebiten.GeoM
}

func NewSprite(img *ebiten.Image) *Sprite {
	b := img.Bounds()
	s := &Sprite{
		Img:   img,
		Rect:  NewRect(NewVec(0, 0), NewVec(float64(b.Dx()), float64(b.Dy()))),
		Scale: NewVec(1, 1),
	}

	s.ColorM.Scale(1, 1, 1, 1)

	return s
}

func (s *Sprite) Draw(tgt *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	Trans(&op.GeoM, s.Rect.Bounds, s.Rect.Pos, TransOpt{
		Scale: s.Scale,
		Rot:   s.Rot,
	})

	op.GeoM.Concat(s.Geom)

	colorm.DrawImage(tgt, s.Img, s.ColorM, &colorm.DrawImageOptions{
		GeoM: op.GeoM,
	})
}
