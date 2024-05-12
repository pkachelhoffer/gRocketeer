package entities

import (
	"gRocketeer/core"
	"gRocketeer/core/entity"
	"gRocketeer/game/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const mainShip = "MainShip"

type LvlCon struct {
	entity.BaseEntity

	g types.IGame
}

func NewLvlCon(g types.IGame) *LvlCon {
	return &LvlCon{
		g: g,
	}
}

func (lvl *LvlCon) Init() {
	ship := CreateShip(lvl.g, ShipPlan1)
	ship.SetPosition(lvl.g.Screen().Div(2))
	lvl.g.GetSys().Add(mainShip, ship)
}

func (lvl *LvlCon) Update() bool {
	lvl.Input()
	return false
}

func (lvl *LvlCon) Input() {
	ship := entity.Get[*Ship](lvl.g.GetSys(), mainShip)
	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		ship.Thrust(core.NewVec(0, -1))
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		ship.Thrust(core.NewVec(0, 1))
	}
	if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		ship.Thrust(core.NewVec(-1, 0))
	}
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		ship.Thrust(core.NewVec(1, 0))
	}
}
