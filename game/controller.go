package game

import (
	"gRocketeer/game/entities"
	"gRocketeer/game/types"
)

type Controller struct {
	g types.IGame

	ship *entities.Ship
}

func NewController(g types.IGame) *Controller {
	return &Controller{g: g}
}

func (c *Controller) LoadLevel() {
	lvl := entities.NewLvlCon(c.g)
	c.g.GetSys().Add("LvlCon", lvl)
}

func (c *Controller) Update() {

}
