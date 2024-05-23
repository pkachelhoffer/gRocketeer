package cpc

import (
	"gRocketeer/core"
	"gRocketeer/game/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type CSprite struct {
	Body  *cp.Body
	Shape *cp.Shape

	space *cp.Space
	sp    *core.Sprite
}

func NewCSprite(img *ebiten.Image, space *cp.Space, mass float64, moment float64) CSprite {
	sp := core.NewSprite(img)
	body := cp.NewBody(mass, moment)
	space.AddBody(body)

	shape := cp.NewBox(body, sp.Rect.Bounds.X, sp.Rect.Bounds.Y, 0.1)
	space.AddShape(shape)

	return CSprite{
		Body:  body,
		Shape: shape,
		sp:    sp,
		space: space,
	}
}

func (cs CSprite) Draw(target *ebiten.Image) {
	cs.sp.Draw(target)
}

func (cs CSprite) Update() {
	cs.sp.Rect.Pos = util.FromVec(cs.Body.Position())
}

func (cs CSprite) SetPosition(pos cp.Vector) {
	cs.Body.SetPosition(pos)
	cs.sp.Rect.Pos = util.FromVec(cs.Body.Position())
}

func (cs CSprite) Terminate() {
	cs.space.RemoveBody(cs.Body)
	cs.space.RemoveShape(cs.Shape)
}
