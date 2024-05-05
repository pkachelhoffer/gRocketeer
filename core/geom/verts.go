package geom

import (
	"gRocketeer/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type Verts struct {
	Rot  float64
	Rect core.Rectangle

	verts   []ebiten.Vertex
	indices []uint16
}

func NewVerts(verts []ebiten.Vertex, indices []uint16, pos core.Vec, rot float64) *Verts {
	v := &Verts{
		Rot:     rot,
		Rect:    getRectFromVerts(verts),
		verts:   verts,
		indices: indices,
	}
	v.Rect.Pos = pos

	return v
}

func getRectFromVerts(verts []ebiten.Vertex) core.Rectangle {
	var (
		minX, minY, maxX, maxY float32
	)

	for i, v := range verts {
		if i == 0 {
			minX = v.DstX
			minY = v.DstY
			maxX = v.DstX
			maxY = v.DstY
			continue
		}

		if v.DstX < minX {
			minX = v.DstX
		}
		if v.DstX > maxX {
			maxX = v.DstX
		}
		if v.DstY < minY {
			minY = v.DstY
		}
		if v.DstY > maxY {
			maxY = v.DstY
		}
	}

	fminX := float64(minX)
	fmaxX := float64(maxX)
	fminY := float64(minY)
	fmaxY := float64(maxY)

	boundsX := fmaxX - fminX
	boundsY := fmaxY - fminY

	return core.NewRect(core.NewVec((boundsX/2.0)+fminX, (boundsY/2.0)+fminY), core.NewVec(boundsX, boundsY))
}

func (v *Verts) Draw(target *ebiten.Image) {
	verts := ApplyGeomVerts(v.verts, v.Rect, v.Rot)

	// Define triangle indices
	indices := []uint16{0, 1, 2} // Indices of the three vertices

	// Draw triangle
	target.DrawTriangles(verts, indices, whiteImage, nil)
}
