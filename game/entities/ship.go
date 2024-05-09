package entities

import (
	"gRocketeer/core"
	"gRocketeer/core/entity"
	"gRocketeer/game/resources"
	"gRocketeer/game/types"
	"gRocketeer/game/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type Ship struct {
	entity.BaseEntity
	g types.IGame

	body *cp.Body

	spShip *core.Sprite
}

func NewShip(g types.IGame) *Ship {
	return &Ship{
		g: g,
	}
}

func (s *Ship) Init() {
	s.spShip = core.NewSprite(resources.ImgShip)

	space := s.g.GetSpace()

	s.body = cp.NewBody(1, cp.MomentForBox(1, 10, 20)*2)

	space.AddBody(s.body)

	shape := cp.NewBox(s.body, s.spShip.Rect.Bounds.X, s.spShip.Rect.Bounds.Y, 0.5)
	shape.SetElasticity(1.5)
	shape.SetFriction(2)
	space.AddShape(shape)

	s.body.SetPosition(cp.Vector{X: s.g.Screen().X / 2, Y: s.g.Screen().Y / 2})
}

func (s *Ship) Draw(target *ebiten.Image) {
	s.spShip.Draw(target)
}

func (s *Ship) Update() bool {
	s.spShip.Rect.Pos = util.FromVec(s.body.Position())
	s.spShip.Rot = s.body.Angle()
	return false
}

func (s *Ship) Thrust() {
	s.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -100}, cp.Vector{X: 0, Y: -10})
}

func (s *Ship) Turn(clockWise bool) {
	if clockWise {
		s.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -20}, cp.Vector{X: -10, Y: 14})
	} else {
		s.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -20}, cp.Vector{X: 10, Y: 14})
	}
}
