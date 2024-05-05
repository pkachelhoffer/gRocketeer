package geom

import (
	"gRocketeer/core"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Rectangle struct {
	Rect core.Rectangle
	Rot  float64
	Col  color.Color

	verts []ebiten.Vertex

	Shader *ebiten.Shader

	img *ebiten.Image
}

func NewRectangle(rect core.Rectangle, rot float64, col color.Color) *Rectangle {
	rec := &Rectangle{
		Rect: rect,
		Rot:  rot,
		Col:  col,
		img:  ebiten.NewImage(100, 100),
	}

	rec.verts = createRectangleVerts(rect, col)

	return rec
}

func (t *Rectangle) Draw(target *ebiten.Image) {
	verts := ApplyGeomVerts(t.verts, t.Rect, t.Rot)

	// Define triangle indices
	indices := []uint16{0, 1, 2, 0, 2, 3} // Indices of the three vertices

	if t.Shader != nil {
		target.DrawTrianglesShader(verts, indices, t.Shader, nil)
	} else {
		target.DrawTriangles(verts, indices, whiteImage, nil)
	}
}
