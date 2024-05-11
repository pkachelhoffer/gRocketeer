package entities

import (
	"gRocketeer/core"
	"gRocketeer/game/resources"
	"gRocketeer/game/types"
	"github.com/jakecoffman/cp"
)

func CreateShip(g types.IGame, plan ShipPlan) *Ship {
	ship := NewShip(g)

	hs := CompSize / 2

	space := g.GetSpace()
	body := cp.NewBody(float64(1*len(plan.Nodes)), cp.INFINITY)

	space.AddBody(body)
	for _, node := range plan.Nodes {
		pos := core.NewVec(float64(node.PosX*CompSize), float64(node.PosY*CompSize))

		verts := []cp.Vector{
			{pos.X - float64(hs), pos.Y - float64(hs)},
			{pos.X + float64(hs), pos.Y - float64(hs)},
			{pos.X + float64(hs), pos.Y + float64(hs)},
			{pos.X - float64(hs), pos.Y + float64(hs)},
		}
		shape := cp.NewPolyShapeRaw(body, 4, verts, 0.5)
		space.AddShape(shape)

		node.Component.SetPhys(body, shape)
		ship.AddComponent(node.Component)
	}

	ship.SetBody(body)

	return ship
}

type ShipPlanNode struct {
	PosX, PosY int
	Component  IShipComp
}

type ShipPlan struct {
	Nodes []ShipPlanNode
}

var ShipPlan1 = ShipPlan{
	Nodes: []ShipPlanNode{
		{
			PosX:      0,
			PosY:      0,
			Component: NewShipCompHull(core.NewSprite(resources.ImgShipHull01)),
		},
		{
			PosX:      -1,
			PosY:      0,
			Component: NewShipCompHull(core.NewSprite(resources.ImgShipHull01)),
		},
		{
			PosX:      1,
			PosY:      0,
			Component: NewShipCompHull(core.NewSprite(resources.ImgShipHull01)),
		},
		{
			PosX:      0,
			PosY:      -1,
			Component: NewShipCompHull(core.NewSprite(resources.ImgShipHull01)),
		},
	},
}
