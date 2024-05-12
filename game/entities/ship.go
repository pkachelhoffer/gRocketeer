package entities

import (
	"gRocketeer/core"
	"gRocketeer/core/entity"
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

	comps []IShipComp
}

func NewShip(g types.IGame) *Ship {
	return &Ship{
		g: g,
	}
}

func (s *Ship) Init() {
	s.body.SetPosition(cp.Vector{X: s.g.Screen().X / 2, Y: s.g.Screen().Y / 2})
}

func (s *Ship) Draw(target *ebiten.Image) {
	for _, c := range s.comps {
		c.Draw(target)
	}
}

func (s *Ship) Update() bool {
	for _, c := range s.comps {
		c.Update()
	}
	return false
}

func (s *Ship) AddComponent(comp IShipComp) {
	s.comps = append(s.comps, comp)
}

func (s *Ship) SetBody(b *cp.Body) {
	s.body = b
}

func (s *Ship) Thrust(direction core.Vec) {
	s.body.ApplyForceAtLocalPoint(cp.Vector{X: direction.X * 100, Y: direction.Y * 100}, cp.Vector{X: 0, Y: -10})
}

func (s *Ship) SetPosition(pos core.Vec) {
	s.body.SetPosition(util.ToVec(pos))
}
