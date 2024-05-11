package entities

import (
	"gRocketeer/core"
	"gRocketeer/core/entity"
	"gRocketeer/game/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

const (
	CompSize = 5
)

var (
	CompSizeHalf = core.NewVec(CompSize/2, CompSize/2)
)

type IShipComp interface {
	entity.Entity

	SetPhys(b *cp.Body, sh *cp.Shape)
}

type ShipCompHull struct {
	entity.BaseEntity

	b  *cp.Body
	sh *cp.PolyShape

	spHull *core.Sprite
}

func (sch *ShipCompHull) SetPhys(b *cp.Body, sh *cp.Shape) {
	sch.b = b
	sch.sh = sh.Class.(*cp.PolyShape)
}

func NewShipCompHull(sp *core.Sprite) *ShipCompHull {
	return &ShipCompHull{
		spHull: sp,
	}
}

func (sch *ShipCompHull) Update() bool {
	lt := sch.sh.Vert(0)
	pos := util.FromVec(sch.b.Position()).Add(core.NewVec(lt.X, lt.Y)).Add(CompSizeHalf)

	sch.spHull.Rect.Pos = pos

	return false
}

func (sch *ShipCompHull) Draw(target *ebiten.Image) {
	sch.spHull.Draw(target)
}
