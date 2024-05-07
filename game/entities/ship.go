package entities

import (
	"gRocketeer/core"
	"gRocketeer/core/cpc"
	"gRocketeer/core/entity"
	"gRocketeer/core/geom"
	"gRocketeer/game/resources"
	"gRocketeer/game/types"
	"gRocketeer/game/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"image/color"
)

type Ship struct {
	entity.BaseEntity
	g types.IGame

	body *cp.Body

	circHead *geom.Circle
	rectBody *geom.Rectangle
	finLeft  *geom.Verts
	finRight *geom.Verts
}

func NewShip(g types.IGame) *Ship {
	return &Ship{
		g: g,
	}
}

func (s *Ship) posHead() core.Vec {
	return s.rectBody.Rect.Top().Add(core.NewVec(0, -s.circHead.Rad))
}

func (s *Ship) posFin(left bool) core.Vec {
	if left {
		return s.rectBody.Rect.BottomLeft().Add(core.NewVec(0, s.finLeft.Rect.Bounds.DownFromCenter()))
	} else {
		return s.rectBody.Rect.BottomRight().Add(core.NewVec(0, s.finRight.Rect.Bounds.DownFromCenter()))
	}
}

func (s *Ship) setPos(pos core.Vec) {
	s.rectBody.Rect.Pos = pos
	s.circHead.Pos = s.posHead()
	s.finLeft.Rect.Pos = s.posFin(true)
	s.finRight.Rect.Pos = s.posFin(false)
}

func (s *Ship) Init() {
	s.rectBody = geom.NewRectangle(core.NewRect(core.Vec{}, core.NewVec(10, 20)), 0, types.ColorRed)

	s.circHead = geom.NewCircle(core.Vec{}, 5, resources.ShaderCircle)

	verts, ind := getFin(core.NewVec(10, 5), types.ColorLightRed)
	s.finLeft = geom.NewVerts(verts, ind, core.Vec{}, 0)

	s.finRight = geom.NewVerts(geom.MirrorHorizontal(verts), ind, s.rectBody.Rect.BottomRight().Add(core.NewVec(0, 3)), 0)

	s.setPos(core.NewVec(s.g.Screen().X/2, s.g.Screen().Y/2))

	space := s.g.GetSpace()

	s.body = cp.NewBody(1, cp.MomentForBox(1, 10, 20))

	space.AddBody(s.body)

	shape := cp.NewBox(s.body, s.rectBody.Rect.Bounds.X, s.rectBody.Rect.Bounds.Y, 0.5)
	space.AddShape(shape)

	shape2 := cp.NewCircle(s.body, 5, cp.Vector{X: 0, Y: -10})
	space.AddShape(shape2)

	s.body.SetPosition(util.ToVec(s.rectBody.Rect.Pos))
}

func (s *Ship) Draw(target *ebiten.Image) {
	//s.rectBody.Draw(target)
	//s.finLeft.Draw(target)
	//s.finRight.Draw(target)
	//s.circHead.Draw(target)
	//s.DrawBody(target)
	cpc.DrawBody(s.body, target, types.ColorGrey)
}

func (s *Ship) Update() bool {
	//s.setPos(util.FromVec(s.body.Position()))
	//s.rectBody.Rect.Pos = util.FromVec(s.body.Position())
	//s.rectBody.Rot = s.body.Angle()
	return false
}

func getFin(bounds core.Vec, col color.Color) ([]ebiten.Vertex, []uint16) {
	r, g, b, a := col.RGBA()
	verts := []ebiten.Vertex{
		{
			DstX:   float32(bounds.LeftFromCenter()),
			DstY:   float32(bounds.DownFromCenter()),
			ColorR: float32(r) / 0xffff,
			ColorG: float32(g) / 0xffff,
			ColorB: float32(b) / 0xffff,
			ColorA: float32(a) / 0xffff,
		},
		{
			DstX:   float32(bounds.LeftFromCenter() + bounds.X/2.5),
			DstY:   float32(bounds.UpFromCenter()),
			ColorR: float32(r) / 0xffff,
			ColorG: float32(g) / 0xffff,
			ColorB: float32(b) / 0xffff,
			ColorA: float32(a) / 0xffff,
		},
		{
			DstX:   float32(bounds.RightFromCenter()),
			DstY:   float32(bounds.UpFromCenter()),
			ColorR: float32(r) / 0xffff,
			ColorG: float32(g) / 0xffff,
			ColorB: float32(b) / 0xffff,
			ColorA: float32(a) / 0xffff,
		},
	}

	return verts, []uint16{0, 1, 2}
}

func (s *Ship) Thrust() {
	s.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -100}, cp.Vector{X: 0, Y: -10})
}

func (s *Ship) Turn(clockWise bool) {
	if clockWise {
		s.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -10}, cp.Vector{X: -5, Y: -10})
	} else {
		s.body.ApplyForceAtLocalPoint(cp.Vector{X: 0, Y: -10}, cp.Vector{X: 5, Y: -10})
	}
}

func (s *Ship) DrawBody(target *ebiten.Image) {
	s.body.EachShape(func(shape *cp.Shape) {
		poly := shape.Class.(*cp.PolyShape)
		r, g, b, a := types.ColorRed.RGBA()
		var vertexes []ebiten.Vertex
		for i := 0; i < poly.Count(); i++ {
			v := poly.Vert(i)
			loc := s.body.LocalToWorld(v)
			vertexes = append(vertexes, ebiten.Vertex{
				DstX:   float32(loc.X),
				DstY:   float32(loc.Y),
				ColorR: float32(r / 255.0),
				ColorG: float32(g / 255.0),
				ColorB: float32(b / 255.0),
				ColorA: float32(a / 255.0),
			})
		}

		verts := geom.NewVerts(vertexes, []uint16{0, 1, 2, 0, 2, 3}, core.Vec{}, 0)
		verts.Draw(target)
	})
}
