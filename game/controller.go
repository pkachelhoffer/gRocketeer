package game

import (
	"gRocketeer/core"
	"gRocketeer/game/entities"
	"gRocketeer/game/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Controller struct {
	g types.IGame

	ship *entities.Ship
}

func NewController(g types.IGame) *Controller {
	return &Controller{g: g}
}

func (c *Controller) ShowGame() {
	c.ship = entities.CreateShip(c.g, entities.ShipPlan1)

	c.g.GetSys().Add("ship", c.ship)

	w := entities.NewWall(c.g, core.NewRect(core.NewVec(c.g.Screen().RightFromCenter(), c.g.Screen().Y-20), core.NewVec(100, 5)), types.ColorWall)
	c.g.GetSys().Add("wallLeft", w)
}

func (c *Controller) Input() {
	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		c.ship.Thrust()
	} else {
		if inpututil.KeyPressDuration(ebiten.KeyE) > 0 {
			c.ship.Turn(true)
		}
		if inpututil.KeyPressDuration(ebiten.KeyQ) > 0 {
			c.ship.Turn(false)
		}
	}
}

func (c *Controller) Update() {
	c.Input()
}
