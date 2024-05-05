package game

import (
	"gRocketeer/core"
	"gRocketeer/game/assets/shaders"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ShaderTest   *ebiten.Shader
	ShaderCircle *ebiten.Shader
)

func init() {
	ShaderTest = core.LoadShaderBytes(shaders.ShaderTest)
	ShaderCircle = core.LoadShaderBytes(shaders.ShaderCircle)
}
