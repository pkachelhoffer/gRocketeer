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
	c.ship = entities.NewShip(c.g)
	c.g.GetSys().Add("ship", c.ship)

	w := entities.NewWall(c.g, core.NewRect(core.NewVec(c.g.Screen().RightFromCenter(), 25), core.NewVec(920, 20)), types.ColorWall)
	c.g.GetSys().Add("wallTop", w)

	w = entities.NewWall(c.g, core.NewRect(core.NewVec(c.g.Screen().RightFromCenter(), c.g.Screen().Y-25), core.NewVec(920, 20)), types.ColorWall)
	c.g.GetSys().Add("wallBottom", w)

	w = entities.NewWall(c.g, core.NewRect(core.NewVec(c.g.Screen().X-25, c.g.Screen().DownFromCenter()), core.NewVec(20, 500)), types.ColorWall)
	c.g.GetSys().Add("wallRight", w)

	w = entities.NewWall(c.g, core.NewRect(core.NewVec(25, c.g.Screen().DownFromCenter()), core.NewVec(20, 500)), types.ColorWall)
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
