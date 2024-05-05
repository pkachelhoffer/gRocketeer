package game

import (
	"gRocketeer/game/entities"
	"gRocketeer/game/types"
)

type Controller struct {
	g types.IGame
}

func NewController(g types.IGame) *Controller {
	return &Controller{g: g}
}

func (c *Controller) ShowGame() {
	ship := entities.NewShip(c.g)
	c.g.Sys().Add("ship", ship)
}
