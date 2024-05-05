package entities

import (
	"fmt"
	"gRocketeer/core"
	"gRocketeer/core/entity"
	"gRocketeer/core/geom"
	"gRocketeer/game/types"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Ship struct {
	entity.BaseEntity
	g types.IGame

	rectBody *geom.Rectangle
	finLeft  *geom.Verts
	finRight *geom.Verts
}

func NewShip(g types.IGame) *Ship {
	return &Ship{
		g: g,
	}
}

func (s *Ship) Init() {
	s.rectBody = geom.NewRectangle(core.NewRect(core.NewVec(s.g.Screen().X/2, s.g.Screen().Y/2), core.NewVec(10, 20)), 0, types.ColorRed)

	fmt.Println(s.rectBody.Rect.BottomLeft(), s.rectBody.Rect.BottomRight())

	verts, ind := getFin(core.NewVec(10, 5), types.ColorRed)
	s.finLeft = geom.NewVerts(verts, ind, s.rectBody.Rect.BottomLeft().Add(core.NewVec(0, 3)), 0)
	s.finRight = geom.NewVerts(geom.MirrorHorizontal(verts), ind, s.rectBody.Rect.BottomRight().Add(core.NewVec(0, 3)), 0)
}

func (s *Ship) Draw(target *ebiten.Image) {
	s.rectBody.Draw(target)
	s.finLeft.Draw(target)
	s.finRight.Draw(target)
}

func getFin(bounds core.Vec, col color.Color) ([]ebiten.Vertex, []uint16) {
	r, g, b, a := col.RGBA()
	verts := []ebiten.Vertex{
		{
			DstX:   float32(bounds.LeftFromCenter()),
			DstY:   float32(bounds.DownFromCenter()),
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
		{
			DstX:   float32(bounds.LeftFromCenter() + 4),
			DstY:   float32(bounds.UpFromCenter()),
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
		{
			DstX:   float32(bounds.RightFromCenter()),
			DstY:   float32(bounds.UpFromCenter()),
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
	}

	return verts, []uint16{0, 1, 2}
}
