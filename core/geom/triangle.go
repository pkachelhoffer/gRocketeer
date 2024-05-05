package geom

import (
	"gRocketeer/core"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

var (
	whiteImage = ebiten.NewImage(3, 3)
)

func init() {
	whiteImage.Fill(color.White)
}

type Triangle struct {
	Rect core.Rectangle
	Rot  float64
	Col  color.Color

	verts []ebiten.Vertex
}

func NewTriangle(rect core.Rectangle, rot float64, col color.Color) *Triangle {
	tri := &Triangle{
		Rect: rect,
		Rot:  rot,
		Col:  col,
	}

	tri.verts = createTriangleVerts(rect, col)

	return tri
}

func (t *Triangle) Draw(target *ebiten.Image) {
	verts := ApplyGeomVerts(t.verts, t.Rect, t.Rot)

	// Define triangle indices
	indices := []uint16{0, 1, 2} // Indices of the three vertices

	// Draw triangle
	target.DrawTriangles(verts, indices, whiteImage, nil)
}
