package game

import (
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
	c.ship = entities.NewShip(c.g)
	c.g.GetSys().Add("ship", c.ship)
}

func (c *Controller) Input() {
	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		c.ship.Thrust()
	} else {
		if inpututil.KeyPressDuration(ebiten.KeyQ) > 0 {
			c.ship.Turn(true)
		} else if inpututil.KeyPressDuration(ebiten.KeyE) > 0 {
			c.ship.Turn(false)
		}
	}
}

func (c *Controller) Update() {
	c.Input()
}
