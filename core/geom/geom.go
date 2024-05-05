package geom

import (
	"gRocketeer/core"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

func createRectangleVerts(rect core.Rectangle, col color.Color) []ebiten.Vertex {
	r, g, b, a := col.RGBA()
	vertices := []ebiten.Vertex{
		{
			DstX:   float32(-(rect.Bounds.X / 2.0)),
			DstY:   float32(rect.Bounds.Y / 2.0),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
		{
			DstX:   float32(-(rect.Bounds.X / 2.0)),
			DstY:   float32(-(rect.Bounds.Y / 2.0)),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
		{
			DstX:   float32(rect.Bounds.X / 2.0),
			DstY:   float32(-(rect.Bounds.Y / 2.0)),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
		{
			DstX:   float32(rect.Bounds.X / 2.0),
			DstY:   float32(rect.Bounds.Y / 2.0),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
	}

	return vertices
}

func createTriangleVerts(rect core.Rectangle, col color.Color) []ebiten.Vertex {
	r, g, b, a := col.RGBA()
	vertices := []ebiten.Vertex{
		{
			DstX:   float32(-(rect.Bounds.X / 2.0)),
			DstY:   float32(rect.Bounds.Y / 2.0),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
		{
			DstX:   0,
			DstY:   float32(-(rect.Bounds.Y / 2.0)),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
		{
			DstX:   float32(rect.Bounds.X / 2.0),
			DstY:   float32(rect.Bounds.Y / 2.0),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(r / 255.0),
			ColorG: float32(g / 255.0),
			ColorB: float32(b / 255.0),
			ColorA: float32(a / 255.0),
		},
	}

	return vertices
}

func ApplyGeomVerts(verts []ebiten.Vertex, rect core.Rectangle, rot float64) []ebiten.Vertex {
	ret := make([]ebiten.Vertex, len(verts))
	for i, v := range verts {
		g := &ebiten.GeoM{}
		g.Rotate(rot * math.Pi / 180)
		g.Translate(rect.Pos.X, rect.Pos.Y)

		x, y := g.Apply(float64(v.DstX), float64(v.DstY))

		ret[i] = v
		ret[i].DstX = float32(x)
		ret[i].DstY = float32(y)
	}

	return ret
}

func MirrorHorizontal(verts []ebiten.Vertex) []ebiten.Vertex {
	ret := make([]ebiten.Vertex, len(verts))
	for i, v := range verts {
		ret[i] = v
		ret[i].DstX *= -1
	}

	return ret
}
