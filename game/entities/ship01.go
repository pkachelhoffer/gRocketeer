package entities

import (
	"gRocketeer/core"
	"gRocketeer/core/cpc"
	"gRocketeer/core/entity"
	"gRocketeer/game/resources"
	"gRocketeer/game/types"
	"gRocketeer/game/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type Ship01 struct {
	entity.BaseEntity

	g types.IGame

	csp cpc.CSprite

	aim cp.Vector
}

func CreateShip01(g types.IGame) *Ship01 {
	csp := cpc.NewCSprite(resources.ImgShip01, g.GetSpace(), 10, cp.INFINITY)
	csp.Shape.SetFilter(cp.ShapeFilter{Categories: types.CatShip})

	return &Ship01{
		g:   g,
		csp: csp,
	}
}

func (s *Ship01) Draw(target *ebiten.Image) {
	//vector.StrokeLine(target, float32(s.sp.Rect.Pos.X), float32(s.sp.Rect.Pos.Y), 10, 10, 1, types.ColorWhite, true)
	s.csp.Draw(target)
}

func (s *Ship01) Update() bool {
	s.csp.Update()
	return false
}

func (s *Ship01) Thrust(direction core.Vec) {
	s.csp.Body.ApplyForceAtLocalPoint(cp.Vector{X: direction.X * 100, Y: direction.Y * 100}, cp.Vector{X: 0, Y: 0})
}

func (s *Ship01) SetPosition(pos core.Vec) {
	s.csp.Body.SetPosition(util.ToVec(pos))
}

func (s *Ship01) Aim(aim cp.Vector) {
	s.aim = aim.Sub(s.csp.Body.Position()).Normalize()
}

func (s *Ship01) Shoot() {
	bul := NewBullet01(s.g)
	bul.csp.SetPosition(s.csp.Body.Position())
	bul.csp.Body.ApplyForceAtWorldPoint(s.aim.Mult(1000), cp.Vector{})
	s.g.GetSys().Add(core.RandomID(), bul)
}

func (s *Ship01) Terminate() {
	s.csp.Terminate()
}
