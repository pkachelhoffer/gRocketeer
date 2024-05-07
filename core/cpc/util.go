package cpc

import (
	_ "embed"
	"gRocketeer/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"image/color"
)

var (
	whiteImage = ebiten.NewImage(3, 3)

	//go:embed shader.go
	shaderCircleBytes []byte

	shaderCircle *ebiten.Shader
)

func init() {
	whiteImage.Fill(color.White)
	shaderCircle = core.LoadShaderBytes(shaderCircleBytes)
}

func DrawBody(body *cp.Body, target *ebiten.Image, col color.Color) {
	body.EachShape(func(shape *cp.Shape) {
		switch t := shape.Class.(type) {
		case *cp.PolyShape:
			drawPolyShape(t, target, col)
		case *cp.Circle:
			drawCircle(t, target, col)
		}
	})
}

func drawCircle(shape *cp.Circle, target *ebiten.Image, col color.Color) {
	r, g, b, a := col.RGBA()

	pos := shape.TransformC()
	opts := &ebiten.DrawRectShaderOptions{}
	opts.Uniforms = make(map[string]interface{})
	opts.Uniforms["Center"] = []float32{
		float32(pos.X),
		float32(pos.Y),
	}
	opts.Uniforms["Radius"] = float32(shape.Radius())
	opts.Uniforms["Color"] = []float32{
		float32(r) / 0xffff,
		float32(g) / 0xffff,
		float32(b) / 0xffff,
		float32(a) / 0xffff,
	}

	opts.GeoM.Translate(pos.X-shape.Radius(), pos.Y-shape.Radius())

	target.DrawRectShader(int(shape.Radius()*2), int(shape.Radius()*2), shaderCircle, opts)
}

func drawPolyShape(shape *cp.PolyShape, target *ebiten.Image, col color.Color) {
	verts, indices := getVerticesPolyShape(shape, col)
	target.DrawTriangles(verts, indices, whiteImage, nil)
}

func getVerticesPolyShape(shape *cp.PolyShape, col color.Color) ([]ebiten.Vertex, []uint16) {
	r, g, b, a := col.RGBA()
	var vertexes []ebiten.Vertex
	for i := 0; i < shape.Count(); i++ {
		v := shape.Vert(i)
		loc := shape.Body().LocalToWorld(v)
		vertexes = append(vertexes, ebiten.Vertex{
			DstX:   float32(loc.X),
			DstY:   float32(loc.Y),
			ColorR: float32(r) / 0xffff,
			ColorG: float32(g) / 0xffff,
			ColorB: float32(b) / 0xffff,
			ColorA: float32(a) / 0xffff,
		})
	}

	return vertexes, []uint16{0, 1, 2, 0, 2, 3}
}
