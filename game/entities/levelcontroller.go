package entities

import (
	"gRocketeer/core"
	"gRocketeer/core/entity"
	"gRocketeer/game/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
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
	ship := CreateShip01(lvl.g)
	ship.SetPosition(lvl.g.Screen().Div(2))
	lvl.g.GetSys().Add(mainShip, ship)
}

func (lvl *LvlCon) Update() bool {
	lvl.Input()
	return false
}

func (lvl *LvlCon) Input() {
	ship := entity.Get[*Ship01](lvl.g.GetSys(), mainShip)
	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		ship.Thrust(core.NewVec(0, -3))
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		ship.Thrust(core.NewVec(0, 3))
	}
	if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		ship.Thrust(core.NewVec(-3, 0))
	}
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		ship.Thrust(core.NewVec(3, 0))
	}

	cx, cy := ebiten.CursorPosition()
	ship.Aim(cp.Vector{X: float64(cx), Y: float64(cy)})
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		ship.Shoot()
	}
}
