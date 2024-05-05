package geom

import (
	"gRocketeer/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type Circle struct {
	Pos  core.Vec
	Rad  float64
	Shad *ebiten.Shader
}

func NewCircle(pos core.Vec, rad float64, shad *ebiten.Shader) *Circle {
	return &Circle{
		Pos:  pos,
		Rad:  rad,
		Shad: shad,
	}
}

func (c *Circle) Draw(target *ebiten.Image) {
	opts := &ebiten.DrawRectShaderOptions{}
	opts.Uniforms = make(map[string]interface{})
	opts.Uniforms["Center"] = []float32{
		float32(c.Pos.X),
		float32(c.Pos.Y),
	}
	opts.Uniforms["Radius"] = float32(c.Rad)
	opts.GeoM.Translate(c.Pos.X-c.Rad, c.Pos.Y-c.Rad)

	target.DrawRectShader(int(c.Rad*2), int(c.Rad*2), c.Shad, opts)
}
